/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package dockerclustertemplate

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

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

type Templater interface {
	Dct(clustername string, namespace string) *DockerClusterTemplate
}

func (d *DockerClusterTemplate) Dct(clustername string, namespace string) *DockerClusterTemplate {
	data := &DockerClusterTemplate{
		APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
		Kind:       "DockerClusterTemplate",
		Metadata: Metadata{
			Name:      clustername,
			Namespace: namespace,
		},
		Spec: Spec{
			Template: Template{
				Spec: map[string]interface{}{},
			},
		},
	}
	return data
}

func Render(t Templater, clustername string, namespace string) ([]byte, error) {
	ns := t.Dct(clustername, namespace)
	yamlData, err := yaml.Marshal(ns)
	if err != nil {
		log.Fatalf("error mashaling data %v", err)
	}

	err = os.WriteFile("test.yaml", yamlData, 0644)
	if err != nil {
		return nil, nil
	}

	return yamlData, nil
}
