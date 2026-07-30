package service

import (
	"TicTacToe/internal/domain"
)

type Winner int

const (
	Draw Winner = iota
	First
	Second
)

type GameService interface {
	NextTurn(cg domain.CurrentGame) (i, j int, err error)
	ValidateGameBoard(gb1, gb2 domain.CurrentGame) (err error)
	IsGameOver(cg domain.CurrentGame) (w Winner, isEnd bool)
}
