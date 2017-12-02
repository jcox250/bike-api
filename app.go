package main

import (
	"bike-api/config"
	"bike-api/controller/rackcontroller"
	"bike-api/model"
	"bike-api/util/httputil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	setConfigVariables(config.Init())

	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var rc rackcontroller.RackController

	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/racks", rc)

	c := httputil.NewCors()
	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":5000", httputil.Log(handler)))
}

func setConfigVariables(c config.Config) {
	model.DBType = c.Database.Type
	model.ConnString = c.Database.ConnString
	httputil.ApiKey = c.ApiKey
	httputil.AllowedOrigins = c.Cors.AllowedOrigins
	httputil.AllowedHeaders = c.Cors.AllowedHeaders
	httputil.AllowedMethods = c.Cors.AllowedMethods
}
