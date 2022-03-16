package main

import (
	"fmt"
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, ты на домашней странице!")
}

func catalog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Привет, ты в каталоге")
}

func main() {
	http.HandleFunc("/", home) //  роутер
	http.HandleFunc("/catalog", catalog) //  роутер
	err := http.ListenAndServe(":5843", nil) // порт веб-сервера

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}