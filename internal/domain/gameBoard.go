package domain

import "errors"

const BoardSize = 3

const (
	FreeCell int = iota
	FirstPlayerOccupied
	SecondPlayerOccupied
)

type GameBoard struct {
	board [BoardSize][BoardSize]int
}

func (gm *GameBoard) Set(i, j, value int) error {
	if value != 0 && value != 1 {
		return errors.New("invalid value")
	}

	if i < 0 || i >= BoardSize || j < 0 || j >= BoardSize {
		return errors.New("invalid coordinates")
	}

	gm.board[i][j] = value

	return nil
}

func (gm *GameBoard) Get(i, j int) (value int, err error) {
	if i < 0 || i >= BoardSize || j < 0 || j >= BoardSize {
		return 0, errors.New("invalid coordinates")
	}

	return gm.board[i][j], nil
}
