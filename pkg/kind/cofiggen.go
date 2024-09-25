/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kind

import (
	"gopkg.in/yaml.v3"
	"log"
)

func configgen() []byte {
	kindConfigYaml := &kindConfig{
		Kind:       "Cluster",
		APIVersion: "kind.x-k8s.io/v1alpha4",
		Networking: Networking{
			IPFamily: "dual",
		},
		Nodes: []Nodes{
			{
				Role: "control-plane",
				ExtraMounts: []ExtraMounts{
					{
						HostPath:      "/var/run/docker.sock",
						ContainerPath: "/var/run/docker.sock",
					},
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(kindConfigYaml)
	// ToDo: Testing err handling is still a thing
	if err != nil {
		log.Printf("error marshalling kind config: %v", err)
	}

	return yamlData
}
