package consts

const HostFile = `; 将所有节点信息在这里填写
;    第一个字段                  为远程服务器内网IP
;    第二个字段 ansible_port     为节点 sshd 监听端口
;    第三个字段 ansible_user     为节点远程登录用户名
;    第四个字段 ansible_ssh_pass 为节点远程登录用户密码
[all]
{{- range .MasterIPs }} 
{{ . }} ansible_port=22 ansible_user={{ $.Ssh.Username }} ansible_ssh_pass={{ $.Ssh.Password }}
{{- end }}
{{- range .NodeIPs }} 
{{ . }} ansible_port=22 ansible_user={{ $.Ssh.Username }} ansible_ssh_pass={{ $.Ssh.Password }}
{{- end }}
; 单 master 节点不需要进行负载均衡，lb节点组留空。
[lb]

; 注意etcd集群必须是1,3,5,7...奇数个节点
[etcd]
{{- range .MasterIPs }} 
{{ . }}
{{- end }}

[kube-master]
{{- range .MasterIPs }} 
{{ . }}
{{- end }}

[kube-worker]
{{- range .MasterIPs }} 
{{ . }}
{{- end }}
{{- range .NodeIPs }} 
{{ . }}
{{- end }}

; 预留组，后续添加master节点使用
[new-master]

; 预留组，后续添加worker节点使用
[new-worker]

; 预留组，后续添加etcd节点使用
[new-etcd]

;-------------------------------------- 以下为基础信息配置 ------------------------------------;
[all:vars]
; 是否跳过节点物理资源校验，Master节点要求2c2g以上，Worker节点要求2c4g以上
skip_verify_node=false
; kubernetes版本
kube_version="1.16.9"
; 负载均衡器
;   有 openresty、nginx、haproxy、envoy  和 slb 可选，默认使用 openresty
;   为什么单 master 集群 apiserver 也使用了负载均衡请参与此讨论： https://github.com/TimeBye/kubeadm-ha/issues/8
lb_mode="openresty"
; 使用负载均衡后集群 apiserver ip，设置 lb_kube_apiserver_ip 变量，则启用负载均衡器 + keepalived
; lb_kube_apiserver_ip="192.168.56.15"
; 使用负载均衡后集群 apiserver port
lb_kube_apiserver_port="8443"

; 网段选择：pod 和 service 的网段不能与服务器网段重叠，
; 若有重叠请配置 "kube_pod_subnet" 和 "kube_service_subnet" 变量设置 pod 和 service 的网段，示例参考
;    如果服务器网段为：10.0.0.1/8
;       pod 网段可设置为：192.168.0.0/18
;       service 网段可设置为 192.168.64.0/18
;    如果服务器网段为：172.16.0.1/12
;       pod 网段可设置为：10.244.0.0/18
;       service 网段可设置为 10.244.64.0/18
;    如果服务器网段为：192.168.0.1/16
;       pod 网段可设置为：10.244.0.0/18
;       service 网段可设置为 10.244.64.0/18
; 集群pod ip段，默认掩码位 18 即 16384 个ip
kube_pod_subnet="10.244.0.0/18"
; 集群service ip段
kube_service_subnet="10.244.64.0/18"
; 分配给节点的 pod 子网掩码位，默认为 24 即 256 个ip，故使用这些默认值可以纳管 16384/256=64 个节点。
kube_network_node_prefix="24"

; node节点最大 pod 数。数量与分配给节点的 pod 子网有关，ip 数应大于 pod 数。
; https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr
kube_max_pods="110"

; 集群网络插件，目前支持flannel,calico,kube-ovn
network_plugin="calico"

; 若服务器磁盘分为系统盘与数据盘，请修改以下路径至数据盘自定义的目录。
; Kubelet 根目录
kubelet_root_dir="/var/lib/kubelet"
; docker容器存储目录
docker_storage_dir="/var/lib/docker"
; Etcd 数据根目录
etcd_data_dir="/var/lib/etcd"
`