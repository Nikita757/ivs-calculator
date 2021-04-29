package interpreter

const (
	OPERATOR = iota
	NUMBER
)

type Token struct {
	tokenType   int
	stringValue string
	floatValue  float64
}

type TreeNode struct {
	token     Token
	leftNode  *TreeNode
	rightNode *TreeNode
}
