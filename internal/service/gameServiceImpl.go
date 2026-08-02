package service

import (
	"TicTacToe/internal/domain"
	"errors"
	"math"
)

type gameService struct{}

func (g gameService) NextTurn(cg domain.CurrentGame) (int, int, error) {
	nextCurrentPlayer := domain.FirstPlayer
	var maxCellPoints = struct {
		points int
		x, y   int
	}{math.MinInt, 0, 0}

	for i := 0; i < domain.BoardSize; i++ {
		for j := 0; j < domain.BoardSize; j++ {

			if val, _ := cg.Get(i, j); val != domain.FreeCell {
				continue
			}

			newCg := cg
			if cg.GetCurrentPlayer() == domain.FirstPlayer {

				if err := newCg.Set(i, j, domain.FirstPlayerOccupied); err != nil {
					return 0, 0, err
				}

				nextCurrentPlayer = domain.SecondPlayer

			} else {
				if err := newCg.Set(i, j, domain.SecondPlayerOccupied); err != nil {
					return 0, 0, err
				}

				nextCurrentPlayer = domain.FirstPlayer
			}

			mx := g.miniMax(newCg, nextCurrentPlayer)
			if maxCellPoints.points < mx {
				maxCellPoints.points = mx
				maxCellPoints.x, maxCellPoints.y = i, j
			}
		}
	}

	return maxCellPoints.x, maxCellPoints.y, nil
}

const CountOfPointsForWin = 10

func (g gameService) miniMax(cg domain.CurrentGame, currentPlayer domain.Player) int {

	otherPlayer := func(pl domain.Player) domain.Player {
		if pl == domain.FirstPlayer {
			return domain.SecondPlayer
		}

		return domain.FirstPlayer
	}

	turnPlayer := func(pl domain.Player) int {
		if pl == domain.FirstPlayer {
			return domain.FirstPlayerOccupied
		}

		return domain.SecondPlayerOccupied
	}

	if w, ok := g.IsGameOver(cg); ok {
		if w == Winner(Draw) {
			return 0
		}

		if w == Winner(First) && cg.GetComputerPlayer() == domain.FirstPlayer ||
			w == Winner(Second) && cg.GetComputerPlayer() == domain.SecondPlayer {

			return CountOfPointsForWin
		}

		return -CountOfPointsForWin
	}

	if currentPlayer == cg.GetComputerPlayer() {
		best := math.MinInt

		for i := 0; i < domain.BoardSize; i++ {
			for j := 0; j < domain.BoardSize; j++ {
				if val, _ := cg.Get(i, j); val != domain.FreeCell {
					continue
				}

				newCg := cg
				newCg.Set(i, j, turnPlayer(currentPlayer))
				score := g.miniMax(newCg, otherPlayer(currentPlayer))
				best = max(best, score)
			}
		}
		return best

	} else {
		best := math.MaxInt
		for i := 0; i < domain.BoardSize; i++ {
			for j := 0; j < domain.BoardSize; j++ {
				if val, _ := cg.Get(i, j); val != domain.FreeCell {
					continue
				}

				newCg := cg
				newCg.Set(i, j, turnPlayer(currentPlayer))
				score := g.miniMax(newCg, otherPlayer(currentPlayer))
				best = min(best, score)
			}
		}
		return best
	}

	return 0
}

func (g gameService) ValidateNextTurn(oldCg, nextCg domain.CurrentGame) error {
	countChangedCells := 0
	countFirstPlayerCells := 0
	countSecondPlayerCells := 0
	replaceNonFreeCells := 0

	for i := 0; i < domain.BoardSize; i++ {
		for j := 0; j < domain.BoardSize; j++ {

			oldCell, _ := oldCg.Get(i, j)
			newCell, _ := nextCg.Get(i, j)

			if oldCell != newCell {
				countChangedCells++

				if oldCell != domain.FreeCell {
					replaceNonFreeCells++
				}
			}

			if newCell == domain.FirstPlayerOccupied {
				countFirstPlayerCells++
			} else if newCell == domain.SecondPlayerOccupied {
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

	if replaceNonFreeCells > 0 {
		return errors.New("modified non free cells")
	}

	if math.Abs(float64(countFirstPlayerCells)-float64(countSecondPlayerCells)) > 1 {

		if countFirstPlayerCells > countSecondPlayerCells {
			return errors.New("it should be the second player's turn now")
		}

		return errors.New("it should be the first player's turn now")
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

		if firstCell == domain.FreeCell {
			continue
		}

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

		if firstCell == domain.FreeCell {
			continue
		}

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

	if firstCell != domain.FreeCell {

		for i := 0; i < domain.BoardSize; i++ {
			if cell, _ := cg.Get(i, i); cell != firstCell {
				w = Winner(Draw)
				break
			}
		}

		if w != Winner(Draw) {
			return w
		}
	}

	firstCell, _ = cg.Get(0, domain.BoardSize-1)
	w = Winner(firstCell)

	if firstCell != domain.FreeCell {

		for i := 0; i < domain.BoardSize; i++ {
			if cell, _ := cg.Get(i, domain.BoardSize-i-1); cell != firstCell {
				w = Winner(Draw)
				break
			}
		}
	}

	return w
}
