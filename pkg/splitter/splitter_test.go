package splitter

import (
	"slices"
	"testing"
)

func TestSplit(t *testing.T) {
	spl := NewBasicStringSplitter()
	testCase1 := "1+(5+3*2/3)+1*8"
	testCase2 := "1+2*(3+2*2)"
	testCase3 := "123*(10/5+25)"

	exceptedCase1 := []string{"1", "+", "(", "5", "+", "3", "*", "2", "/", "3", ")", "+", "1", "*", "8"}
	exceptedCase2 := []string{"1", "+", "2", "*", "(", "3", "+", "2", "*", "2", ")"}
	exceptedCase3 := []string{"123", "*", "(", "10", "/", "5", "+", "25", ")"}

	t1 := spl.Split(testCase1)
	t2 := spl.Split(testCase2)
	t3 := spl.Split(testCase3)

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
