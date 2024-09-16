package kind

import (
	"gopkg.in/yaml.v3"
	"log"
)

func Configgen() []byte {
	kindConfigYaml := &kindConfig{
		Kind:       "Cluster",
		APIVersion: "kind.x-k8s.io/v1alpha4",
		Networking: struct {
			IPFamily string `yaml:"ipFamily"`
		}{IPFamily: "dual"},
		Nodes: []struct {
			Role        string `yaml:"role"`
			ExtraMounts []struct {
				HostPath      string `yaml:"hostPath"`
				ContainerPath string `yaml:"containerPath"`
			} `yaml:"extraMounts"`
		}{
			{
				Role: "control-plane",
				ExtraMounts: []struct {
					HostPath      string `yaml:"hostPath"`
					ContainerPath string `yaml:"containerPath"`
				}([]struct {
					HostPath      string
					ContainerPath string
				}{
					{
						HostPath:      "/var/run/docker.sock",
						ContainerPath: "/var/run/docker.sock",
					},
				}),
			},
		},
	}

	yamlData, err := yaml.Marshal(kindConfigYaml)
	if err != nil {
		log.Printf("error marshalling kind config: %v", err)
	}

	return yamlData
}
