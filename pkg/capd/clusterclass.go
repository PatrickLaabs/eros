/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package capd

import (
	"github.com/PatrickLaabs/eros/structs/clusterclass"
	"gopkg.in/yaml.v2"
	"log"
)

func ClusterClass() (yamlData []byte) {
	data := &clusterclass.ClusterClass{
		APIVersion: "cluster.x-k8s.io/v1beta1",
		Kind:       "ClusterClass",
		Metadata: clusterclass.Metadata{
			Name:      "quick-start",
			Namespace: "default",
		},
		Spec: clusterclass.Spec{
			ControlPlane: clusterclass.ControlPlane{
				MachineInfrastructure: clusterclass.MachineInfrastructure{
					Ref: clusterclass.Ref{
						APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
						Kind:       "DockerMachineTemplate",
						Name:       "quick-start-control-plane",
					},
				},
				Ref: clusterclass.Ref{
					APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
					Kind:       "KubeadmControlPlaneTemplate",
					Name:       "quick-start-control-plane",
				},
			},
			Infrastructure: clusterclass.Infrastructure{
				Ref: clusterclass.Ref{
					APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
					Kind:       "DockerClusterTemplate",
					Name:       "quick-start-cluster",
				},
			},
			Patches: []clusterclass.Patches{
				{
					Definitions: []clusterclass.Definitions{
						{JSONPatches: []clusterclass.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/imageRepository",
								ValueFrom: clusterclass.ValueFrom{
									Variable: "imageRepository",
								},
							},
						},
							Selector: clusterclass.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: clusterclass.MatchResources{
									ControlPlane: true,
								},
							}},
					},
					Description: "Sets the imageRepository used for the KubeadmControlPlane.",
					EnabledIf:   `{{ ne .imageRepository "" }}`,
					Name:        "imageRepository",
				},
				{
					Definitions: []clusterclass.Definitions{
						{JSONPatches: []clusterclass.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/etcd",
								ValueFrom: clusterclass.ValueFrom{
									Template: "local:\n  imageTag: {{ .etcdImageTag }}\n",
								},
							},
						},
							Selector: clusterclass.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: clusterclass.MatchResources{
									ControlPlane: true,
								},
							}},
					},
					Description: "Sets tag to use for the etcd image in the KubeadmControlPlane.",
					Name:        "etcdImageTag",
				},
				{
					Definitions: []clusterclass.Definitions{
						{JSONPatches: []clusterclass.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/dns",
								ValueFrom: clusterclass.ValueFrom{
									Template: "imageTag: {{ .coreDNSImageTag }}\n",
								},
							},
						},
							Selector: clusterclass.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: clusterclass.MatchResources{
									ControlPlane: true,
								},
							},
						},
					},
					Description: "Sets tag to use for the etcd image in the KubeadmControlPlane.",
					Name:        "coreDNSImageTag",
				},
				{
					Definitions: []clusterclass.Definitions{
						{JSONPatches: []clusterclass.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/customImage",
								ValueFrom: clusterclass.ValueFrom{
									Template: "kindest/node:{{ .builtin.machineDeployment.version | replace \"+\" \"_\" }}\n",
								},
							},
						},
							Selector: clusterclass.Selector{
								APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
								Kind:       "DockerMachineTemplate",
								MatchResources: clusterclass.MatchResources{
									MachineDeploymentClass: clusterclass.MachineDeploymentClass{Names: []string{
										"default-worker",
									}},
								},
							},
						},
						{JSONPatches: []clusterclass.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/template/customImage",
								ValueFrom: clusterclass.ValueFrom{
									Template: "kindest/node:{{ .builtin.machinePool.version | replace \"+\" \"_\" }}\n",
								},
							},
						},
							Selector: clusterclass.Selector{
								APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
								Kind:       "DockerMachinePoolTemplate",
								MatchResources: clusterclass.MatchResources{
									MachinePoolClass: clusterclass.MachinePoolClass{Names: []string{
										"default-worker",
									}},
								},
							},
						},
						{JSONPatches: []clusterclass.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/customImage",
								ValueFrom: clusterclass.ValueFrom{
									Template: "kindest/node:{{ .builtin.controlPlane.version | replace \"+\" \"_\" }}\n",
								},
							},
						},
							Selector: clusterclass.Selector{
								APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
								Kind:       "DockerMachineTemplate",
								MatchResources: clusterclass.MatchResources{
									ControlPlane: true,
								},
							},
						},
					},
					Description: "Sets the container image that is used for running dockerMachines for the controlPlane and default-worker machineDeployments.",
					Name:        "customImage",
				},
				{
					Definitions: []clusterclass.Definitions{
						{JSONPatches: []clusterclass.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/apiServer/extraArgs",
								Value: clusterclass.MixedValue{
									Single: &clusterclass.AdValue{
										AdmissionControlConfigFile: "/etc/kubernetes/kube-apiserver-admission-pss.yaml",
									},
								},
							},
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/apiServer/extraVolumes",
								Value: clusterclass.MixedValue{
									Multi: []clusterclass.Value{
										{
											HostPath:  "/etc/kubernetes/kube-apiserver-admission-pss.yaml",
											MountPath: "/etc/kubernetes/kube-apiserver-admission-pss.yaml",
											Name:      "admission-pss",
											PathType:  "File",
											ReadOnly:  true,
										},
									},
								},
							},
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/files",
								ValueFrom: clusterclass.ValueFrom{
									Template: `- content: |
    apiVersion: apiserver.config.k8s.io/v1
    kind: AdmissionConfiguration
    plugins:
    - name: PodSecurity
      configuration:
        apiVersion: pod-security.admission.config.k8s.io/v1{{ if semverCompare "< v1.25" .builtin.controlPlane.version }}beta1{{ end }}
        kind: PodSecurityConfiguration
        defaults:
          enforce: "{{ .podSecurityStandard.enforce }}"
          enforce-version: "latest"
          audit: "{{ .podSecurityStandard.audit }}"
          audit-version: "latest"
          warn: "{{ .podSecurityStandard.warn }}"
          warn-version: "latest"
        exemptions:
          usernames: []
          runtimeClasses: []
          namespaces: [kube-system]
  path: /etc/kubernetes/kube-apiserver-admission-pss.yaml
`,
								},
							},
						},
							Selector: clusterclass.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: clusterclass.MatchResources{
									ControlPlane: true,
								},
							},
						},
					},
					Description: "Adds an admission configuration for PodSecurity to the kube-apiserver.",
					EnabledIf:   `{{ .podSecurityStandard.enabled }}`,
					Name:        "podSecurityStandard",
				},
			},
			Variables: []clusterclass.Variables{
				{
					Name:     "imageRepository",
					Required: true,
					Schema: clusterclass.Schema{
						OpenAPIV3Schema: clusterclass.MixedOpenAPIs{
							OpenAPIV3Schema: &clusterclass.OpenAPIV3Schema{
								Default:     "",
								Description: "imageRepository sets the container registry to pull images from. If empty, nothing will be set and the from of kubeadm will be used.",
								Example:     "registry.k8s.io",
								Type:        "string",
							},
						},
					},
				},
				{
					Name:     "etcdImageTag",
					Required: true,
					Schema: clusterclass.Schema{
						OpenAPIV3Schema: clusterclass.MixedOpenAPIs{
							OpenAPIV3Schema: &clusterclass.OpenAPIV3Schema{
								Default:     "",
								Description: "etcdImageTag sets the tag for the etcd image.",
								Example:     "3.5.3-0",
								Type:        "string",
							},
						},
					},
				},
				{
					Name:     "coreDNSImageTag",
					Required: true,
					Schema: clusterclass.Schema{
						OpenAPIV3Schema: clusterclass.MixedOpenAPIs{
							OpenAPIV3Schema: &clusterclass.OpenAPIV3Schema{
								Default:     "",
								Description: "coreDNSImageTag sets the tag for the coreDNS image.",
								Example:     "v1.8.5",
								Type:        "string",
							},
						},
					},
				},
				{
					Name:     "podSecurityStandard",
					Required: false,
					Schema: clusterclass.Schema{
						OpenAPIV3Schema: clusterclass.MixedOpenAPIs{
							OpenAPIV3SchemaNoDefault: &clusterclass.OpenAPIV3SchemaNoDefault{
								Properties: clusterclass.Properties{
									Audit: clusterclass.Audit{
										Default:     "restricted",
										Description: "audit sets the level for the audit PodSecurityConfiguration mode. One of privileged, baseline, restricted.",
										Type:        "string",
									},
									Enabled: clusterclass.Enabled{
										Default:     true,
										Description: "enabled enables the patches to enable Pod Security Standard via AdmissionConfiguration.",
										Type:        "boolean",
									},
									Enforce: clusterclass.Enforce{
										Default:     "baseline",
										Description: "enforce sets the level for the enforce PodSecurityConfiguration mode. One of privileged, baseline, restricted.",
										Type:        "string",
									},
									Warn: clusterclass.Warn{
										Default:     "restricted",
										Description: "warn sets the level for the warn PodSecurityConfiguration mode. One of privileged, baseline, restricted.",
										Type:        "string",
									},
								},
								Type: "object",
							},
						},
					},
				},
			},
			Workers: clusterclass.Workers{
				MachineDeployments: []clusterclass.MachineDeployments{
					{
						Class: "default-worker",
						Template: clusterclass.Template{
							Bootstrap: clusterclass.Bootstrap{
								Ref: clusterclass.Ref{
									APIVersion: "bootstrap.cluster.x-k8s.io/v1beta1",
									Kind:       "KubeadmConfigTemplate",
									Name:       "quick-start-default-worker-bootstraptemplate",
								},
							},
							Infrastructure: clusterclass.Infrastructure{
								Ref: clusterclass.Ref{
									APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
									Kind:       "DockerMachineTemplate",
									Name:       "quick-start-default-worker-machinetemplate",
								},
							},
						},
					},
				},
				MachinePools: []clusterclass.MachinePools{
					{
						Class: "default-worker",
						Template: clusterclass.Template{
							Bootstrap: clusterclass.Bootstrap{
								Ref: clusterclass.Ref{
									APIVersion: "bootstrap.cluster.x-k8s.io/v1beta1",
									Kind:       "KubeadmConfigTemplate",
									Name:       "quick-start-default-worker-bootstraptemplate",
								},
							},
							Infrastructure: clusterclass.Infrastructure{
								Ref: clusterclass.Ref{
									APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
									Kind:       "DockerMachinePoolTemplate",
									Name:       "quick-start-default-worker-machinepooltemplate",
								},
							},
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
