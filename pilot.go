package main

import (
	"io"
	"log"
	"net/http"
)

func pilot(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed read body request:", err)
		w.Write([]byte("Internal Server Error"))
		return
	}
	if droneConn == nil {
		w.Write([]byte("There is no connected client."))
		return
	}
	if err := droneConn.WriteJSON(string(b)); err != nil {
		log.Println("Failed send message to client:", err)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Write([]byte("Success."))
	return
}
