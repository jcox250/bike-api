package main

import (
	"bike-api/config"
	"bike-api/model"
)

func main() {
	setPackageConfiguration(config.Init())

}

func setPackageConfiguration(c config.Config) {
	model.DBType = c.Database.Type
	model.ConnString = c.Database.ConnString
}
