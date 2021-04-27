package interpreter

import (
	"fmt"
	"ivs-calculator/pkg/mathfunc"
	"log"
	"strconv"
	"unicode"
)

type SyntaxError struct {
	index int
}

func (se *SyntaxError) Error() string {
	return fmt.Sprintf("Invalid syntax at index: %d.", se.index)
}

var operOrder = map[string]struct {
	prec   int
	rAssoc bool
}{
	"-": {2, false},
	"+": {2, false},
	"*": {3, false},
	"/": {3, false},
	"%": {3, false},
	"√": {4, true},
	"!": {4, true},
	"^": {4, true},
	"m": {5, true},
	"p": {5, true},
}

func Eval(input string) (float64, error)

// can return the index of invalid syntax,
// so that it can be highlighted in the GUI
func Parse(input string) (*TreeNode, error)

/**
 * evalOperator: evaluates operator node
 *
 * Calls Interpret() on left and right children of the node, then based
 * on the node's token.stringValue calls the correct function.
 *
 * @param node Pointer to the node being evaluated
 * @return float64 resulting from the called operator function
 * @return error if called on an unknown operator,
 * or when an error occurs when interpreting child nodes
 * or when calling the operator function
 */
func evalOperator(node *TreeNode) (float64, error) {
	left, err1 := Interpret(node.leftNode)
	if err1 != nil {
		return 0, err1
	}

	stringValue := node.token.stringValue

	// handle one operand operators
	if stringValue == "abs" {
		return mathfunc.AbsoluteValue(left), nil
	} else if stringValue == "fac" {
		return mathfunc.Factorial(left)
	}

	right, err2 := Interpret(node.rightNode)
	if err2 != nil {
		return 0, err2
	}

	// handle two operand operators
	switch stringValue {
	case "+":
		return mathfunc.Add(left, right), nil
	case "*":
		return mathfunc.Multiply(left, right), nil
	case "-":
		return mathfunc.Subtract(left, right), nil
	case "/":
		return mathfunc.Divide(left, right)
	case "mod":
		return mathfunc.Modulo(left, right)
	case "pow":
		return mathfunc.Power(left, right)
	case "root":
		return mathfunc.Root(left, right)
	default:
		return 0, fmt.Errorf("invalid operator: '%v'", node.token.stringValue)
	}
}

/**
 * evalNumber: evaluates number node by returning it's stored float64 value
 *
 * @param node Pointer to the node being evaluated
 * @return float64 number stored by the node's token
 */
func evalNumber(node *TreeNode) float64 {
	return node.token.floatValue
}

/**
 * Interpret: calculates the result as float64 of the expression represented by the parametr root
 *
 * @param root Pointer to the AST node being evaluated
 * @return float64 result of the whole expression
 * @return error if there was an error when evaluating the AST - see evalOperator for details
 */
func Interpret(root *TreeNode) (float64, error) {
	if root == nil { // interpreting an empty tree or node desn't make sense
		return 0, fmt.Errorf("cannot interpret an empty node")
	}

	if root.token.tokenType == OPERATOR {
		return evalOperator(root)
	} else if root.token.tokenType == NUMBER {
		return evalNumber(root), nil
	} else {
		return 0, fmt.Errorf("invalid token type: %d", root.token.tokenType)
	}
}

func toSlice(in string) ([]string, []int) {
	outSlice := make([]string, 0)
	wrongSynt := make([]int, 0)
	brackPos := make([]int, 0)
	absPos := make([]int, 0)
	openedBracket := false
	openedAbs := false
	consNum := false
	number := ""
	for i, tokenRune := range in {
		token := string(tokenRune)
		if token == "(" || token == ")" || token == "+" || token == "-" || token == "*" || token == "/" || token == "!" || token == "^" || token == "√" || token == "|" || token == "%" {
			if consNum {
				consNum = false
				outSlice = append(outSlice, number)
				number = ""
			}
			if token == "|" {
				if openedAbs {
					openedAbs = false
					absPos = absPos[:len(absPos)-1]
				} else {
					openedAbs = true
					absPos = append(absPos, i)
				}
			}

			if token == "(" {
				openedBracket = true
				brackPos = append(brackPos, i)
			}
			if token == ")" {
				if !openedBracket {
					wrongSynt = append(wrongSynt, i)
				} else {
					brackPos = brackPos[:len(brackPos)-1]
					openedBracket = false
				}
			}

			outSlice = append(outSlice, token)
		} else if unicode.IsDigit(tokenRune) || (consNum && (token == "," || token == ".")) {
			if token == "," || token == "." {
				number += "."
				continue
			}

			number += token
			consNum = true
		} else {
			wrongSynt = append(wrongSynt, i)
		}
	}
	if len(brackPos) != 0 {
		for _, pos := range brackPos {
			wrongSynt = append(wrongSynt, pos)
		}
	}
	if len(absPos) != 0 {
		for _, pos := range absPos {
			wrongSynt = append(wrongSynt, pos)
		}
	}

	if consNum {
		outSlice = append(outSlice, number)
	}

	return outSlice, wrongSynt
}

