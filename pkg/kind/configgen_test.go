/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kind

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

func TestConfiggen(t *testing.T) {

	// Generate the YAML from the configgen function
	yamlData := configgen()

	// Define the expected fields to check in the output
	var config kindConfig
	err := yaml.Unmarshal(yamlData, &config)
	if err != nil {
		t.Fatalf("Failed to unmarshal yaml: %v", err)
	}

	// Check if the generated YAML has the expected values
	assert.Equal(t, "Cluster", config.Kind, "Expected kind to be 'Cluster'")
	assert.Equal(t, "kind.x-k8s.io/v1alpha4", config.APIVersion, "Expected apiVersion to be 'kind.x-k8s.io/v1alpha4'")
	assert.Equal(t, "dual", config.Networking.IPFamily, "Expected ipFamily to be 'dual'")

	// Ensure that there is one node and it's a control-plane with the expected mounts
	assert.Equal(t, 1, len(config.Nodes), "Expected exactly one node")
	assert.Equal(t, "control-plane", config.Nodes[0].Role, "Expected node role to be 'control-plane'")
	assert.Equal(t, 1, len(config.Nodes[0].ExtraMounts), "Expected exactly one extraMount")
	assert.Equal(t, "/var/run/docker.sock", config.Nodes[0].ExtraMounts[0].HostPath, "Expected HostPath to be '/var/run/docker.sock'")
	assert.Equal(t, "/var/run/docker.sock", config.Nodes[0].ExtraMounts[0].ContainerPath, "Expected ContainerPath to be '/var/run/docker.sock'")

	// Additional check to ensure YAML contains some specific strings
	assert.True(t, strings.Contains(string(yamlData), "kind: Cluster"), "Expected yaml to contain 'kind: Cluster'")
	assert.True(t, strings.Contains(string(yamlData), "apiVersion: kind.x-k8s.io/v1alpha4"), "Expected yaml to contain 'apiVersion: kind.x-k8s.io/v1alpha4'")
}
