package unionfind

import "fmt"

type QU struct {
	id    []int
	count int
}

func NewQU(n int) *QU {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	return &QU{a, len(a)}
}

func (qu *QU) String() string {
	return fmt.Sprintf("id: %v count: %d", qu.id, qu.count)
}

func (qu *QU) Count() int {
  return qu.count
}

func (qu *QU) Connected(q, p int) bool {
  qid, pid := qu.Find(q), qu.Find(p)
  return qid == pid
}

func (qu *QU) Find(p int) int {
  p1 := qu.id[p]
  for p1 != p {
    p, p1 = p1, qu.id[p1]
  }
  return p
}

func (qu *QU) Union(p, q int) {
  pid, qid := qu.Find(p), qu.Find(q)
  if pid != qid {
      qu.id[pid] = qid
      qu.count--
  }
}
