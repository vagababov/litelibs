package unionfind

type UnionFind interface {
	Union(p, q int)
	Connected(p, q int) bool
	Find(p int) int
	Count() int
}
