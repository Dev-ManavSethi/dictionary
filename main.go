package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Dev-ManavSethi/dictionary/models"

	"golang.org/x/net/websocket"

	"github.com/Dev-ManavSethi/dictionary/controllers"
)

func init() {

	models.Root = &models.DictionaryRoot

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Home)
	mux.HandleFunc("/add", controllers.Add)
	mux.Handle("/search", websocket.Handler(controllers.Search))

	log.Println("Listening on: " + os.Getenv("PORT"))
	log.Fatalln(http.ListenAndServe(":"+os.Getenv("PORT"), mux))

}
