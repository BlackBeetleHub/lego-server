package main

import (
	"fmt"
	"github.com/easy/lego/connection"
	"encoding/json"
	entries "github.com/easy/lego/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *App) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "{ name: successful response  }")
}

func (app *App) GetAllWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	email := r.Form.Get("email")
	pass := r.Form.Get("pass")
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
	strPass := r.FormValue("pass")
	strDetails := r.FormValue("details")

	var details entries.Details
	err := json.Unmarshal([]byte(strDetails), &details)
	if err != nil {
		fmt.Fprintf(w, string(err.Error()))
		return;
	}

	fmt.Fprintf(w, "200 OK")
	strInsertQ := "INSERT INTO account (details, hash) VALUES ('" + strDetails + "','" + strPass + "');"
	app.connector.Insert(strInsertQ)
}

func (app *App) AddCustomWord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	strPass := r.FormValue("value")
	strDetails := r.FormValue("details")
	var details entries.Details
	err := json.Unmarshal([]byte(strDetails), &details)
	if err != nil {
		fmt.Fprintf(w, string(err.Error()))
		return;
	}

	fmt.Fprintf(w, "200 OK")
	strInsertQ := "INSERT INTO account (details, hash) VALUES ('" + strDetails + "','" + strPass + "');"
	app.connector.Insert(strInsertQ)
}


//details->>'name' = ?
func (app *App) AccountID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	strEmail := r.Form.Get("email")
	println(string(strEmail) + " ? empty ?")
	selector := "SELECT id FROM account where details->>'email'='"+ strEmail +"'"
	var id int
	app.connector.Get(selector, &id)
	fmt.Fprintf(w, strconv.Itoa(id))
}