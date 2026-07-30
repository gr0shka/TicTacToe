package domain

import (
	"github.com/google/uuid"
)

type CurrentGame struct {
	GameBoard
	id uuid.UUID
}

func (cg CurrentGame) GetID() uuid.UUID {
	return cg.id
}

func (cg *CurrentGame) SetID(id uuid.UUID) {
	cg.id = id
}
