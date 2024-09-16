package kind

type kindConfig struct {
	Kind       string `yaml:"kind"`
	APIVersion string `yaml:"apiVersion"`
	Networking struct {
		IPFamily string `yaml:"ipFamily"`
	} `yaml:"networking"`
	Nodes []struct {
		Role        string `yaml:"role"`
		ExtraMounts []struct {
			HostPath      string `yaml:"hostPath"`
			ContainerPath string `yaml:"containerPath"`
		} `yaml:"extraMounts"`
	} `yaml:"nodes"`
}
