package app

import (
	"github.com/gorilla/mux"
	"golang-restful-api/app/handler"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router = mux.NewRouter().StrictSlash(true)
	a.setRouters()
}

func (a *App) setRouters() {
	// Routing for handling the students
	a.Get("/students", a.handleRequest(handler.GetAllStudents))
	a.Post("/students", a.handleRequest(handler.CreateStudent))
	a.Get("/students/{id}", a.handleRequest(handler.GetStudent))
	a.Put("/students/{id}", a.handleRequest(handler.UpdateStudent))
	a.Delete("/students/{id}", a.handleRequest(handler.DeleteStudent))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
