package main

import (
	d "github.com/PatrickLaabs/eros/structs/dockerclustertemplate"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

/*
Testing interfaces
*/

func main() {
	namespace := "default"
	clustername := "eros-mgmt-cluster"

	templaterFunc := d.TemplaterFunc(d.NewClusterTemplate)

	data := templaterFunc.DockerClusterTemplate(clustername, namespace)
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Printf("yaml.Marshal err %v", err)
	}

	err = os.WriteFile("test.yaml", yamlData, 0644)
}
