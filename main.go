package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/drone", drone)
	r.HandleFunc("/pilot", pilot)

	log.Println("Server started at 8000")
	http.ListenAndServe(":8000", r)
}
