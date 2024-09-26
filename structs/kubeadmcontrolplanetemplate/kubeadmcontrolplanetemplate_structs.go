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
	ExtraArgs map[string]string `yaml:"extraArgs"`
}

type InitConfiguration struct {
	NodeRegistration map[string]interface{} `yaml:"nodeRegistration"`
}

type JoinConfiguration struct {
	NodeRegistration map[string]interface{} `yaml:"nodeRegistration"`
}
