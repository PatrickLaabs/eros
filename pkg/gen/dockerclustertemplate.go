package gen

import (
	"github.com/PatrickLaabs/eros/pkg/dockerclustertemplate"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func DockerClusterTemplate() {
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

	err = os.WriteFile("dockerclustertemplate.yaml", yamlData, 0644)
	if err != nil {
		log.Fatalf("error writing file %v", err)
	}
}
