package binary

func (b *binaryTree) Balance() {
	sl := b.ToSlice()
	b.radice = newNodo(0)
	b.radice.inserisciBilanciato(sl, 0, len(sl)-1)
}

func (n *nodo) inserisciBilanciato(values []int, min, max int) {
	nel := max - min

	if nel == 0 {
		n.val = values[min]
	} else if nel == 1 {
		n.val = values[min]
		n.dx = newNodo(values[max])
	} else if nel == 2 {
		n.sx = newNodo(values[min])
		n.val = values[min+1]
		n.dx = newNodo(values[max])
	} else {
		centro := (max + min) / 2
		n.val = values[centro]
		n.sx = newNodo(0)
		n.sx.inserisciBilanciato(values, min, centro-1)
		n.dx = newNodo(0)
		n.dx.inserisciBilanciato(values, centro+1, max)
	}
}
