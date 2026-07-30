package service

type GameService interface {
	NextTurn() GameService
	ValidationGameBoard() bool
	IsGameOver() bool
}
