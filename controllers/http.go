package controllers

import (
	"log"
	"net/http"

	"github.com/Dev-ManavSethi/dictionary/search"
)

func Home(w http.ResponseWriter, r *http.Request) {

}

func Add(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {

	}

	word := r.FormValue("word")
	meaning := r.FormValue("meaning")

	ok := search.Add(word, meaning)

	if ok {
		log.Println("Added: ", word)
	}
	if !ok {
		log.Println("Not able to add: ", word)
	}

}
