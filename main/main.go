package main

import (
	"os"
)

func main() {
	app := App{}
	pathConfig := os.Args
	println(pathConfig[1])
	app.Init(pathConfig[1])
	app.Start()
	return
}
