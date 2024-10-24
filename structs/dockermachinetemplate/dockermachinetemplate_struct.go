package dockermachinetemplate

type DockerMachineTemplate struct {
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
	Template Template `yaml:"template"`
}

type Template struct {
	Spec TemplateSpec `yaml:"spec"`
}

type TemplateSpec struct {
	ExtraMounts []ExtraMounts `yaml:"extraMounts"`
}

type ExtraMounts struct {
	ContainerPath string `yaml:"containerPath"`
	HostPath      string `yaml:"hostPath"`
}

type Templater interface {
	DockerMachineTemplate(clustername string, namespace string) *DockerMachineTemplate
}

type TemplateFunc func(clustername string, namespace string) *DockerMachineTemplate

func (t TemplateFunc) DockerMachineTemplate(clustername string, namespace string) *DockerMachineTemplate {
	return t(clustername, namespace)
}

func NewDockerMachineTemplate(clustername string, namespace string) *DockerMachineTemplate {
	return &DockerMachineTemplate{
		APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
		Kind:       "DockerMachineTemplate",
		Metadata: Metadata{
			Name:      clustername,
			Namespace: namespace,
		},
		Spec: Spec{
			Template: Template{
				Spec: TemplateSpec{
					ExtraMounts: []ExtraMounts{
						{
							ContainerPath: "/var/run/docker.sock",
							HostPath:      "/var/run/docker.sock",
						},
					},
				},
			},
		},
	}
}
