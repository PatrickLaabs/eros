package main

import (
	d "github.com/PatrickLaabs/eros/structs/dockerclustertemplate"
	dt "github.com/PatrickLaabs/eros/structs/dockermachinetemplate"
	k "github.com/PatrickLaabs/eros/structs/kubeadmcontrolplanetemplate"
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

	// DockerClusterTemplate YAML
	dockerClusterTempaltetemplaterFunc := d.TemplaterFunc(d.NewClusterTemplate)

	data := dockerClusterTempaltetemplaterFunc.DockerClusterTemplate(clustername, namespace)
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Printf("yaml.Marshal err %v", err)
	}

	err = os.WriteFile("dockerclustertemplate.yaml", yamlData, 0644)

	// KubeadmControlPlane YAML
	kubeadmTemplaterFunc := k.TemplaterFunc(k.NewKubeadmControlPlaneTemplate)

	data2 := kubeadmTemplaterFunc.KubeadmControlPlaneTemplate(clustername, namespace)
	yamlData2, err := yaml.Marshal(data2)
	if err != nil {
		log.Printf("yaml.Marshal err %v", err)
	}

	err = os.WriteFile("NewKubeadmControlPlaneTemplate.yaml", yamlData2, 0644)

	// DockerMachineTemplate YAML
	dockerMachineTemplateTemplaterFunc := dt.TemplateFunc(dt.NewDockerMachineTemplate)

	data3 := dockerMachineTemplateTemplaterFunc.DockerMachineTemplate(clustername, namespace)
	yamlData3, err := yaml.Marshal(data3)
	if err != nil {
		log.Printf("yaml.Marshal err %v", err)
	}

	err = os.WriteFile("DockerMachineTemplate.yaml", yamlData3, 0644)
}
