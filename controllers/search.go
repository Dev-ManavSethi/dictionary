package controllers

import (
	"github.com/Dev-ManavSethi/dictionary/models"
	"github.com/Dev-ManavSethi/dictionary/search"
	"github.com/Dev-ManavSethi/dictionary/utils"
	"golang.org/x/net/websocket"
)

func Search(ws *websocket.Conn) {

	for {
		var SearchReq models.SearchRequest
		err := websocket.JSON.Receive(ws, &SearchReq)

		if err != nil {
			utils.HandleError(err, "Error recieving Search req through ws", "")
			//do something
		} else {
			Word := SearchReq.Word
			Exists := search.Search(Word)

			if !Exists {

				ws.Write([]byte("DNE"))
			}

			if Exists {
				ws.Write([]byte("E"))
			}

		}
	}

}
