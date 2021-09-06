package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HTTP Router")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
