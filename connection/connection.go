package connection

import (
	"encoding/json"
	jsonResponse "github.com/easy/lego/json"
	"github.com/easy/lego/json/lingualeo"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"log"
	"errors"
)

var cookieJar, _ = cookiejar.New(nil)

var client = &http.Client{
	Jar: cookieJar,
}

const stringUrlApi = "http://api.lingualeo.com/"
const stringUrlCommon = "http://lingualeo.com/"
const countWordsInPage = 100

type Connector interface {
	Connect()
}

type SimpleConnector struct {
	Login string
	Pass  string
}

func (s *SimpleConnector) SetPass(pass string) {
	s.Pass = pass
}

func (s *SimpleConnector) SetLogin(login string) {
	s.Login = login
}

func (s *SimpleConnector) Connect() error {
	authRequest := stringUrlApi + "api/login"
	authParams := url.Values{
		"email":    {s.Login},
		"password": {s.Pass},
	}
	resp, err := client.PostForm(authRequest, authParams)
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error connect to lingualeo account: %s", resp)
		err = errors.New("Error connect to account")
	}

	defer resp.Body.Close()
	return err
}

//TODO: urlBuilder for "string_url_common" and rename variable
func (s *SimpleConnector) GetPageOfDictionary(index int) jsonResponse.Dictionary {
	requestStr := stringUrlCommon + "userdict/json?groupId=dictionary&filter=learned&page=" + strconv.Itoa(index)
	requestArgs := url.Values{}
	resp, err := client.PostForm(requestStr, requestArgs)
	defer resp.Body.Close()
	var m lingualeo.LeoDictionaryImpl
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&m)
	if err != nil {
		panic(err.Error())
	}
	return &m
}

func (s *SimpleConnector) GetCountWords() int {
	return s.GetPageOfDictionary(0).GetCountWords()
}

func (s *SimpleConnector) GetAllWords() []jsonResponse.Word {
	var result []jsonResponse.Word
	chRes := make(chan *[]jsonResponse.Word)
	countWords := s.GetCountWords()
	countPages := (countWords / countWordsInPage) + 1

	for i := 1; i <= countPages; i++ {
		go func(i int) {
			tmpArray := new([]jsonResponse.Word)
			*tmpArray = s.GetPageOfDictionary(i).GetWords()
			chRes <- tmpArray
		}(i)
	}

	for i := 0; i < countPages; i++ {
		tmpData := <-chRes
		result = append(result, *tmpData...)
	}
	return result
}

func (s *SimpleConnector) AddWord(word, translate, context string) {
	requestStr := stringUrlApi + "addword"
	requestArgs := url.Values{
		"word":    {word},
		"tword":   {translate},
		"context": {context},
	}
	resp, err := client.PostForm(requestStr, requestArgs)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
}
