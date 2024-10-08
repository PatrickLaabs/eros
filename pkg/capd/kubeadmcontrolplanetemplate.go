package capd

import (
	k "github.com/PatrickLaabs/eros/structs/kubeadmcontrolplanetemplate"
	"gopkg.in/yaml.v2"
	"log"
)

func KubeadmControlPlaneTemplate() (yamlData []byte) {
	data := &k.KubeadmControlPlaneTemplate{
		APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
		Kind:       "KubeadmControlPlaneTemplate",
		Metadata: k.Metadata{
			Name:      "quick-start-control-plane",
			Namespace: "default",
		},
		Spec: k.Spec{
			Template: k.ControlPlaneTemplate{
				Spec: k.ControlPlaneSpec{
					KubeadmConfigSpec: k.KubeadmConfigSpec{
						ClusterConfiguration: k.ClusterConfiguration{
							APIServer: k.APIServer{
								CertSANs: []string{
									"localhost",
									"127.0.0.1",
									"0.0.0.0",
									"host.docker.internal",
								},
							},
							ControllerManager: k.ControllerManager{
								ExtraArgs: map[string]string{
									"enable-hostpath-provisioner": "true",
								},
							},
						},

						InitConfiguration: k.InitConfiguration{
							NodeRegistration: map[string]interface{}{},
						},
						JoinConfiguration: k.JoinConfiguration{
							NodeRegistration: map[string]interface{}{},
						},
					},
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Fatalf("error mashaling data %v", err)
	}

	return yamlData
}
