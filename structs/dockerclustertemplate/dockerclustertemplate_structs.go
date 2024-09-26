package dockerclustertemplate

type DockerClusterTemplate struct {
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
	Spec map[string]interface{} `yaml:"spec"`
}
