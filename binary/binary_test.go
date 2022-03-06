package binary_test

import (
	"fmt"
	"testing"

	"github.com/MrTip0/albero/binary"
)

func TestBalanceWith1Element(t *testing.T) {
	tree := binary.InitTree(5)
	tree.Balance()
}

func TestBalanceWith2Element(t *testing.T) {
	tree := binary.InitTree(5)
	tree.InserisciOrdinato(8)
	tree.Balance()
}

func TestBalanceWith3Element(t *testing.T) {
	tree := binary.InitTree(5)
	tree.InserisciOrdinato(2)
	tree.InserisciOrdinato(6)
	tree.Balance()
}

func TestInsert10Numbers(t *testing.T) {
	toInsert := []int{1, 7, 12, 54, 11, 65, 23, 62, 32}
	tree := binary.InitTree(10)

	for _, v := range toInsert {
		tree.InserisciOrdinato(v)
	}

	expected := "[1 7 10 11 12 23 32 54 62 65]"

	if r := fmt.Sprintf("%v", tree.ToSlice()); r != expected {
		t.Errorf("Expected %s, got %s", expected, r)
	}
}
