package capd

import (
	"github.com/PatrickLaabs/eros/structs/dockerclustertemplate"
	"gopkg.in/yaml.v2"
	"log"
)

func DockerClusterTemplate() (yamlData []byte) {
	data := &dockerclustertemplate.DockerClusterTemplate{
		APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
		Kind:       "DockerClusterTemplate",
		Metadata: dockerclustertemplate.Metadata{
			Name:      "quick-start-cluster",
			Namespace: "default",
		},
		Spec: dockerclustertemplate.Spec{
			Template: dockerclustertemplate.Template{
				Spec: map[string]interface{}{},
			},
		},
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("error mashaling data %v", err)
	}

	return yamlData
}
