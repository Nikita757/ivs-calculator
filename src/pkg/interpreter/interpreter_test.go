package interpreter

import (
	"errors"
	"ivs-calculator/pkg/mathfunc"
	"testing"
)

//================
// Interpret tests

func TestInterpretEmpty(t *testing.T) {
	var nilNode *TreeNode = nil
	InterpretErrorTestCase(t, nilNode, errors.New("cannot interpret an empty node"))

	var emptyNode = &TreeNode{Token{OPERATOR, "", 0}, nil, nil}
	InterpretErrorTestCase(t, emptyNode, errors.New("cannot interpret an empty node"))

	var emptyChildren = &TreeNode{Token{OPERATOR, "+", 0}, nil, nil}
	InterpretErrorTestCase(t, emptyChildren, errors.New("cannot interpret an empty node"))
}

func TestInterpretOperators(t *testing.T) {
	InterpretResultTestCase(t, getTreeForOperator("+"), 9.0)
	InterpretResultTestCase(t, getTreeForOperator("-"), 1.0)
	InterpretResultTestCase(t, getTreeForOperator("*"), 20.0)

	res, _ := mathfunc.Divide(5.0, 4.0)
	InterpretResultTestCase(t, getTreeForOperator("/"), res)

	res, _ = mathfunc.Modulo(5.0, 4.0)
	InterpretResultTestCase(t, getTreeForOperator("mod"), res)

	res, _ = mathfunc.Power(5.0, 4.0)
	InterpretResultTestCase(t, getTreeForOperator("pow"), res)

	res, _ = mathfunc.Root(5.0, 4.0)
	InterpretResultTestCase(t, getTreeForOperator("root"), res)

	var tree = &TreeNode{Token{OPERATOR, "+", 0},
		&TreeNode{Token{OPERATOR, "*", 0},
			&TreeNode{Token{NUMBER, "", 3}, nil, nil},
			&TreeNode{Token{NUMBER, "", 2}, nil, nil}},
		&TreeNode{Token{NUMBER, "", 4.0}, nil, nil}}
	InterpretResultTestCase(t, tree, 10.0)

	tree = &TreeNode{Token{OPERATOR, "fac", 0},
		&TreeNode{Token{NUMBER, "", 5}, nil, nil}, nil}
	InterpretResultTestCase(t, tree, 120)

	tree = &TreeNode{Token{OPERATOR, "fac", 0},
		&TreeNode{Token{OPERATOR, "abs", 0},
			&TreeNode{Token{NUMBER, "", -5}, nil, nil}, nil}, nil}
	InterpretResultTestCase(t, tree, 120)
}

func getTreeForOperator(op string) *TreeNode {
	var tree = &TreeNode{Token{OPERATOR, op, 0},
		&TreeNode{Token{NUMBER, "", 5.0}, nil, nil},
		&TreeNode{Token{NUMBER, "", 4.0}, nil, nil}}
	return tree
}

// Test constants and operator validity
func TestInterpretInvalidData(t *testing.T) {
	var tree = &TreeNode{Token{2, "+", 0},
		&TreeNode{Token{NUMBER, "", 5.0}, nil, nil},
		&TreeNode{Token{NUMBER, "", 4.0}, nil, nil}}
	InterpretErrorTestCase(t, tree, errors.New("invalid token type: 2"))

	tree = &TreeNode{Token{OPERATOR, "(", 0},
		&TreeNode{Token{NUMBER, "", 5.0}, nil, nil},
		&TreeNode{Token{NUMBER, "", 4.0}, nil, nil}}
	InterpretErrorTestCase(t, tree, errors.New("invalid operator: '('"))
}

func TestInterpretNumber(t *testing.T) {
	var tree = &TreeNode{Token{NUMBER, "", 8.123}, nil, nil}
	InterpretResultTestCase(t, tree, 8.123)
}

