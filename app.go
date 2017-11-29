package main

import (
	"bike-api/config"
	"bike-api/controller/rackcontroller"
	"bike-api/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Enable line numbers for logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	setPackageConfiguration(config.Init())

	var rc rackcontroller.RackController

	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/racks", rc)

	http.ListenAndServe(":5000", nil)

}

func setPackageConfiguration(c config.Config) {
	model.DBType = c.Database.Type
	model.ConnString = c.Database.ConnString
}
