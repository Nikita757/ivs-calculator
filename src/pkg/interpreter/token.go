package interpeter

import "fmt"

const (
	OPERATOR = iota
	NUMBER
)

type Token struct {
	tokenType int
	stringValue string
	floatValue float64
}

type BinaryTree struct {
	token Token
	leftNode *BinaryTree
	rightNode *BinaryTree
}
