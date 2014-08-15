package unionfind

import "fmt"

type QF struct {
	id    []int
	count int
}

func New(n int) *QF {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	return &QF{a, len(a)}
}

func (qf *QF) String() string {
	return fmt.Sprintf("id: %v count: %d", qf.id, qf.count)
}

func (qf *QF) Count() int {
	return qf.count
}

func (qf *QF) Union(q, p int) {
	pid := qf.id[p]
	qid := qf.id[q]
	if !qf.Connected(q, p) {
		qf.count--
		for i := 0; i < len(qf.id); i++ {
			if qf.id[i] == pid {
				qf.id[i] = qid
			}
		}
	}
}

func (qf *QF) Connected(q, p int) bool {
	return qf.id[q] == qf.id[p]
}

func (qf *QF) Find(p int) int {
	return qf.id[p]
}
