package binary

func (b *binaryTree) Balance() {
	sl := b.toSliceNodo()
	b.radice = sl[b.len/2]
	if b.len == 2 {
		b.radice.inserisciSx(sl, 0, 0)
	} else if b.len > 2 {
		b.radice.inserisciSx(sl, 0, (b.len/2)-1)
		b.radice.inserisciDx(sl, (b.len/2)+1, b.len-1)
	}
}

func (b *binaryTree) toSliceNodo() []*nodo {
	i := 0
	r := make([]*nodo, b.len)
	b.radice.toSliceNodo(r, &i)
	return r
}

func (n *nodo) toSliceNodo(s []*nodo, i *int) {
	if n.sx != nil {
		n.sx.toSliceNodo(s, i)
	}
	n.sx = nil
	s[*i] = n
	*i++
	if n.dx != nil {
		n.dx.toSliceNodo(s, i)
	}
	n.dx = nil
}

func (n *nodo) inserisciSx(values []*nodo, min, max int) {
	nel := max - min
	if nel == 0 {
		n.sx = values[min]
	} else if nel == 1 {
		n.sx = values[max]
		n.sx.inserisciSx(values, min, min)
	} else {
		media := (max + min) / 2
		n.sx = values[media]
		n.sx.inserisciSx(values, min, media-1)
		n.sx.inserisciDx(values, media+1, max)
	}
}

func (n *nodo) inserisciDx(values []*nodo, min, max int) {
	nel := max - min
	if nel == 0 {
		n.dx = values[min]
	} else if nel == 1 {
		n.dx = values[max]
		n.dx.inserisciSx(values, min, min)
	} else {
		media := (max + min) / 2
		n.dx = values[media]
		n.dx.inserisciSx(values, min, media-1)
		n.dx.inserisciDx(values, media+1, max)
	}
}
