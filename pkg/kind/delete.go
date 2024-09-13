/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package kind

import "os/exec"

func Delete() {
	cmd := exec.Command("kind", "delete", "clusters", "-A")
	cmd.Run()
}
