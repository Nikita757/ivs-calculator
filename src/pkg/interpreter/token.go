package interpeter

import "fmt"

const (
	OPERATOR = iota
	NUMBER
)

type Token struct {
	tokenType int
	value string
}

type BinaryTree struct {
	token Token
	leftNode *BinaryTree
	rightNode *BinaryTree
}
