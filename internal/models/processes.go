package models

import "github.com/google/uuid"

type Processes map[uuid.UUID][]uuid.UUID

func NewProcesses() Processes {
	return make(map[uuid.UUID][]uuid.UUID)
}

func (ps Processes) Add(fID, qID uuid.UUID) {
	ql, ok := ps[fID]
	if !ok {
		ql = make([]uuid.UUID, 0, 1)
	}

	ql = append(ql, qID)

	ps[fID] = ql
}