func inToPost(input []string) ([]string, int) {
	post := make([]string, 0)
	stack := make([]string, 0)
	afterOpPar := false
	lastDig := false
	for i, token := range input {
		switch token {
		case "(":
			lastDig = false
			stack = append(stack, token)
		case ")":
			lastDig = false
			if len(stack) == 0 {
				return nil, i
			}
			for {
				operator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if operator == "(" {
					break
				}
				post = append(post, operator)
				afterOpPar = true
			}
		case "+", "-", "/", "*", "%", "^", "!", "√":
			curOp := token
			if i == 0 && curOp == "-" {
				curOp = "m"
			} else if i == 0 && curOp == "+" {
				curOp = "p"
			} else if len(stack) > 0 {
				lastOp := stack[len(stack)-1]

				if !afterOpPar {
					if lastOp != ")" && curOp == "-" && lastOp != "p" && lastOp != "m" && !lastDig {
						curOp = "m"
					}

					if lastOp != ")" && curOp == "+" && lastOp != "p" && lastOp != "m" && !lastDig {
						curOp = "p"
					}
				}

				for {
					if !operOrder[curOp].rAssoc && operOrder[curOp].prec <= operOrder[lastOp].prec || operOrder[curOp].rAssoc && operOrder[curOp].prec < operOrder[lastOp].prec {
						operator := stack[len(stack)-1]
						stack = stack[:len(stack)-1]
						post = append(post, operator)

						if len(stack) == 0 {
							break
						} else {
							lastOp = stack[len(stack)-1]
						}
					} else {
						break
					}
				}
			}

			stack = append(stack, curOp)
			if token != "!" {
				lastDig = false
			}
			afterOpPar = false
		default:
			lastDig = true
			post = append(post, token)
		}
	}

	for len(stack) > 0 {
		post = append(post, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return post, -1
}

func toTreeOper(stack []*TreeNode, token string) []*TreeNode {
	var (
		t *Token
		l *TreeNode
		r *TreeNode
	)
	switch token {
	case "+", "-", "/", "*", "^", "%":
		l = stack[len(stack)-2]
		r = stack[len(stack)-1]
		log.Println(l, r)
	case "!", "√":
		l = stack[len(stack)-1]
		r = nil
	}

	if token == "%" {
		t = NewToken(OPERATOR, "mod", 0.0)
	} else if token == "^" {
		t = NewToken(OPERATOR, "pow", 0.0)
	} else if token == "√" {
		t = NewToken(OPERATOR, "root", 0.0)
	} else if token == "!" {
		t = NewToken(OPERATOR, "fac", 0.0)
	} else if token == "m" {
		unT := NewToken(NUMBER, "-1", -1.0)
		unN := NewNode(unT)
		l := stack[len(stack)-1]
		t := NewToken(OPERATOR, "*", 0.0)
		n := NewParent(t, l, unN)
		stack[len(stack)-1] = n

		return stack
	} else if token == "p" {
		unT := NewToken(NUMBER, "1", 1.0)
		unN := NewNode(unT)
		l := stack[len(stack)-1]
		t := NewToken(OPERATOR, "*", 0.0)
		n := NewParent(t, l, unN)
		stack[len(stack)-1] = n

		return stack
	} else {
		t = NewToken(OPERATOR, token, 0.0)
	}
	n := NewParent(t, l, r)
	if r == nil {
		stack[len(stack)-1] = n

		return stack
	}

	stack = stack[:len(stack)-1]
	stack[len(stack)-1] = n

	return stack
}

func postToTree(post []string) *TreeNode {
	stack := make([]*TreeNode, 0)

	for _, token := range post {
		switch token {
		case "+", "-", "/", "*", "^", "!", "%", "√", "m", "p":
			stack = toTreeOper(stack, token)
		default:
			fl, err := strconv.ParseFloat(token, 64)
			if err != nil {
				log.Println("error at: ", token)
			}

			t := NewToken(NUMBER, token, fl)
			n := NewNode(t)
			stack = append(stack, n)
			log.Println(n)
		}
	}

	return stack[0]
}
