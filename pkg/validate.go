package game

import (
	"fmt"
)

func validateRange(v Axis) error {
	//if v >= 0 && v < BoardSize {
	if v < BoardSize {
		return nil
	}
	return ErrAxisOutOfRange
}

func validateDirection(f *Figure, m *Move) error {
	match := false
	inverted := false // FIXME
	possibleMoves := GetFigureDirections(f, inverted)

	for _, target := range possibleMoves {
		if target.X == m.To.X && target.Y == m.To.Y {
			match = true
			break
		}
	}

	if !match {
		return ErrInvalidMove
	}

	return nil
}

func validateOwner(f *Figure) {
	if f.Owner != White && f.Owner != Black {
		panic(ErrImplementationIssue)
	}
}

func (g *Game) validateMove(m *Move) (*Figure, *Figure, error) {
	var err error

	//fmt.Printf("# %s\n", m) // DEBUG

	// Coordinate checks
	if err = validateRange(m.From.X); err != nil {
		return nil, nil, err
	}

	if err = validateRange(m.From.Y); err != nil {
		return nil, nil, err
	}
	if err = validateRange(m.To.X); err != nil {
		return nil, nil, err
	}

	if err = validateRange(m.To.Y); err != nil {
		return nil, nil, err
	}

	f1, err := g.GetFigure(m.From.X, m.From.Y)
	if err != nil {
		panic(err)
	}
	if f1 != nil {
		validateOwner(f1)
	}

	f2, err := g.GetFigure(m.To.X, m.To.Y)
	if err != nil {
		panic(err)
	}
	if f2 != nil {
		validateOwner(f2)
	}

	// Validate turn owner
	// TODO: if target same, then can't move - otherwise take it

	if f1.Owner != g.CurrentTurnOwner {
		fmt.Printf("Not my turn: %s\n", f1)
		return nil, nil, ErrInvalidTurnOwner
	}

	if err = validateDirection(f1, m); err != nil {
		return nil, nil, err
	}

	return f1, f2, nil
}
