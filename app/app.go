package app

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
	app.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}}", app.redirect).Methods("GET")
}

func (app *App) createShortlink(res http.ResponseWriter, req *http.Request) {
	var shortReq shortenReq
	if err := json.NewDecoder(req.Body).Decode(&shortReq); err != nil {
		fmt.Println("decode error")
		return
	}

	// check shortenReq cretaria
	if err := validator.Validate(shortReq); err != nil {
		fmt.Println("validate error")
		return
	}

	defer req.Body.Close()

	fmt.Printf("%v\n", shortReq)
}

func (app *App) getShortlink(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	s := vals.Get("shortlink")

	fmt.Printf("%s\n", s)
}

func (app *App) redirect(w http.ResponseWriter, r *http.Request) {
	// need to returen status code 302
	vars := mux.Vars(r) // return a dict
	fmt.Printf("%s\n", vars["shortlink"])
}

// Run the app at certain address
func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}
