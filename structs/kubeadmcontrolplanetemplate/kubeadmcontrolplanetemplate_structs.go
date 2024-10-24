/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kubeadmcontrolplanetemplate

type KubeadmControlPlaneTemplate struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type Spec struct {
	Template ControlPlaneTemplate `yaml:"template"`
}

type ControlPlaneTemplate struct {
	Spec ControlPlaneSpec `yaml:"spec"`
}

type ControlPlaneSpec struct {
	KubeadmConfigSpec KubeadmConfigSpec `yaml:"kubeadmConfigSpec"`
}

type KubeadmConfigSpec struct {
	ClusterConfiguration ClusterConfiguration `yaml:"clusterConfiguration"`
	InitConfiguration    InitConfiguration    `yaml:"initConfiguration"`
	JoinConfiguration    JoinConfiguration    `yaml:"joinConfiguration"`
}

type ClusterConfiguration struct {
	APIServer         APIServer         `yaml:"apiServer"`
	ControllerManager ControllerManager `yaml:"controllerManager"`
}

type APIServer struct {
	CertSANs []string `yaml:"certSANs"`
}

type ControllerManager struct {
	ExtraArgs EnablHostPathProvisioner `yaml:"extraArgs"`
}

type EnablHostPathProvisioner struct {
	EnableHostPathProvisioner string `yaml:"enable-hostpath-provisioner"`
}

type InitConfiguration struct {
	NodeRegistration map[string]interface{} `yaml:"nodeRegistration"`
}

type JoinConfiguration struct {
	NodeRegistration map[string]interface{} `yaml:"nodeRegistration"`
}

type Templater interface {
	KubeadmControlPlaneTemplate(clustername string, namespace string) *KubeadmControlPlaneTemplate
}

type TemplaterFunc func(clustername string, namespace string) *KubeadmControlPlaneTemplate

func (t TemplaterFunc) KubeadmControlPlaneTemplate(clustername string, namespace string) *KubeadmControlPlaneTemplate {
	return t(clustername, namespace)
}

func NewKubeadmControlPlaneTemplate(clustername string, namespace string) *KubeadmControlPlaneTemplate {
	return &KubeadmControlPlaneTemplate{
		APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
		Kind:       "KubeadmControlPlaneTemplate",
		Metadata: Metadata{
			Name:      clustername,
			Namespace: namespace,
		},
		Spec: Spec{
			Template: ControlPlaneTemplate{
				ControlPlaneSpec{
					KubeadmConfigSpec: KubeadmConfigSpec{
						ClusterConfiguration: ClusterConfiguration{
							APIServer: APIServer{
								CertSANs: []string{
									"localhost",
									"127.0.0.1",
									"0.0.0.0",
									"host.docker.internal",
								},
							},
							ControllerManager: ControllerManager{
								ExtraArgs: EnablHostPathProvisioner{
									EnableHostPathProvisioner: "true",
								},
							},
						},
						InitConfiguration: InitConfiguration{
							NodeRegistration: map[string]interface{}{},
						},
						JoinConfiguration: JoinConfiguration{
							NodeRegistration: map[string]interface{}{},
						},
					},
				},
			},
		},
	}
}
