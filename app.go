package main

import "bike-api/config"

var Config config.Config

func main() {
	Config = config.Init()

}
