package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sparta/config"
	"sparta/core"
)

func main() {
	config.InitConfig()
	core.ConnectDB()

	r := mux.NewRouter().StrictSlash(false)
	api := r.PathPrefix("/api").Subrouter()

	for _, route := range routes {
		api.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(makeHandler(route.HandlerFunc))
	}

	// 需要走 JWTHttpMiddleware
	//rest := api.PathPrefix("/rest").Subrouter()
	//for _, route := range restRoutes {
	//	rest.
	//		Methods(route.Method).
	//		Path(route.Pattern).
	//		Name(route.Name).
	//		Handler(makeHandler(route.HandlerFunc))
	//}
	//rest.Use(middleware.JWTHTTPMiddleware)

	http.Handle("/", r)
	err := http.ListenAndServe(":5005", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
