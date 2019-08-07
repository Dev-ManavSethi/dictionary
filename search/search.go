package search

import (
	"github.com/Dev-ManavSethi/dictionary/models"
)

func Search(word string) bool {

	NodeExists := subNodeExists(models.Root, word[0])

	if !NodeExists {
		return false
	}
	if NodeExists {

		subNode := SubNodeSearch(models.Root, word[0])
		i := 1
		for NodeExists != false && i < len(word) {

			NodeExists = subNodeExists(subNode, word[i])

			if NodeExists {

				subNode = SubNodeSearch(subNode, word[i])
			}
			i++
		}

		if NodeExists == false {
			return false
		}

		if NodeExists == true && subNode.IsEnd == false && i == len(word) {

			return false
		}
	}

	return true

}

func subNodeExists(node *models.Node, alphabet byte) bool {

	if node == nil {
		return false
	}
	for _, subNode := range node.Alphabets {
		if subNode.Alphabet == alphabet {
			return true
		}
	}

	return false

}

func SubNodeSearch(node *models.Node, alphabet byte) *models.Node {
	if node == nil {
		return nil
	}

	for _, subNode := range node.Alphabets {
		if subNode.Alphabet == alphabet {

			return subNode
		}
	}

	return nil
}
