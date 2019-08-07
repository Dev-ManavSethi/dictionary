package search

import (
	"github.com/Dev-ManavSethi/dictionary/models"
)

func Add(word, meaning string) bool {

	i := 0

	NodeExists := subNodeExists(models.Root, word[i])
	//var meaningg string = meaning
	if !NodeExists {
		//add node
		subNode := models.Root

		for i < len(word) {

			var isEnd bool
			if i == len(word)-1 {
				isEnd = true
			}
			subsubNode := &models.Node{
				Alphabet:  word[i],
				Alphabets: []*models.Node{},
				IsEnd:     isEnd,
			}

			if isEnd {
				subsubNode.Meaning = meaning
			}
			subNode.Alphabets = append(subNode.Alphabets, subsubNode)

			subNode = subsubNode
			i++

		}
	}

	if NodeExists {
		var subNode *models.Node
		subNode = models.Root
		for NodeExists == true {

			subSubNode := SubNodeSearch(subNode, word[i])

			i++

			if subSubNode != nil {

				subNode = subSubNode
			}

			NodeExists = subNodeExists(subNode, word[i])

		}
		if NodeExists == false {
			//start adding

			for i < len(word) {
				var isEnd bool
				if i == len(word)-1 {
					isEnd = true
				}

				subsubNode := &models.Node{
					Alphabet:  word[i],
					Alphabets: []*models.Node{},
					IsEnd:     isEnd,
				}
				subNode.Alphabets = append(subNode.Alphabets, subsubNode)
				subNode = subsubNode
				i++

			}
		}
	}

	return true

}
