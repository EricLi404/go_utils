package go_utils

type QuickFind struct {
	pair   []APairOfInt
	ele    map[int]int
	unions map[int][]int
	max    int
}

func (q *QuickFind) initEleList() {
	q.ele = make(map[int]int, q.max+1)
	q.unions = make(map[int][]int, q.max+1)
	for i := 0; i < q.max; i++ {
		q.ele[i] = i + 1
	}
}

func (q *QuickFind) do() {
	for _, vv := range q.pair {
		pN := APairOfInt{Key: q.ele[vv.Key], Value: q.ele[vv.Value]}
		minV := IntMin(pN.Key, pN.Value)
		for k, _ := range q.ele {
			if q.ele[k] == pN.Key || q.ele[k] == pN.Value {
				q.ele[k] = minV
			}
		}
	}
	for k, v := range q.ele {
		q.unions[v] = append(q.unions[v], k)
	}
}

func (q *QuickFind) Result() map[int]int {
	return q.ele
}

func (q *QuickFind) GetUnions() map[int][]int {
	return q.unions
}

func (q *QuickFind) GetUnionsByLimit(limit int) map[int][]int {
	rl := make(map[int][]int, len(q.unions))
	for k, v := range q.unions {
		if len(v) >= limit {
			rl[k] = v
		}
	}
	return rl
}

func NewQuickFind(max int, pair []APairOfInt) *QuickFind {
	qf := new(QuickFind)
	qf.pair = pair
	qf.max = max
	qf.initEleList()
	qf.do()
	return qf
}
