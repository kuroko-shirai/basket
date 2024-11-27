package models

type Processes map[int][]int

func NewProcesses() Processes {
	return make(map[int][]int)
}

func (ps Processes) Add(fID, qID int) {
	ql, ok := ps[fID]
	if !ok {
		ql = make([]int, 0, 1)
	}

	ql = append(ql, qID)

	ps[fID] = ql
}