// Test propagation of errors that arise in mathfunc
func TestInterpretOpError(t *testing.T) {
	var tree = &TreeNode{Token{OPERATOR, "+", 0},
		&TreeNode{Token{NUMBER, "", 4.0}, nil, nil},
		&TreeNode{Token{OPERATOR, "/", 0},
			&TreeNode{Token{NUMBER, "", 10.0}, nil, nil},
			&TreeNode{Token{NUMBER, "", 0.0}, nil, nil}}}
	InterpretErrorTestCase(t, tree, errors.New("cannot divide by zero"))

	tree = &TreeNode{Token{OPERATOR, "fac", 0},
		&TreeNode{Token{NUMBER, "", -1}, nil, nil}, nil}
	InterpretErrorTestCase(t, tree, errors.New("cannot calculate factorial of negative numbers"))

}

func InterpretErrorTestCase(t *testing.T, tree *TreeNode, expectedError error) {
	out, err := Interpret(tree)
	if out != 0.0 {
		t.Errorf("Interpret(%v) out = %s should be 0.0", tree, err)
	}
	if err == nil || (err.Error() != expectedError.Error()) {
		t.Errorf("Interpret(%v) err = %s should be %s", tree, err, expectedError)
	}
}

func InterpretResultTestCase(t *testing.T, tree *TreeNode, expectedOutput float64) {
	out, err := Interpret(tree)
	if err != nil {
		t.Errorf("Interpret(%v) err = %s should be nil", tree, err)
	}
	if out != expectedOutput {
		t.Errorf("Interpret(%v) out = %f should be %f", tree, out, expectedOutput)
	}
}

func TestToSlice(t *testing.T) {
	in := "1010+10/5"
	expOut := []string{"1010", "+", "10", "/", "5"}
	
	out, err := toSlice(in);
	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}


	in = "(50+(30/10)*5-2^5+5.5)"
	expOut = []string{"(", "50", "+", "(", "30", "/", "10", ")", "*", "5", "-", "2", "^", "5", "+", "5.5", ")"}

	out, err = toSlice(in);
	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}

	in = "√(5^|-5|-1)+5%5"
	expOut = []string{"2", "√", "(", "5", "^", "-", "5", "|", "-", "1", ")", "+", "5", "%", "5"}
	
	out, _ = toSlice(in);

	if (len(expOut) == len(out)) {
		for i := 0; i < len(expOut); i++ {
			if (expOut[i] != out[i]) {
				t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
			}
		}
	} else {
		t.Errorf("Length is %d should be %d", len(out), len(expOut))
	}

	in = ""
	out, _ = toSlice(in);

	if (out != nil) {
		t.Errorf("Slice should be nil")
	}

	in = ")"
	out, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "3√9"
	expOut = []string{"3", "√", "9"}
	out, err = toSlice(in);

	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}

	in = "(√16)"
	expOut = []string{"(", "2", "√", "16", ")"}
	out, err = toSlice(in);

	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}

	in = "(^"
	out, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "--5"
	expOut = []string{"+", "5"}
	out, err = toSlice(in);

	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}

	in = "+-5"
	expOut = []string{"-", "5"}
	out, err = toSlice(in);

	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}

	in = "++5"
	expOut = []string{"+", "5"}
	out, err = toSlice(in);

	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}

	in = "-+5"
	expOut = []string{"-", "5"}
	out, err = toSlice(in);

	if (len(err) > 0) {
		t.Errorf("Error given should be no error")
	} else {
		if (len(expOut) == len(out)) {
			for i := 0; i < len(expOut); i++ {
				if (expOut[i] != out[i]) {
					t.Errorf("Sliced out '%s' should be '%s'", out[i], expOut[i])
				}
			}
		} else {
			t.Errorf("Length is %d should be %d", len(out), len(expOut))
		}
	}

	in = "1*%5"
	_, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "5|||"
	_, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "()("
	_, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "(()"
	_, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "(*. 5"
	_, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "5..5"
	_, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}

	in = "5..5"
	_, err = toSlice(in);

	if (len(err) <= 0) {
		t.Errorf("No error given should be error")
	}
}
