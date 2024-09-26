package clusterclass

type ClusterClass struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}
type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}
type Ref struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Name       string `yaml:"name"`
}
type MachineInfrastructure struct {
	Ref Ref `yaml:"ref"`
}
type ControlPlane struct {
	MachineInfrastructure MachineInfrastructure `yaml:"machineInfrastructure"`
	Ref                   Ref                   `yaml:"ref"`
}
type Infrastructure struct {
	Ref Ref `yaml:"ref"`
}

// MixedValue Custom type to handle both single and multiple values
type MixedValue struct {
	Single *AdValue // For single-value case
	Multi  []Value  // For multi-value case
}
type ValueFrom struct {
	Variable string `yaml:"variable,omitempty"`
	Template string `yaml:"template,omitempty"`
}
type JSONPatches struct {
	Op        string     `yaml:"op"`
	Path      string     `yaml:"path"`
	ValueFrom ValueFrom  `yaml:"valueFrom,omitempty"`
	Value     MixedValue `yaml:"value,omitempty"`
}
type AdValue struct {
	AdmissionControlConfigFile string `yaml:"admission-control-config-file,omitempty"`
}
type Value struct {
	AdmissionControlConfigFile string `yaml:"admission-control-config-file,omitempty"`
	HostPath                   string `yaml:"hostPath,omitempty"`
	MountPath                  string `yaml:"mountPath,omitempty"`
	Name                       string `yaml:"name,omitempty"`
	PathType                   string `yaml:"pathType,omitempty"`
	ReadOnly                   bool   `yaml:"readOnly,omitempty"`
}
type MatchResources struct {
	ControlPlane           bool                   `yaml:"controlPlane,omitempty"`
	MachineDeploymentClass MachineDeploymentClass `yaml:"machineDeploymentClass,omitempty"`
	MachinePoolClass       MachinePoolClass       `yaml:"machinePoolClass,omitempty"`
}

type MachineDeploymentClass struct {
	Names []string `yaml:"names"`
}

type MachinePoolClass struct {
	Names []string `yaml:"names"`
}

type Selector struct {
	APIVersion     string         `yaml:"apiVersion"`
	Kind           string         `yaml:"kind"`
	MatchResources MatchResources `yaml:"matchResources"`
}
type Definitions struct {
	JSONPatches []JSONPatches `yaml:"jsonPatches"`
	Selector    Selector      `yaml:"selector"`
}
type Patches struct {
	Definitions []Definitions `yaml:"definitions"`
	Description string        `yaml:"description"`
	EnabledIf   string        `yaml:"enabledIf,omitempty"`
	Name        string        `yaml:"name"`
}

type OpenAPIV3Schema struct {
	Default     string     `yaml:"default"`
	Description string     `yaml:"description,omitempty"`
	Example     string     `yaml:"example,omitempty"`
	Type        string     `yaml:"type,omitempty"`
	Properties  Properties `yaml:"properties,omitempty"`
}

type OpenAPIV3SchemaNoDefault struct {
	Description string     `yaml:"description,omitempty"`
	Example     string     `yaml:"example,omitempty"`
	Type        string     `yaml:"type,omitempty"`
	Properties  Properties `yaml:"properties,omitempty"`
}

type MixedOpenAPIs struct {
	OpenAPIV3Schema          *OpenAPIV3Schema
	OpenAPIV3SchemaNoDefault *OpenAPIV3SchemaNoDefault
}

type Properties struct {
	Audit   Audit   `yaml:"audit"`
	Enabled Enabled `yaml:"enabled"`
	Enforce Enforce `yaml:"enforce"`
	Warn    Warn    `yaml:"warn"`
}

type Audit struct {
	Default     string `yaml:"default,omitempty"`
	Description string `yaml:"description,omitempty"`
	Type        string `yaml:"type,omitempty"`
}

type Enabled struct {
	Default     bool   `yaml:"default,omitempty"`
	Description string `yaml:"description,omitempty"`
	Type        string `yaml:"type,omitempty"`
}

type Enforce struct {
	Default     string `yaml:"default,omitempty"`
	Description string `yaml:"description,omitempty"`
	Type        string `yaml:"type,omitempty"`
}

type Warn struct {
	Default     string `yaml:"default,omitempty"`
	Description string `yaml:"description,omitempty"`
	Type        string `yaml:"type,omitempty"`
}

type Schema struct {
	OpenAPIV3Schema MixedOpenAPIs `yaml:"openAPIV3Schema,omitempty"`
}

type Variables struct {
	Name     string `yaml:"name"`
	Required bool   `yaml:"required"`
	Schema   Schema `yaml:"schema"`
}
type Bootstrap struct {
	Ref Ref `yaml:"ref"`
}
type Template struct {
	Bootstrap      Bootstrap      `yaml:"bootstrap"`
	Infrastructure Infrastructure `yaml:"infrastructure"`
}
type MachineDeployments struct {
	Class    string   `yaml:"class"`
	Template Template `yaml:"template"`
}
type MachinePools struct {
	Class    string   `yaml:"class,omitempty"`
	Template Template `yaml:"template,omitempty"`
}
type Workers struct {
	MachineDeployments []MachineDeployments `yaml:"machineDeployments,omitempty"`
	MachinePools       []MachinePools       `yaml:"machinePools,omitempty"`
}
type Spec struct {
	ControlPlane   ControlPlane   `yaml:"controlPlane,omitempty"`
	Infrastructure Infrastructure `yaml:"infrastructure,omitempty"`
	Patches        []Patches      `yaml:"patches,omitempty"`
	Variables      []Variables    `yaml:"variables,omitempty"`
	Workers        Workers        `yaml:"workers,omitempty"`
}

func (mo MixedOpenAPIs) MarshalYAML() (interface{}, error) {
	if mo.OpenAPIV3Schema != nil {
		return mo.OpenAPIV3Schema, nil
	} else if mo.OpenAPIV3SchemaNoDefault != nil {
		return mo.OpenAPIV3SchemaNoDefault, nil
	}
	return nil, nil
}

// MarshalYAML Custom YAML marshaling function for MixedValue
func (mv MixedValue) MarshalYAML() (interface{}, error) {
	if mv.Single != nil {
		// Marshal as a single object
		return mv.Single, nil
	} else if len(mv.Multi) > 0 {
		// Marshal as an array of objects
		return mv.Multi, nil
	}
	return nil, nil
}
