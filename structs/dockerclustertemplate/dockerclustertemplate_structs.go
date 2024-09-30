/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

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

type Templater interface {
	DockerClusterTemplate(clustername string, namespace string) *DockerClusterTemplate
}

type TemplaterFunc func(clustername string, namespace string) *DockerClusterTemplate

func (t TemplaterFunc) DockerClusterTemplate(clustername string, namespace string) *DockerClusterTemplate {
	return t(clustername, namespace)
}

func NewClusterTemplate(clustername string, namespace string) *DockerClusterTemplate {
	return &DockerClusterTemplate{
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
}
