/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kind

type kindConfig struct {
	Kind       string     `yaml:"kind"`
	APIVersion string     `yaml:"apiVersion"`
	Networking Networking `yaml:"networking"`
	Nodes      []Nodes    `yaml:"nodes"`
}
type Networking struct {
	IPFamily string `yaml:"ipFamily"`
}
type ExtraMounts struct {
	HostPath      string `yaml:"hostPath"`
	ContainerPath string `yaml:"containerPath"`
}
type Nodes struct {
	Role        string        `yaml:"role"`
	ExtraMounts []ExtraMounts `yaml:"extraMounts"`
}
