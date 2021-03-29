package interpreter

import (
	"errors"
	"fmt"
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

func Interpret(root *TreeNode) (float64, error)
