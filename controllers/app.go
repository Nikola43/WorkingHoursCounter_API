package controllers

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/panjf2000/ants/v2"
	"log"
	"net/http"
)

var myPool *ants.Pool

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.Use(JwtAuthentication) //attach JWT auth middleware
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	// configure allowed requests
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	// create pool
	myPool, _ = ants.NewPool(1000000)
	defer myPool.Release()

	// run server
	fmt.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(addr, handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(a.Router)))
}

func (a *App) initializeRoutes() {
	// USER
	a.Router.HandleFunc("/api/user/login", Login).Methods("GET")
	/*
		a.Router.HandleFunc("/api/user/new", SingUp).Methods("POST")
		a.Router.HandleFunc("/api/user", GetAll).Methods("GET")
		a.Router.HandleFunc("/api/user/{id:[0-9]+}", GetUserByID).Methods("GET")
		a.Router.HandleFunc("/api/user/{invite_id:[0-9]+}/{invited_id:[0-9]+}", InviteUser).Methods("GET")
		a.Router.HandleFunc("/api/user/{id:[0-9]+}/invited", GetNumberOfInvitedUsers).Methods("GET")
		a.Router.HandleFunc("/api/user/total", GetNumberOfUsers).Methods("GET")
		a.Router.HandleFunc("/api/test", TestRoutine).Methods("GET")
	*/

}
