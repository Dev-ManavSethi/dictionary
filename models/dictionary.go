package models

type (
	Node struct {
		Alphabets []*Node
		Alphabet  byte
		IsEnd     bool
	}
)

var (

	DictionaryRoot Node
	Root *Node
)


