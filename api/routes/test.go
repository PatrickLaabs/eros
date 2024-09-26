package routes

import (
	"github.com/PatrickLaabs/eros/pkg/capd"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	capd.ClusterClass()
	capd.DockerClusterTemplate()
	capd.KubeadmControlPlaneTemplate()
}
