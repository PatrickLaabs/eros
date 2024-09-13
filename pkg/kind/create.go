/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kind

import (
	"fmt"
	kind "sigs.k8s.io/kind/pkg/cluster"
)

func Create() {
	p := kind.NewProvider()

	if err := p.Create(
		"test",
		kind.CreateWithConfigFile("./kind-config.yaml"),
	); err != nil {
		fmt.Printf("error creating kind cluster: %v", err)
	}
}
