package hashset

type IntHashSet map[int]struct{}

func (i IntHashSet) Add(n int) {
	i[n] = struct{}{}
}

func (i IntHashSet) Remove(n int) {
	delete(i, n)
}

func (i IntHashSet) Has(n int) bool {
	_, ok := i[n]
	return ok
}
