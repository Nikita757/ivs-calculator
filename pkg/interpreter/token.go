package interpreter

/**
 * Constants to define type of a token
 */
const (
	OPERATOR = iota
	NUMBER
)

/**
 * Token: Data structure for tokens
 */
type Token struct {
	tokenType   int
	stringValue string
	floatValue  float64
}

/**
 * TreeNode: Data structure for nodes of the tree
 */
type TreeNode struct {
	token     Token
	leftNode  *TreeNode
	rightNode *TreeNode
}

/**
 * NewParent: Creates new parent node
 *
 * @param pToken Pointer to a parent token
 * @param lKid Pointer to the left child
 * @param rKid Pointer to the right child
 * @return *TreeNode Pointer to the created parent node
 */
func NewParent(pToken *Token, lKid, rKid *TreeNode) *TreeNode {
	parent := NewNode(pToken)
	parent.leftNode = lKid
	parent.rightNode = rKid
	return parent
}

/**
 * NewNode: Creates new node
 *
 * @param t Pointer to the token
 * @return *TreeNode Pointer to the created node
 */
func NewNode(t *Token) *TreeNode {
	n := &TreeNode{token: *t, leftNode: nil, rightNode: nil}
	return n
}

/**
 * NewToken: Creates new token
 *
 * @param tType Type of the token can be an OPERATOR or a NUMBER
 * @param strVal String value of the token
 * @param flVal Float value of the token
 * @return *Token Pointer to the new token
 */
func NewToken(tType int, strVal string, flVal float64) *Token {
	t := &Token{tokenType: tType, stringValue: strVal, floatValue: flVal}
	return t
}
