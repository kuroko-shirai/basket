package storage

type processes map[int][]int

func newProcesses() processes {
	return make(map[int][]int)
}

func (ftq processes) Add(fID, qID int) {
	ql, ok := ftq[fID]
	if !ok {
		ql = make([]int, 0, 1)
	}

	ql = append(ql, qID)

	ftq[fID] = ql
}
