package main

import (
	"fmt"
	"github.com/easy/lego/connection"
	"encoding/json"
	entries "github.com/easy/lego/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	entries2 "github.com/easy/lego/db/entries"
	"log"
)

func (app *App) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "{ name: successful response  }")
}

//Endopoint: get all known words from service in user registered
//Params: email, pass. (Only for lingualeo) in future jsonToken and parse need ConnectorServices
//Return: json array words
func (app *App) GetAllWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	email := r.Form.Get("email")
	pass := r.Form.Get("pass")
	sp := connection.SimpleConnector{Login: string(email), Pass: string(pass)}
	err := sp.Connect()
	if err != nil {
		log.Printf("Error connect to lingualeo account: %s, %s", email, pass)
		fmt.Fprintf(w, "Internal error: Error connect to ligualeo account")
		return
	}
	dict := sp.GetAllWords()
	resJson, err := json.Marshal(&dict)
	if err != nil {
		log.Print("Error connect to lingualeo account")
		fmt.Fprintf(w, "Internal error: not valid json response")
		return
	}
	fmt.Fprintf(w, string(resJson))
}

//Endpoint: Create new account in db
//Params: details, pass. Details - is json which contained difficult information (email, username, ... , etc)
//Return: result code operation
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

//Endpoint: Auth user
//Params: email, password
//Return result operation add
func (app *App) LogIn (w http.ResponseWriter, r * http.Request, _ httprouter.Params) {
	strEmail := r.FormValue("email")
	strPass := r.FormValue("password")
	selector := "SELECT * FROM account where details->>'email'='"+ strEmail +"'"
	var account entries2.Account
	app.connector.Get(selector, &account)
	if account.Hash != strPass {
		fmt.Fprint(w, "400")
		return
	}
}

//Endpoint: Add know word in db
//Params: id_user, word
//Return result operation add word in account_word and word tables
func (app *App) AddCustomWord(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//TODO: check exist account
	accountID := r.FormValue("id_user")
	wordValue := r.FormValue("word")

	isWordExistInDB := false
	fmt.Println(wordValue)
	app.connector.Get(entries2.ExistWord(wordValue), &isWordExistInDB)

	if !isWordExistInDB {
		app.connector.Insert(entries2.AddWord(wordValue))
	}

	wordID := ""

	app.connector.Get(entries2.GetWordID(wordValue), &wordID)
	app.connector.Insert(entries2.AddUserWord(accountID, wordID))

	fmt.Fprintf(w, "200 OK")
}

//TODO: add endpoint for add services

//Endpoint: Get account_id from email (details->>'email' = ?)
//params: email
//return account_id
func (app *App) AccountID(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	strEmail := r.Form.Get("email")
	println(string(strEmail) + " ? empty ?")
	selector := "SELECT id FROM account where details->>'email'='"+ strEmail +"'"
	var id int
	app.connector.Get(selector, &id)
	fmt.Fprintf(w, strconv.Itoa(id))
}


func (app *App) GetAllCustomWords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Yes")
	accountID := r.FormValue("id_user")
	words := []entries2.Word{}
	app.connector.Select(entries2.GetAllCustomWords(accountID), &words)

	responseJson,err := json.Marshal(words)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(responseJson))
	fmt.Fprintf(w, string(responseJson))
}