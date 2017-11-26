package main

import (
	"github.com/easy/lego/support"
	"github.com/BurntSushi/toml"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

type App struct {
	config    *support.Config
	connector support.Connector
	router    *fasthttprouter.Router
}

func (app *App) GetConfig() *support.Config {
	return app.config
}

func (app *App) Init(path string) {
	_, err := toml.DecodeFile(path, &app.config);
	if err != nil {
		panic("Error init config: " + err.Error())
	}
	app.connector = &support.PostgresConnector{ConnectorString: app.config.GetDBStringConnector()}
	app.router = fasthttprouter.New()

	app.router.GET("/", app.Index)
	app.router.GET("/get_all_words", app.GetAllWords)
	app.router.POST("/create_account", app.CreateAccount)

}

func (app *App) Start() {
	err := app.connector.Connect()
	if err != nil {
		panic(err.Error())
	}
	address := app.config.ListenIP + ":" + app.config.ListenPort;
	log.Fatal(fasthttp.ListenAndServe(address, app.router.Handler))

}

func (app *App) Stop() {
	app.connector.Close()
}
