package rackcontroller

import (
	"bike-api/model/rack"
	"bike-api/util/httputil"
	"encoding/json"
	"net/http"
)

type RackController rack.Rack

func (rc RackController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		rc.get(w, r)
	} else {
		httputil.NotImplemented(w)
	}
}

func (rc RackController) get(w http.ResponseWriter, r *http.Request) {
	racks, err := rack.GetAll()
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}

	data, err := json.Marshal(racks)
	if err != nil {
		httputil.InternalServerError(w, err)
		return
	}
	httputil.WriteResponse(w, data)
}
