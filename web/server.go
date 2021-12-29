package main

import (
	"log"
	"net/http"

	"web/art"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.HandleFunc("/", art.Home)
	http.HandleFunc("/ascii-art", art.Asciiart)
	log.Println("\nStarting the web server on localhost:8070")
	if err := http.ListenAndServe(":8070", nil); err != nil {
		log.Println(err.Error())
		return
	}
}
