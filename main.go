package main

import "fmt"

// UnionFind keeps the items in disjoint sets.
// It provides 3 operations:
// * Union - unions two sets
// * Same - checks whether two sets have the same representative
// * Find - finds the representative for given item
type UnionFind struct {
	parents []uint64
	ranks   []uint64
}

func (uf UnionFind) String() string {
	return fmt.Sprintf("%v\n%v", uf.parents, uf.ranks)
}

func New(size int) *UnionFind {
	uf := &UnionFind{
		parents: make([]uint64, size),
		ranks:   make([]uint64, size),
	}

	for i := 0; i < size; i++ {
		uf.parents[i] = uint64(i)
	}

	return uf
}

func (u *UnionFind) Union(a, b uint64) {
	a = u.Find(a)
	b = u.Find(b)

	// TODO ranks

	u.parents[b] = a
}

func (u *UnionFind) Same(a, b uint64) bool {
	return u.Find(a) == u.Find(b)
}

func (u *UnionFind) Find(a uint64) uint64 {
	if u.parents[a] == a {
		return a
	} else {
		// compress path
		u.parents[a] = u.Find(u.parents[a])
		return u.parents[a]
	}
}

func main() {
	uf := New(8)
	fmt.Println(uf)

	uf.Union(0, 1)
	uf.Union(2, 3)

	fmt.Println(uf)

	fmt.Println(uf.Find(3))
	fmt.Println(uf.Same(0, 1))
}
