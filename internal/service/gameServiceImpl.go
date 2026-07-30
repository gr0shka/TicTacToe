package service

import "TicTacToe/internal/domain"

type gameService struct {
	domain.CurrentGame
}

func (g gameService) NextTurn(cg domain.CurrentGame) (i, j int, err error) {
	//TODO implement me
	panic("implement me")
}

func (g gameService) ValidateGameBoard(gb1, gb2 domain.CurrentGame) (err error) {
	//TODO implement me
	panic("implement me")
}

func (g gameService) IsGameOver(cg domain.CurrentGame) (w Winner, isEnd bool) {
	//TODO implement me
	panic("implement me")
}
