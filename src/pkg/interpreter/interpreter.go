package interpreter

import (
	"fmt"
	"ivs-calculator/pkg/mathfunc"
	"strconv"
	"unicode"
)

type SyntaxError struct {
	index int
}

func (se *SyntaxError) Error() string {
	return fmt.Sprintf("Invalid syntax at index: %d.", se.index)
}

// order of operations by their precedence in expression
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

/**
 * Parse: parses inputted math expression from infix notation into a binary expression tree
 *
 * Calls toSlice() on string expression to produce a valid slice of expression
 * If expression contained wrong syntax then a slice with positions of those mistakes is returned with a nil root
 * After it calls intoPost() to convert infix slice into a postfix one, since it's easier to convert into a tree
 * In the end postToTree() is being called to convert postfix slice into a tree
 *
 *
 * @param input infix expression to get parsed
 * @return *TreeNode root of a binary expression tree
 * @return []int slice with positions of syntax errors, if such've been found
 */
func Parse(input string) (*TreeNode, []int) {
	expSlice, wrongSynt := toSlice(input)
	if len(wrongSynt) != 0 {
		return nil, wrongSynt
	}
	post := inToPost(expSlice)
	root := postToTree(post)
	return root, nil
}

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

/**
 * toSlice: converts math expression from string to slice
 *
 * @param in inputted math expression in form of a string
 * @return []string slice consisted of inputted math expression
 * @return []int slice of mistakes in mathematical notation. Consider as an error return, if length of it is > 0
 */
func toSlice(in string) ([]string, []int) {
	if in == "" {
		return nil, nil
	}
	outSlice := make([]string, 0)
	wrongSynt := make([]int, 0)
	brackPos := make([]int, 0)
	absPos := make([]int, 0)
	openedAbs := false
	openedBr := false
	consNum := false
	isFloat := false
	wantPow := false
	closedBr := false
	number := ""
	for i, tokenRune := range in {
		token := string(tokenRune)
		if token == "(" || token == ")" || token == "+" || token == "-" || token == "*" || token == "/" || token == "!" || token == "^" || token == "√" || token == "|" || token == "%" {
			if i == 0 && (token == ")" || token == "*" || token == "/" || token == "!" || token == "^" || token == "%") {
				wrongSynt = append(wrongSynt, i)
			}
			// append a number to slice if it's construction is over
			if consNum {
				consNum = false
				outSlice = append(outSlice, number)
				number = ""
			}
			//
			if wantPow {
				wantPow = false
				wrongSynt = append(wrongSynt, i)
				continue
			}
			if token == "^" {
				prev := outSlice[len(outSlice)-1]
				_, err := strconv.Atoi(prev)
				if err != nil {
					wrongSynt = append(wrongSynt, i)
					continue
				}
				wantPow = true
			}
			if token == "-" && i > 0 {
				if outSlice[len(outSlice)-1] == "-" {
					outSlice[len(outSlice)-1] = "+"
				} else if outSlice[len(outSlice)-1] == "+" {
					outSlice[len(outSlice)-1] = "-"
				}
			}
			if token == "+" && i > 0 {
				if outSlice[len(outSlice)-1] == "+" {
					outSlice[len(outSlice)-1] = "+"
				} else if outSlice[len(outSlice)-1] == "-" {
					outSlice[len(outSlice)-1] = "-"
				}
			}
			if token == "*" || token == "/" || token == "!" || token == "%" {
				prev := outSlice[len(outSlice)-1]
				if prev == "*" || prev == "/" || prev == "!" || prev == "%" || prev == "+" || prev == "-" {
					wrongSynt = append(wrongSynt, i)
					continue
				}
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
				if closedBr {
					wrongSynt = append(wrongSynt, i)
					continue
				}
				openedBr = true
				brackPos = append(brackPos, i)
			}
			if token == ")" {
				if len(brackPos) == 0 {
					wrongSynt = append(wrongSynt, i)
					continue
				}
				if openedBr {
					brackPos = brackPos[:len(brackPos)-1]
				}
				closedBr = true
				brackPos = brackPos[:len(brackPos)-1]
				outSlice = append(outSlice, token)
				continue
			}
			closedBr = false
			openedBr = false
			outSlice = append(outSlice, token)
		} else if unicode.IsDigit(tokenRune) || (consNum && (token == "," || token == ".")) {
			if closedBr {
				closedBr = false
				wrongSynt = append(wrongSynt, i)
				continue
			}
			if wantPow {
				wantPow = false
			}
			if token == "," || token == "." {
				if !isFloat {
					number += "."
					isFloat = true
					continue
				}
				wrongSynt = append(wrongSynt, i)
				continue
			}

			number += token
			consNum = true
		} else if unicode.IsSpace(tokenRune) {
			continue
		} else {
			wrongSynt = append(wrongSynt, i)
			continue
		}
	}
	// if some brackets left in stack append their positions in wrongSynt
	if len(brackPos) != 0 {
		for _, pos := range brackPos {
			wrongSynt = append(wrongSynt, pos)
		}
	}
	// if some abs brackets left in stack append their positions in wrongSynt
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

/**
 * inToPost: converts infix expression into postfix
 *
 * @param input expression in infix notation
 * @return []string slice of expression in postfix notation
 */
func inToPost(input []string) []string {
	post := make([]string, 0)
	stack := make([]string, 0)
	openedAbs := false
	afterOpPar := false
	lastDig := false
	for i, token := range input {
		switch token {
		case "(":
			lastDig = false
			stack = append(stack, token)
		case ")":
			lastDig = false
			for {
				operator := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if operator == "(" {
					break
				}
				post = append(post, operator)
				afterOpPar = true
			}
		case "|":
			lastDig = false
			if openedAbs {
				for {
					operator := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					if operator == "|" {
						break
					}
					post = append(post, operator)
					afterOpPar = true
					post = append(post, "abs")
				}
			} else {
				openedAbs = true
				stack = append(stack, token)
			}
		case "+", "-", "/", "*", "%", "^", "!", "√":
			curOp := token
			if i == 0 && curOp == "-" { // checking if current operator is unary minus in the beginning of an expression
				curOp = "m"
			} else if i == 0 && curOp == "+" { // checking if current operator is unary plus in the beginning of an expression
				curOp = "p"
			} else if len(stack) > 0 {
				lastOp := stack[len(stack)-1]

				// checking if current operator is unary plus or minus
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
	return post
}

/**
 * toTreeOper: converts operator into according node type
 *
 * @param stack slice of nodes
 * @param token operator from postfix expression
 * @return []*TreeNode returns updated stack of nodes with assigned operands to operator
 */
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
	case "!", "√", "abs":
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
	} else if token == "abs" {
		t = NewToken(OPERATOR, token, 0.0)
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

/**
 * postToTree: converts postfix expression to a binary expression tree
 *
 * @param post slice of a postfix expression to be converted into a tree
 * @return *TreeNode root of a tree
 */
func postToTree(post []string) *TreeNode {
	stack := make([]*TreeNode, 0)

	for _, token := range post {
		switch token {
		case "+", "-", "/", "*", "^", "!", "%", "√", "abs", "m", "p":
			stack = toTreeOper(stack, token)
		default:
			fl, _ := strconv.ParseFloat(token, 64)

			t := NewToken(NUMBER, token, fl)
			n := NewNode(t)
			stack = append(stack, n)
		}
	}

	return stack[0]
}
