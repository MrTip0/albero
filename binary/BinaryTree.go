package binary

import "fmt"

func InitTree(val int) binaryTree {
	radice := new(nodo)
	radice.dx = nil
	radice.sx = nil
	radice.val = val
	return binaryTree{len: 1, radice: radice}
}

type binaryTree struct {
	len    int
	radice *nodo
}

type nodo struct {
	sx, dx *nodo
	val    int
}

func (b *binaryTree) InserisciOrdinato(val int) error {
	err := b.radice.InserisciOrdinato(val)
	if err == nil {
		b.len++
	}
	return err
}

func (b *binaryTree) Trova(val int) bool {
	return b.radice.Trova(val)
}

func (b *binaryTree) ToSlice() []int {
	i := 0
	r := make([]int, b.len)
	b.radice.ToSlice(r, &i)
	return r
}

func (n *nodo) InserisciOrdinato(val int) error {
	var toInit *nodo

	if val == n.val {
		return fmt.Errorf("valore già presente")

	} else if val > n.val {
		if n.dx == nil {
			toInit = new(nodo)
			n.dx = toInit
		} else {
			return n.dx.InserisciOrdinato(val)
		}

	} else {
		if n.sx == nil {
			toInit = new(nodo)
			n.sx = toInit
		} else {
			return n.sx.InserisciOrdinato(val)
		}
	}

	toInit.dx = nil
	toInit.sx = nil
	toInit.val = val
	return nil
}

func (n *nodo) ToSlice(s []int, i *int) {
	if n.sx != nil {
		n.sx.ToSlice(s, i)
	}
	s[*i] = n.val
	*i++
	if n.dx != nil {
		n.dx.ToSlice(s, i)
	}
}

func (n *nodo) Trova(val int) bool {
	if n.val == val {
		return true
	} else if val > n.val {
		if n.dx != nil {
			return n.dx.Trova(val)
		} else {
			return false
		}
	} else {
		if n.sx != nil {
			return n.sx.Trova(val)
		} else {
			return false
		}
	}
}
