package service

import "TicTacToe/internal/domain"

type GameService interface {
	NextTurn() domain.GameBoard
	ValidationGameBoard() bool
	IsGameOver() bool
}
