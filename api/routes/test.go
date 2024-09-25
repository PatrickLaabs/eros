package routes

import (
	"github.com/PatrickLaabs/eros/pkg/gen/capd/clusterclass"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	clusterclass.CapdLocal()
}
