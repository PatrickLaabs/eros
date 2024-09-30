/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package capd

import (
	dct "github.com/PatrickLaabs/eros/structs/dockerclustertemplate"
	"log"
)

//func DockerClusterTemplate() (yamlData []byte) {
//	data := &dockerclustertemplate.DockerClusterTemplate{
//		APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
//		Kind:       "DockerClusterTemplate",
//		Metadata: dockerclustertemplate.Metadata{
//			Name:      "quick-start-cluster",
//			Namespace: "default",
//		},
//		Spec: dockerclustertemplate.Spec{
//			Template: dockerclustertemplate.Template{
//				Spec: map[string]interface{}{},
//			},
//		},
//	}
//
//	yamlData, err := yaml.Marshal(data)
//	if err != nil {
//		log.Fatalf("error mashaling data %v", err)
//	}
//
//	return yamlData
//}

func DockerClusterTemplate(clustername string, namespace string) (err error) {
	_, err = dct.Render(&dct.DockerClusterTemplate{}, clustername, namespace)
	if err != nil {
		log.Fatalf("failed to render docker-cluster template: %v", err)
		return err
	}
	return err
}
