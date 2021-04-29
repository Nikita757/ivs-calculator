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

func NewParent(pToken *Token, lKid, rKid *TreeNode) *TreeNode {
	parent := NewNode(pToken)
	parent.leftNode = lKid
	parent.rightNode = rKid
	return parent
}

func NewNode(t *Token) *TreeNode {
	n := &TreeNode{token: *t, leftNode: nil, rightNode: nil}
	return n
}

func NewToken(tType int, strVal string, flVal float64) *Token {
	t := &Token{tokenType: tType, stringValue: strVal, floatValue: flVal}
	return t
}
