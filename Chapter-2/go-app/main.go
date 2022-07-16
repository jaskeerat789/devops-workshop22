package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.helloWorld).Methods("GET")
}

func (a *App) run(addr string) {
	fmt.Println("This Server is a Fucking Spaceship 🚀")
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("<h1>Welcome to DevOps Workshop</h1><p>Golang app v1.0</p>"))
}

func main() {
	a := App{}
	a.initialize()
	port := os.Getenv("PORT")
	a.run(fmt.Sprintf(":%s", port))
}
