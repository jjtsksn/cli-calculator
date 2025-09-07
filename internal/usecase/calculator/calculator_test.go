package calculator

import (
	"slices"
	"testing"

	"github.com/jjtsksn/cli-calculator/pkg/splitter"
)

func TestInfinixToRPN(t *testing.T) {
	calc := NewCalculator()

	testCase1 := []string{"1", "+", "(", "5", "+", "3", "*", "2", "/", "3", ")", "+", "1", "*", "8"}
	testCase2 := []string{"1", "+", "2", "*", "(", "3", "+", "2", "*", "2", ")"}
	testCase3 := []string{"123", "*", "(", "10", "/", "5", "+", "25", ")"}

	exceptedCase1 := []string{"1", "5", "3", "2", "*", "3", "/", "+", "+", "1", "8", "*", "+"}
	exceptedCase2 := []string{"1", "2", "3", "2", "2", "*", "+", "*", "+"}
	exceptedCase3 := []string{"123", "10", "5", "/", "25", "+", "*"}

	calc.infixToRPN(testCase1)
	t1 := calc.rpn
	calc.rpn = []string{}
	calc.infixToRPN(testCase2)
	t2 := calc.rpn
	calc.rpn = []string{}
	calc.infixToRPN(testCase3)
	t3 := calc.rpn

	if !slices.Equal(t1, exceptedCase1) {
		t.Fatalf("Excpected: %v\nReceived: %v\n", exceptedCase1, t1)
	}
	if !slices.Equal(t2, exceptedCase2) {
		t.Fatalf("Excpected: %v\nReceived: %v\n", exceptedCase2, t2)
	}
	if !slices.Equal(t3, exceptedCase3) {
		t.Fatalf("Excpected: %v\nReceived: %v\n", exceptedCase3, t3)
	}
}

func Test(t *testing.T) {
	spl := splitter.NewBasicStringSplitter()
	calc := NewCalculator()
	testCase1 := "1+(5+3*2/3)+1*8"
	testCase2 := "1+2*(3+2*2)"
	testCase3 := "123*(10/5+25)"

	t1, _ := calc.Calculate(testCase1, spl)
	t2, _ := calc.Calculate(testCase2, spl)
	t3, _ := calc.Calculate(testCase3, spl)

	if t1 != "16" {
		t.Fatalf("Excpected: 16\nReceived: %v\n", t1)
	}
	if t2 != "15" {
		t.Fatalf("Excpected: 15\nReceived: %v\n", t2)
	}
	if t3 != "3321" {
		t.Fatalf("Excpected: 3321\nReceived: %v\n", t3)
	}
}
