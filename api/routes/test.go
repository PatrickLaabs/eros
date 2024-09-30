/*
Copyright © 2024 Patrick Laabs patrick.laabs@me.com
*/

package routes

import (
	"github.com/PatrickLaabs/eros/pkg/capd"
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	namespace := "default"
	clustername := "eros-mgmt-cluster"

	capd.ClusterClass()
	err := capd.DockerClusterTemplate(clustername, namespace)
	if err != nil {
		log.Fatalf(err.Error())
	}
	capd.KubeadmControlPlaneTemplate()
}
