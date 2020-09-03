package resource

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	c7nclient "github.com/choerodon/c7nctl/pkg/client"
	c7nconsts "github.com/choerodon/c7nctl/pkg/common/consts"
	c7nerrors "github.com/choerodon/c7nctl/pkg/common/errors"
	c7ncfg "github.com/choerodon/c7nctl/pkg/config"
	c7nslaver "github.com/choerodon/c7nctl/pkg/slaver"
	c7nutils "github.com/choerodon/c7nctl/pkg/utils"
	std_errors "github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	yaml_v2 "gopkg.in/yaml.v2"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"path/filepath"
	"text/template"
)

const (
	eurekaClientServerUrlTpl = "%s://%s:8000/eureka/"
	resourceDomainUrlTpl     = "%s://%s"
)

type InstallDefinition struct {
	// api 版本
	Version string
	// Choerodon 平台版本
	PaaSVersion string
	Metadata    Metadata
	Spec        Spec
}

type Metadata struct {
	Name      string
	Namespace string
}

type Spec struct {
	Basic     Basic
	Resources v1.ResourceRequirements
	Release   []*Release
	Runner    *Release `json:"runner"`
	Component []*Release
}

type Basic struct {
	CommonLabels       map[string]string
	DefaultAccessModes []v1.PersistentVolumeAccessMode
	StorageClass       string

	Prefix string
	// 默认为空
	ImageRepository string
	ChartRepository string
	DatasourceTpl   string
	ThinMode        bool

	SkipInput bool
	Timeout   int
	Slaver    c7nslaver.Slaver
}

func (i *InstallDefinition) GetInstallDefinition(resource string) error {
	// TODO 如果 resource 是域名形式, https:// 会变成 https:/
	res, err := c7nutils.GetResource(filepath.Join(resource, c7nconsts.InstallConfigPath))
	if err != nil {
		return err
	}

	rdJson, err := yaml.ToJSON(res)
	if err != nil {
		panic(err)
	}
	// slaver 使用了 core_v1.ContainerPort, 必须先转 JSON
	_ = json.Unmarshal(rdJson, i)

	if i.Spec.Basic.DefaultAccessModes == nil {
		i.Spec.Basic.DefaultAccessModes = []v1.PersistentVolumeAccessMode{"ReadWriteOnce"}
	}
	return nil
}

// TODO 渲染也是有先后顺序的
func (i *InstallDefinition) RenderRelease(r *Release) error {
	task, err := c7nclient.GetTask(r.Name)
	if err != nil {
		if std_errors.Is(err, c7nerrors.TaskInfoIsNotFoundError) {
			task = c7nclient.NewReleaseTask(r.Name, r.Namespace, r.Version, r.Prefix)
		} else {
			return std_errors.WithMessage(err, fmt.Sprintf("render release %s failed", r.Name))
		}
	}
	if task.Status == c7nconsts.UninitializedStatus {
		// TODO 将合并 config.yaml 移到获取 installDef 中
		if err = i.renderValues(r); err != nil {
			return err
		}
		if err := i.render(r); err != nil {
			return err
		}

		// 保存渲染完成的 r
		task.Values = r.Values
		// 执行 Release Job 时需要
		task.Resource = *r.Resource
		task.Status = c7nconsts.RenderedStatus
		if _, err = c7nclient.SaveTask(*task); err != nil {
			return err
		}
	} else {
		// 当 r 渲染完成但是没有完成安装——c7nctl install 会中断，二次执行
		r.Values = task.Values
		r.Resource = &task.Resource
		// 重新渲染 preCommand 等，避免在 TaskInfo 加入 PreCommand 导致循环依赖
		if err := i.render(r); err != nil {
			return err
		}
	}
	if err = i.CheckReleaseDomain(r.Values); err != nil {
		return err
	}
	log.Infof("Successfully rendered the Release %s", r.Name)
	return nil
}

func (i *InstallDefinition) RenderComponent(rls *Release) error {
	rlsByte, _ := yaml_v2.Marshal(rls)
	renderedRls, err := i.renderTpl(rls.Name, string(rlsByte))
	if err != nil {
		return err
	}
	return yaml_v2.Unmarshal(renderedRls.Bytes(), rls)
}

// 必须基于 InstallDefinition 渲染 value.yaml 文件
func (i *InstallDefinition) RenderHelmValues(r *Release, resource, helmValues string) (map[string]interface{}, error) {
	rlsVals := r.HelmValues()
	var fileValsByte bytes.Buffer
	fileVals, err := r.ValuesRaw(resource, helmValues)
	if err != nil {
		return nil, err
	}
	fileValsByte, err = i.renderTpl(r.Name+"-file-values", fileVals)
	if err != nil {
		return nil, err
	}

	return c7nutils.Vals(rlsVals, fileValsByte.String())
}

