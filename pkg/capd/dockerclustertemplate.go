/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package capd

import (
	d "github.com/PatrickLaabs/eros/structs/dockerclustertemplate"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func DockerClusterTemplate(clustername string, namespace string) (err error) {
	templaterFunc := d.TemplaterFunc(d.NewClusterTemplate)
	data := templaterFunc.DockerClusterTemplate(clustername, namespace)

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Printf("error marshalling template data: %v", err)
	}

	err = os.WriteFile("test.yaml", yamlData, 0644)
	return err
}
