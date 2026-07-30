package domain

import (
	"github.com/google/uuid"
)

type Player int

const (
	FirstPlayer Player = iota
	SecondPlayer
)

type CurrentGame struct {
	GameBoard
	currentPlayer  Player
	computerPlayer Player
	id             uuid.UUID
}

func (cg CurrentGame) GetID() uuid.UUID {
	return cg.id
}

func (cg *CurrentGame) SetID(id uuid.UUID) {
	cg.id = id
}

func (cg CurrentGame) GetCurrentPlayer() Player {
	return cg.currentPlayer
}

func (cg CurrentGame) GetComputerPlayer() Player {
	return cg.computerPlayer
}