func (i *InstallDefinition) SetPrefix(prefix string) {
	i.Spec.Basic.Prefix = prefix
}
func (i *InstallDefinition) SetImageRepository(imageRepo string) {
	i.Spec.Basic.ImageRepository = imageRepo
}
func (i *InstallDefinition) SetChartRepository(chartRepo string) {
	i.Spec.Basic.ChartRepository = chartRepo
}
func (i *InstallDefinition) SetDatasourceTpl(dsTpl string) {
	i.Spec.Basic.DatasourceTpl = dsTpl
}
func (i *InstallDefinition) SetThinMode(thinMode bool) {
	i.Spec.Basic.ThinMode = thinMode
}
func (i *InstallDefinition) SetStorageClass(sc string) {
	i.Spec.Basic.StorageClass = sc
}

// 将 config.yml 中的值合并到 Release.Resource
func (i *InstallDefinition) MergerConfig(uc *c7ncfg.C7nConfig) {
	i.mergerResource(uc)

	if uc.GetStorageClass() != "" {
		i.SetStorageClass(uc.GetStorageClass())
	}
	if uc.GetPrefix() != "" {
		i.SetPrefix(uc.GetPrefix())
	}
	if uc.GetImageRepository() != "" {
		i.SetImageRepository(uc.GetImageRepository())
	}
	if uc.GetChartRepository() != "" {
		i.SetChartRepository(uc.GetChartRepository())
	}
	if uc.GetDatasourceTpl() != "" {
		i.SetDatasourceTpl(uc.GetDatasourceTpl())
	}
	if uc.Spec.ThinMode {
		i.SetThinMode(true)
	}
}

func (i *InstallDefinition) CheckReleaseDomain(values []c7nclient.ChartValue) error {
	for _, v := range values {
		// TODO 添加本地方式检查域名
		if v.Check == "clusterdomain" {
			log.Debugf("Value %s: %s, checking: %s", v.Name, v.Value, v.Check)
			if err := i.Spec.Basic.Slaver.CheckClusterDomain(v.Value); err != nil {
				log.Errorf("请检查您的域名: %s 已正确解析到集群", v.Value)
				return err
			}
		}
	}
	return nil
}

func (i *InstallDefinition) mergerResource(uc *c7ncfg.C7nConfig) {
	for _, rls := range i.Spec.Release {
		if res := uc.GetResource(rls.Name); res == nil {
			log.Debugf("There is no resource in config.yaml of Release %s", rls.Name)
		} else {
			// 直接使用外部配置
			if res.External {
				rls.Resource = res
			} else {
				if res.Domain != "" {
					if !c7nutils.CheckDomain(res.Domain) {
						log.Errorf("domain name %s is not in compliance with the rules", res.Domain)
					}
					rls.Resource.Domain = res.Domain
				}
				if res.Schema != "" {
					if !c7nutils.CheckSchema(res.Schema) {
						log.Errorf("schema %s is not in compliance with the rules", res.Schema)
					}
					rls.Resource.Schema = res.Schema
				}
				if res.Username != "" {
					rls.Resource.Username = res.Username
				}
				if res.Password != "" {
					rls.Resource.Password = res.Password
				}
				// TODO 其他的配置项是否需要初始化到 rls 中
				/*
					rlsType := reflect.TypeOf(*rls.Resource)
					rlsValue := reflect.ValueOf(rls.Resource).Elem()
					resType := reflect.TypeOf(*res)
					resValue := reflect.ValueOf(*res)
					for i := 0; i < resType.NumField(); i ++ {
						if f, ok := rlsType.FieldByName(resType.Field(i).Name); ok {
							switch resValue.Field(i).Kind() {
							case reflect.String: {
								rlsValue.FieldByName(f.Name).SetString(resValue.FieldByName(f.Name).String())
							}
							case reflect.Bool: {
								rlsValue.FieldByName(f.Name).SetBool(resValue.FieldByName(f.Name).Bool())
							}
							case reflect.Int32: {
								rlsValue.FieldByName(f.Name).SetInt(resValue.FieldByName(f.Name).Int())
							}
							default:
								log.Debugf("can't convert Field %s from C7nConfig to InstallDefinition", f.Name)
							}
						} else {
							log.Debugf("Resource of C7nConfig field %s isn't in Release of InstallDefinition", resType.Field(i).Name)
						}
					}
				*/
			}
		}
	}
}

// 渲染 release
func (i *InstallDefinition) render(r *Release) error {
	rlsByte, _ := yaml_v2.Marshal(r)
	renderedRls, err := i.renderTpl(r.Name, string(rlsByte))
	if err != nil {
		return err
	}
	if err := yaml_v2.Unmarshal(renderedRls.Bytes(), r); err != nil {
		return std_errors.WithMessage(err, fmt.Sprintf("Unmarshal Release %s failed", r))
	}
	return nil
}

