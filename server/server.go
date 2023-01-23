package server

import (
	"fmt"
	"net/http"
)

func Run() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/detail/", EachArtistHandler)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public/"))))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
