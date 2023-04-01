package http

import (
	"net/http"
	"github.com/flystary/aiops/modules/transfer/tools/proc"
)

func configProcHttpRoutes() {
	http.HandleFunc("/counter/all", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, proc.GetAll())
	})

}