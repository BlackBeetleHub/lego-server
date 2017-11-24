package main

import "testing"

func TestApp_GetConfig(t *testing.T) {
	var app App
	app.Init("../config.toml")

	config:=app.GetConfig()

	println(config.Database)
}
