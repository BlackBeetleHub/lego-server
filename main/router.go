package main

import (
	"fmt"
	"github.com/easy/lego/connection"
	"encoding/json"
	entries "github.com/easy/lego/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *App) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "{ name: 123 }")
}

func (app *App) GetAllWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	email := r.FormValue("email")
	pass := r.FormValue("pass")
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
	fmt.Fprintf(w, string(resJson))
}

func (app *App) CreateAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	/*ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods" ,  "GET, POST, PUT, OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers" ,  "Origin, X-Requested-With, Content-Type, Accept")
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Request.Header.Set("Access-Control-Allow-Methods" ,  "GET, POST, PUT, OPTIONS")
	ctx.Request.Header.Set("Access-Control-Allow-Headers" ,  "Origin, X-Requested-With, Content-Type, Accept")
	*///
	//pass:= ctx.QueryArgs().Peek("pass")
	//ctx.Response.Header.Set("Allow", "GET, HEAD, OPTIONS")
	strDetails := r.FormValue("details")
	strPass := r.FormValue("pass")

	println(string(strPass))

	str := string(strDetails)

	println(str)


	var details entries.Details

	println(strDetails + " empty?")

	err := json.Unmarshal([]byte(strDetails), &details)
	if err != nil {
		fmt.Fprintf(w, string(err.Error()))
		return;
	}
	fmt.Fprintf(w, "200 OK HUI TAM PLAVAL")
	/*rawDetails := string(strDetails)
	// add in database and return 200 OK

	strInsertQ := "INSERT INTO account (id, details, encoded_password) VALUES (''," + rawDetails + "'');"

	app.connector.Insert(strInsertQ)*/

}