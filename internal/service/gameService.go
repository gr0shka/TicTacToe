package service

import (
	"TicTacToe/internal/domain"
)

type Winner int

const (
	First Winner = iota
	Second
	Draw
)

type GameService interface {
	NextTurn(cg domain.CurrentGame) (int, int, error)
	ValidateNextTurn(oldCg, nextCg domain.CurrentGame) error
	IsGameOver(cg domain.CurrentGame) (Winner, bool)
}
