package service

import (
	"TicTacToe/internal/domain"
	"errors"
	"math"
)

type gameService struct{}

func (g gameService) NextTurn(cg domain.CurrentGame) (int, int, error) {
	//TODO implement me
	panic("implement me")
}

func (g gameService) ValidateNextTurn(oldCg, nextCg domain.CurrentGame) error {
	countChangedCells := 0
	countFirstPlayerCells := 0
	countSecondPlayerCells := 0

	for i := 0; i < domain.BoardSize; i++ {
		for j := 0; j < domain.BoardSize; j++ {

			oldCell, _ := oldCg.Get(i, j)
			newCell, _ := nextCg.Get(i, j)

			if oldCell != newCell {
				countChangedCells++
			}

			if newCell == domain.FirstPlayerOccupied {
				countFirstPlayerCells++
			} else {
				countSecondPlayerCells++
			}
		}
	}

	if countChangedCells > 1 {
		return errors.New("more than 1 cell modified")
	}
	if countChangedCells < 1 {
		return errors.New("no changes in board")
	}

	if math.Abs(float64(countFirstPlayerCells)-float64(countSecondPlayerCells)) > 1 {

		if countFirstPlayerCells > countSecondPlayerCells {
			return errors.New("it should be the second player's turn now")
		} else {
			return errors.New("it should be the first player's turn now")
		}
	}

	return nil
}

func (g gameService) IsGameOver(cg domain.CurrentGame) (Winner, bool) {
	if w := checkHorizontalLines(cg); w != Winner(Draw) {
		return w, true
	}

	if w := checkVerticalLines(cg); w != Winner(Draw) {
		return w, true
	}

	if w := checkDiagonalLines(cg); w != Winner(Draw) {
		return w, true
	}

	isEnd := true
	for i := 0; i < domain.BoardSize; i++ {
		for j := 0; j < domain.BoardSize; j++ {
			if v, _ := cg.Get(i, j); v == domain.FreeCell {
				isEnd = false
			}
		}
	}

	return Draw, isEnd
}

func checkHorizontalLines(cg domain.CurrentGame) Winner {
	w := Winner(Draw)

	for i := 0; i < domain.BoardSize; i++ {
		firstCell, _ := cg.Get(i, 0)
		w = Winner(firstCell)

		for j := 0; j < domain.BoardSize; j++ {
			if cell, _ := cg.Get(i, j); cell != firstCell {
				w = Winner(Draw)
				break
			}
		}
	}

	return w
}

func checkVerticalLines(cg domain.CurrentGame) Winner {
	w := Winner(Draw)

	for j := 0; j < domain.BoardSize; j++ {
		firstCell, _ := cg.Get(0, j)
		w = Winner(firstCell)

		for i := 0; i < domain.BoardSize; i++ {
			if cell, _ := cg.Get(i, j); cell != firstCell {
				w = Winner(Draw)
				break
			}
		}
	}

	return w
}

func checkDiagonalLines(cg domain.CurrentGame) Winner {
	firstCell, _ := cg.Get(0, 0)
	w := Winner(firstCell)

	for i := 0; i < domain.BoardSize; i++ {
		if cell, _ := cg.Get(i, i); cell != firstCell {
			w = Winner(Draw)
			break
		}
	}

	if w != Winner(Draw) {
		return w
	}

	firstCell, _ = cg.Get(domain.BoardSize-1, domain.BoardSize-1)
	w = Winner(firstCell)

	for i := domain.BoardSize - 1; i >= 0; i-- {
		if cell, _ := cg.Get(i, i); cell != firstCell {
			w = Winner(Draw)
			break
		}
	}

	return w
}
