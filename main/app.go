package main

import (
	"github.com/easy/lego/support"
	"github.com/BurntSushi/toml"
)

type App struct{
	config *support.Config
	q *support.Connector
}

func (a *App) GetConfig() *support.Config {
	return a.config
}

func (a *App) Init(path string) {
	if _, err := toml.DecodeFile(path, &a.config);
	err != nil {
		panic("Error init config: " + err.Error())
	}
}
/*

GetConfig() *support.Config
Init()
Start()
Stop()*/
