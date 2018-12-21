package app

import (
	"github.com/choerodon/c7n/pkg/common"
	"github.com/choerodon/c7n/pkg/config"
	"github.com/choerodon/c7n/pkg/helm"
	"github.com/choerodon/c7n/pkg/install"
	kube2 "github.com/choerodon/c7n/pkg/kube"
	"github.com/choerodon/c7n/pkg/upgrade"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/vinkdong/gox/log"
	"io/ioutil"
	helm_env "k8s.io/helm/pkg/helm/environment"
	"k8s.io/helm/pkg/kube"
	"os"
)

var (
	tlsServerName string // overrides the server name used to verify the hostname on the returned certificates from the server.
	tlsCaCertFile string // path to TLS CA certificate file
	tlsCertFile   string // path to TLS certificate file
	tlsKeyFile    string // path to TLS key file
	tlsVerify     bool   // enable TLS and verify remote certificates
	tlsEnable     bool   // enable TLS

	tlsCaCertDefault = "$HELM_HOME/ca.pem"
	tlsCertDefault   = "$HELM_HOME/cert.pem"
	tlsKeyDefault    = "$HELM_HOME/key.pem"

	tillerTunnel     *kube.Tunnel
	settings         helm_env.EnvSettings
	ResourceFile     string
	client           *helm.Client
	defaultNamespace = "choerodon"
	UserConfig       *config.Config
)

const (
	repoUrl       = "https://openchart.choerodon.com.cn/choerodon/c7n/"
	C7nLabelKey   = "c7n-usage"
	C7nLabelValue = "c7n-installer"
)

func getUserConfig(filePath string) *config.Config {
	if filePath == "" {
		log.Debugf("no user config defined by `-c`")
		return nil
	}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error(err)
		os.Exit(124)
	}
	userConfig := &config.Config{}
	err = yaml.Unmarshal(data, userConfig)
	if err != nil {
		log.Error(err)
		os.Exit(124)
	}
	return userConfig

}
func TearDown() {
	tillerTunnel.Close()
}

func GetInstall(cmd *cobra.Command, args []string) *install.Install {
	var ResourceFile string
	var err error
	// get install configFile
	ResourceFile, err = cmd.Flags().GetString("resource-file")
	if err != nil {
		log.Error(err)
		os.Exit(123)
	}
	configFile, err := cmd.Flags().GetString("config-file")
	UserConfig = getUserConfig(configFile)

	prefix, _ := cmd.Flags().GetString("prefix")

	r := config.ResourceDefinition{}
	r.LocalFile = ResourceFile
	var installDef = &install.Install{}

	installDef.Prefix = prefix

	version, err := cmd.Flags().GetString("version")
	if err != nil {
		log.Error(err)
		os.Exit(128)
	}

	data, err := r.GetResourceDate(version)
	if err != nil {
		log.Error(err)

	}
	yaml.Unmarshal(data, installDef)
	if installDef.Version == "" {
		log.Error("get install config error")
		os.Exit(127)
	}

	installDef.UserConfig = UserConfig

	commonLabels := make(map[string]string)
	commonLabels[C7nLabelKey] = C7nLabelValue
	installDef.CommonLabels = commonLabels
	// prepare environment
	tillerTunnel = kube2.GetTunnel()
	helmClient := &helm.Client{
		Tunnel: tillerTunnel,
	}
	helmClient.InitClient()
	installDef.HelmClient = helmClient

	if disable, _ := cmd.Flags().GetBool("no-timeout"); disable {
		installDef.Timeout = 60 * 60 * 24
	}

	if UserConfig == nil {
		installDef.Namespace = "c7n-system"
	} else {
		installDef.Namespace = UserConfig.Metadata.Namespace
	}

	if installDef.SkipInput, err = cmd.Flags().GetBool("skip-input"); err != nil {
		log.Error(err)
		os.Exit(127)
	}
	return installDef
}

func Install(cmd *cobra.Command, args []string, mail string) error {

	InstallDef := GetInstall(cmd, args)
	InstallDef.Mail = mail

	defer TearDown()
	//tunnel.Close()
	// do install
	return InstallDef.Run(args...)
}

func Upgrade(cmd *cobra.Command, args []string) error {
	ResourceFile, err := cmd.Flags().GetString("resource-file")
	if err != nil {
		return err
	}
	r := config.ResourceDefinition{}
	r.LocalFile = ResourceFile
	version, err := cmd.Flags().GetString("version")
	data, err := r.GetUpgradeResourceDate(version)
	if err != nil {
		return err
	}
	u := upgrade.Upgrader{}
	yaml.Unmarshal(data, &u)
	//tunnel.Close()
	defer TearDown()
	// do upgrade
	return u.Run(args...)
}

func Delete(cmd *cobra.Command, args []string) error {
	var err error

	defer TearDown()
	//tunnel.Close()

	// prepare environment
	tillerTunnel = kube2.GetTunnel()

	kubeClient := kube2.GetClient()

	helmClient := &helm.Client{
		Tunnel:     tillerTunnel,
		KubeClient: kubeClient,
	}
	helmClient.InitClient()

	ns, err := cmd.Flags().GetString("namespace")
	if err != nil {
		return err
	}

	ctx := install.Context{
		Client:    kubeClient,
		Namespace: ns,
		Metrics:   common.Metrics{},
	}

	for _, a := range args {
		if err := ctx.DeleteSucceed(a, ns, install.ReleaseTYPE); err == nil {
			log.Successf("deleted %s", a)
		} else {
			log.Error(err)
			log.Errorf("delete %s failed", a)
		}
		ctx.DeleteSucceedTask(a)
	}

	// do delete
	return err
}
