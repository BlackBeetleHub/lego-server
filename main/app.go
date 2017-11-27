package main

import (
	"github.com/easy/lego/support"
	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)

type App struct {
	config    *support.Config
	connector support.Connector
	router    *httprouter.Router
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
	app.router = httprouter.New()
	app.router.GET("/", app.Index)
	app.router.GET("/get_all_words", app.GetAllWords)
	app.router.POST("/create_account", app.CreateAccount)
}

func (app *App) Start() {
	err := app.connector.Connect()
	if err != nil {
		panic(err.Error())
	}

	handler := cors.Default().Handler(app.router)
	http.ListenAndServe("0.0.0.0:4000", handler)
}

func (app *App) Stop() {
	app.connector.Close()
}