// 传指针的方式好呢，还是返回值的方式好？
//
// 在渲染 release 前将 values 渲染完成
// 获取用户输入或者根据 value 的模版值渲染
func (i *InstallDefinition) renderValues(rls *Release) error {
	if rls.Values == nil {
		log.Debugf("release %s values is empty", rls.Name)
		return nil
	}
	for idx, v := range rls.Values {
		// 输入 value
		if v.Input.Enabled && !i.Spec.Basic.SkipInput {
			var err error
			var value string
			if v.Input.Password {
				v.Input.Twice = true
				value, err = c7nutils.AcceptUserPassword(v.Input)
			} else {
				value, err = c7nutils.AcceptUserInput(v.Input)
			}
			if err != nil {
				return err
			}
			// v.Values 是复制
			rls.Values[idx].Value = value
		} else {
			v, err := i.renderTpl(v.Name+"-values", v.Value)
			if err != nil {
				return err
			}
			rls.Values[idx].Value = v.String()
		}
	}
	return nil
}

// 根据模版和 InstallDefinition 渲染
func (i *InstallDefinition) renderTpl(name, tplStr string) (bytes.Buffer, error) {
	tpl, err := template.New(name).Funcs(c7nutils.C7nFunc).Parse(tplStr)
	if err != nil {
		return bytes.Buffer{}, err
	}
	var result bytes.Buffer
	err = tpl.Execute(&result, i)
	if err != nil {
		return bytes.Buffer{}, std_errors.WithMessage(err, fmt.Sprintf("Failed to render release %s", name))
	}
	return result, nil
}

/*
  template 内嵌函数
*/
func (i *InstallDefinition) WithPrefix() string {
	if i.Spec.Basic.Prefix == "" {
		return ""
	}
	return i.Spec.Basic.Prefix + "-"
}

func (i *InstallDefinition) GetReleaseName(rlsName string) string {
	return i.WithPrefix() + rlsName
}

// TODO add storageClassName()
func (i *InstallDefinition) GetStorageClass() string {
	//return c7nctx.Ctx.UserConfig.GetStorageClassName()
	return i.Spec.Basic.StorageClass
}

func (i *InstallDefinition) GetDatabaseUrl(rls string) string {
	return fmt.Sprintf(i.Spec.Basic.DatasourceTpl, i.GetReleaseName("mysql"), i.GetReleaseName(rls))
}

func (i *InstallDefinition) GetResource(rls string) *c7ncfg.Resource {
	for _, r := range i.Spec.Release {
		if r.Name == rls {
			return r.Resource
		}
	}
	log.Fatal("Release cannot be empty")
	return nil
}

func (i *InstallDefinition) GetReleaseValue(rls, value string) string {
	for _, r := range i.Spec.Release {
		if r.Name == rls {
			for _, v := range r.Values {
				if v.Name == value {
					return v.Value
				}
			}
			log.WithField("Release values", value).Fatal("Release value cannot be empty")
		}
	}
	log.WithField("Release", rls).Fatal("Release cannot be empty")
	return ""
}

func (i *InstallDefinition) EncryptGitlabAccessToken() string {
	token := i.GetReleaseValue("gitlab-service", "env.open.GITLAB_PRIVATETOKEN")
	dbKeyBase := i.GetReleaseValue("gitlab", "core.env.GITLAB_SECRETS_DB_KEY_BASE")
	str := fmt.Sprintf("%s%s", token, dbKeyBase[:32])

	hash := sha256.New()
	hash.Write([]byte(str))

	// to lowercase hexits
	hex.EncodeToString(hash.Sum(nil))

	// to base64
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func (i *InstallDefinition) GetPersistence(rls string, index int) *Persistence {
	for _, r := range i.Spec.Release {
		if r.Name == rls && len(r.Persistence) > index {
			return r.Persistence[index]
		}
	}
	log.WithField("Release", rls).Fatal("Release cannot be empty")
	return nil
}

func (i *InstallDefinition) GetRunnerPersistence(index int) *Persistence {
	if len(i.Spec.Runner.Persistence) > index {
		return i.Spec.Runner.Persistence[index]
	}
	log.WithField("Release", "gitlab-runner").Fatal("Release cannot be empty")
	return nil
}

func (i *InstallDefinition) GetRunnerValues(values string) string {
	for _, v := range i.Spec.Runner.Values {
		if v.Name == values {
			return v.Value
		}
	}
	return ""
}

func (i *InstallDefinition) IsThinMode() bool {
	return i.Spec.Basic.ThinMode
}

func (i *InstallDefinition) GetEurekaUrl() string {
	for _, r := range i.Spec.Release {
		if r.Name == c7nconsts.HzeroRegister {
			return fmt.Sprintf(eurekaClientServerUrlTpl, r.Resource.Schema, r.Resource.Host)
		}
	}
	return ""
}

func (i *InstallDefinition) GetResourceDomainUrl(rls string) string {
	for _, r := range i.Spec.Release {
		if r.Name == rls {
			return fmt.Sprintf(resourceDomainUrlTpl, r.Resource.Schema, r.Resource.Domain)
		}
	}
	return ""
}

func (i *InstallDefinition) GetImageRepository() string {
	return i.Spec.Basic.ImageRepository
}
