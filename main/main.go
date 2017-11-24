package main

import (
	"github.com/easy/lego/connection"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"fmt"
	"encoding/json"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(ctx, "{ name: 123 }")
}

func getAllWords(ctx *fasthttp.RequestCtx) {
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

func main(){
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/get_all_words", getAllWords)

	log.Fatal(fasthttp.ListenAndServe("0.0.0.0:4000", router.Handler))
	return
}