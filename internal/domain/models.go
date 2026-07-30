package domain

import(
	"github.com/google/uuid"
)

type GameBoard struct {
	var board [][]int
}

type CurrentGame struct {
	GameBoard
	id uuid.UUID
}