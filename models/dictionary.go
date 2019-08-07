package models

type (
	Node struct {
		Alphabets []*Node
		Alphabet  byte
		IsEnd     bool
		Meaning   string
	}
)

var (
	DictionaryRoot Node
	Root           *Node
)
