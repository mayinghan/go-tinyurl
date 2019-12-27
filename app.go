package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
)

// App Encapsulates Env, Router and middlewares
type App struct {
	Router *mux.Router
}

type shortenReq struct {
	URL                 string `json:"url" validate:"nonzero"`
	ExpirationInMinutes int64  `json:"experiation_in_min" validate:"min=0"`
}

type shortLinkRes struct {
	Shortlink string `json:"shortlink"`
}

// Initialize for App
func (app *App) Initialize() {
	// set log formatter
	log.SetFlags(log.LstdFlags | log.Lshortfile) //日志日期 ｜ 日志文件名
	app.Router = mux.NewRouter()
	app.initializeRouters()
}

func (app *App) initializeRouters() {
	app.Router.HandleFunc("/api/shorten", app.createShortlink).Methods("POST")
	app.Router.HandleFunc("/api/info", app.getShortlink).Methods("GET")
	app.Router.HandleFunc("/api/{shortlink:[a-zA-Z0-9]{1,11}}", app.redirect).Methods("GET")
}

func (app *App) createShortlink(res http.ResponseWriter, req *http.Request) {
	var shortReq shortenReq
	if err := json.NewDecoder(req.Body).Decode(&shortReq); err != nil {
		return
	}

	// check shortenReq cretaria
	if err := validator.Validate(shortReq); err != nil {
		return
	}

	defer req.Body.Close()

	fmt.Printf("%v\n", shortReq)
}

func (app *App) getShortlink(res http.ResponseWriter, req *http.Request) {

}

func (app *App) redirect(res http.ResponseWriter, req *http.Request) {

}

func main() {

}
