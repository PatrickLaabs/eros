package clusterclass

import (
	"github.com/PatrickLaabs/eros/pkg/gen"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func CapdLocal() {
	data := &gen.ClusterClass{
		APIVersion: "cluster.x-k8s.io/v1beta1",
		Kind:       "ClusterClass",
		Metadata: gen.Metadata{
			Name:      "quick-start",
			Namespace: "default",
		},
		Spec: gen.Spec{
			ControlPlane: gen.ControlPlane{
				MachineInfrastructure: gen.MachineInfrastructure{
					Ref: gen.Ref{
						APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
						Kind:       "DockerMachineTemplate",
						Name:       "quick-start-control-plane",
					},
				},
				Ref: gen.Ref{
					APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
					Kind:       "KubeadmControlPlaneTemplate",
					Name:       "quick-start-control-plane",
				},
			},
			Infrastructure: gen.Infrastructure{
				Ref: gen.Ref{
					APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
					Kind:       "DockerClusterTemplate",
					Name:       "quick-start-cluster",
				},
			},
			Patches: []gen.Patches{
				{
					Definitions: []gen.Definitions{
						{JSONPatches: []gen.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/imageRepository",
								ValueFrom: gen.ValueFrom{
									Variable: "imageRepository",
								},
							},
						},
							Selector: gen.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: gen.MatchResources{
									ControlPlane: true,
								},
							}},
					},
					Description: "Sets the imageRepository used for the KubeadmControlPlane.",
					EnabledIf:   `{{ ne .imageRepository "" }}`,
					Name:        "imageRepository",
				},
				{
					Definitions: []gen.Definitions{
						{JSONPatches: []gen.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/etcd",
								ValueFrom: gen.ValueFrom{
									Template: "local:\n  imageTag: {{ .etcdImageTag }}\n",
								},
							},
						},
							Selector: gen.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: gen.MatchResources{
									ControlPlane: true,
								},
							}},
					},
					Description: "Sets tag to use for the etcd image in the KubeadmControlPlane.",
					Name:        "etcdImageTag",
				},
				{
					Definitions: []gen.Definitions{
						{JSONPatches: []gen.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/dns",
								ValueFrom: gen.ValueFrom{
									Template: "imageTag: {{ .coreDNSImageTag }}\n",
								},
							},
						},
							Selector: gen.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: gen.MatchResources{
									ControlPlane: true,
								},
							},
						},
					},
					Description: "Sets tag to use for the etcd image in the KubeadmControlPlane.",
					Name:        "coreDNSImageTag",
				},
				{
					Definitions: []gen.Definitions{
						{JSONPatches: []gen.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/customImage",
								ValueFrom: gen.ValueFrom{
									Template: "kindest/node:{{ .builtin.machineDeployment.version | replace \"+\" \"_\" }}\n",
								},
							},
						},
							Selector: gen.Selector{
								APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
								Kind:       "DockerMachineTemplate",
								MatchResources: gen.MatchResources{
									MachineDeploymentClass: gen.MachineDeploymentClass{Names: []string{
										"default-worker",
									}},
								},
							},
						},
						{JSONPatches: []gen.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/template/customImage",
								ValueFrom: gen.ValueFrom{
									Template: "kindest/node:{{ .builtin.machinePool.version | replace \"+\" \"_\" }}\n",
								},
							},
						},
							Selector: gen.Selector{
								APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
								Kind:       "DockerMachinePoolTemplate",
								MatchResources: gen.MatchResources{
									MachinePoolClass: gen.MachinePoolClass{Names: []string{
										"default-worker",
									}},
								},
							},
						},
						{JSONPatches: []gen.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/customImage",
								ValueFrom: gen.ValueFrom{
									Template: "kindest/node:{{ .builtin.controlPlane.version | replace \"+\" \"_\" }}\n",
								},
							},
						},
							Selector: gen.Selector{
								APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
								Kind:       "DockerMachineTemplate",
								MatchResources: gen.MatchResources{
									ControlPlane: true,
								},
							},
						},
					},
					Description: "Sets the container image that is used for running dockerMachines for the controlPlane and default-worker machineDeployments.",
					Name:        "customImage",
				},
				{
					Definitions: []gen.Definitions{
						{JSONPatches: []gen.JSONPatches{
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/apiServer/extraArgs",
								Value: gen.MixedValue{
									Single: &gen.AdValue{
										AdmissionControlConfigFile: "/etc/kubernetes/kube-apiserver-admission-pss.yaml",
									},
								},
							},
							{
								Op:   "add",
								Path: "/spec/template/spec/kubeadmConfigSpec/clusterConfiguration/apiServer/extraVolumes",
								Value: gen.MixedValue{
									Multi: []gen.Value{
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
								ValueFrom: gen.ValueFrom{
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
							Selector: gen.Selector{
								APIVersion: "controlplane.cluster.x-k8s.io/v1beta1",
								Kind:       "KubeadmControlPlaneTemplate",
								MatchResources: gen.MatchResources{
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
			Variables: []gen.Variables{
				{
					Name:     "imageRepository",
					Required: true,
					Schema: gen.Schema{
						OpenAPIV3Schema: gen.MixedOpenAPIs{
							OpenAPIV3Schema: &gen.OpenAPIV3Schema{
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
					Schema: gen.Schema{
						OpenAPIV3Schema: gen.MixedOpenAPIs{
							OpenAPIV3Schema: &gen.OpenAPIV3Schema{
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
					Schema: gen.Schema{
						OpenAPIV3Schema: gen.MixedOpenAPIs{
							OpenAPIV3Schema: &gen.OpenAPIV3Schema{
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
					Schema: gen.Schema{
						OpenAPIV3Schema: gen.MixedOpenAPIs{
							OpenAPIV3SchemaNoDefault: &gen.OpenAPIV3SchemaNoDefault{
								Properties: gen.Properties{
									Audit: gen.Audit{
										Default:     "restricted",
										Description: "audit sets the level for the audit PodSecurityConfiguration mode. One of privileged, baseline, restricted.",
										Type:        "string",
									},
									Enabled: gen.Enabled{
										Default:     true,
										Description: "enabled enables the patches to enable Pod Security Standard via AdmissionConfiguration.",
										Type:        "boolean",
									},
									Enforce: gen.Enforce{
										Default:     "baseline",
										Description: "enforce sets the level for the enforce PodSecurityConfiguration mode. One of privileged, baseline, restricted.",
										Type:        "string",
									},
									Warn: gen.Warn{
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
			Workers: gen.Workers{
				MachineDeployments: []gen.MachineDeployments{
					{
						Class: "default-worker",
						Template: gen.Template{
							Bootstrap: gen.Bootstrap{
								Ref: gen.Ref{
									APIVersion: "bootstrap.cluster.x-k8s.io/v1beta1",
									Kind:       "KubeadmConfigTemplate",
									Name:       "quick-start-default-worker-bootstraptemplate",
								},
							},
							Infrastructure: gen.Infrastructure{
								Ref: gen.Ref{
									APIVersion: "infrastructure.cluster.x-k8s.io/v1beta1",
									Kind:       "DockerMachineTemplate",
									Name:       "quick-start-default-worker-machinetemplate",
								},
							},
						},
					},
				},
				MachinePools: []gen.MachinePools{
					{
						Class: "default-worker",
						Template: gen.Template{
							Bootstrap: gen.Bootstrap{
								Ref: gen.Ref{
									APIVersion: "bootstrap.cluster.x-k8s.io/v1beta1",
									Kind:       "KubeadmConfigTemplate",
									Name:       "quick-start-default-worker-bootstraptemplate",
								},
							},
							Infrastructure: gen.Infrastructure{
								Ref: gen.Ref{
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

	err = os.WriteFile("genYaml.yaml", yamlData, 0644)
	if err != nil {
		log.Fatalf("error writing file %v", err)
	}
}
