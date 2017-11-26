package main

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"github.com/easy/lego/connection"
	"encoding/json"
	entries "github.com/easy/lego/json"
)

func (app *App) Index(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(ctx, "{ name: 123 }")
}

func (app *App) GetAllWords(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	email := ctx.QueryArgs().Peek("email")
	pass := ctx.QueryArgs().Peek("pass")
	sp := connection.SimpleConnector{Login: string(email), Pass: string(pass)}
	err := sp.Connect()
	if err != nil {
		println(err.Error())
	}
	dict := sp.GetAllWords()
	resJson, err := json.Marshal(&dict)
	if err != nil {
		panic("marshal error")
	}
	fmt.Fprintf(ctx, string(resJson))
}

func (app *App) CreateAccount(ctx *fasthttp.RequestCtx) {
	//pass:= ctx.QueryArgs().Peek("pass")
	strDetails := ctx.PostArgs().Peek("details")
	var details entries.Details

	err := json.Unmarshal(strDetails, &details)
	if err != nil {
		fmt.Fprintf(ctx, string(err.Error()))
		return;
	}

	rawDetails := string(strDetails)
	// add in database and return 200 OK

	strInsertQ := "INSERT INTO account (id, details, encoded_password) VALUES (''," + rawDetails + "'');"

	app.connector.Insert(strInsertQ)

}