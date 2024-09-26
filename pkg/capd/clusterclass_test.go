package capd

import (
	"github.com/google/go-cmp/cmp"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"testing"
)

// TestCapdLocal tests the CapdLocal function by comparing the generated YAML with the wanted YAML.
func TestCapdLocal(t *testing.T) {
	// Run the function that generates the YAML file
	p := ClusterClass()

	// Load the wanted YAML file (which should be present in a specific directory in your project)
	wantedYamlPath := filepath.Join("../../testdata", "clusterclass.yaml")
	wantedYamlData, err := os.ReadFile(wantedYamlPath)
	if err != nil {
		t.Fatalf("failed to read wanted YAML file: %v", err)
	}

	// Unmarshal both YAMLs to make comparison more structured
	var genYamlObject, wantedYamlObject interface{}

	if err := yaml.Unmarshal(p, &genYamlObject); err != nil {
		t.Fatalf("failed to unmarshal generated YAML: %v", err)
	}

	if err := yaml.Unmarshal(wantedYamlData, &wantedYamlObject); err != nil {
		t.Fatalf("failed to unmarshal wanted YAML: %v", err)
	}

	// Compare both YAMLs
	if diff := cmp.Diff(wantedYamlObject, genYamlObject); diff != "" {
		t.Errorf("generated YAML does not match the wanted YAML. Diff:\n%s", diff)
	}
}
