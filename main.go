package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
	"webserver/api"
	"webserver/auth"
)

func main() {
	//to view static files just open a web browser at http://localhost:8000

	//this is the default router
	r := mux.NewRouter()

	authSubRouter := r.PathPrefix("/auth").Subrouter()
	authSubRouter.Use(auth.CheckUser)
	apiSubRouter := authSubRouter.PathPrefix("/api").Subrouter()

	//if we want an optional parameter we must specify the route two times
	//test this with wget -q -O - --header 'X-AUTH-USERNAME: antares' http://localhost:8000/auth/api/hello
	apiSubRouter.Path("/hello").Methods(http.MethodGet).HandlerFunc(api.SayHello)

	//the placeholder "name" will be available in the handler func as a variable
	//test this with wget -q -O - --header 'X-AUTH-USERNAME: antares' http://localhost:8000/auth/api/hello/MYNAME
	apiSubRouter.Path("/hello/{name}").Methods(http.MethodGet).HandlerFunc(api.SayHello)

	// This will serve files under http://localhost:8000/<filename>
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend")))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		fmt.Println(t)
		return nil
	})

	log.Fatal(srv.ListenAndServe())
}
