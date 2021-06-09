package main

import (
	"io"
	"net/http"
)

func pilot(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	err := droneConn.WriteJSON(string(b))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("success"))
	return
}
