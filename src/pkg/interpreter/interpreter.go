package interpreter

import (
	"errors"
	"fmt"
	"ivs-calculator/pkg/mathfunc"
)

type SyntaxError struct {
	index int
}

func (se *SyntaxError) Error() string {
	return fmt.Sprintf("Invalid syntax at index: %d.", se.index)
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
		return mathfunc.Substract(left, right), nil
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
