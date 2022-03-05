package binary

import "fmt"

func InitTree(val int) binaryTree {
	radice := newNodo(val)
	return binaryTree{len: 1, radice: radice}
}

func newNodo(val int) *nodo {
	r := new(nodo)
	r.val = val
	r.dx = nil
	r.sx = nil
	return r
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
	if val == n.val {
		return fmt.Errorf("valore già presente")

	} else if val > n.val {
		if n.dx == nil {
			n.dx = newNodo(val)
		} else {
			return n.dx.InserisciOrdinato(val)
		}

	} else {
		if n.sx == nil {
			n.sx = newNodo(val)
		} else {
			return n.sx.InserisciOrdinato(val)
		}
	}
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
