package game

import (
	"errors"
	"fmt"
)

type Axis uint8

const (
	AxisOut Axis = 0xF // > (0..8)
)

type Coordinates struct {
	X, Y Axis
}

type Move struct {
	From, To Coordinates
	//Figure   FigureType
}

func NewMove(oldX, oldY, newX, newY Axis) *Move {
	m := &Move{
		From: Coordinates{
			X: oldX,
			Y: oldY,
		},
		To: Coordinates{
			X: newX,
			Y: newY,
		},
	}
	return m
}

func (m *Move) String() string {
	return fmt.Sprintf(
		"{From:{X: %d, Y: %d}, To:{X: %d, Y: %d}}",
		m.From.X,
		m.From.Y,
		m.To.X,
		m.To.Y,
	)
}

func getPawnDirections(f *Figure) []Coordinates {
	var targets []Coordinates

	//newX := f.Position.X // same
	newY := f.Position.Y + 1

	if err := validateRange(newY); errors.Is(err, ErrAxisOutOfRange) {
		return nil // no new possible coordinates to move
	}

	targets = append(targets, Coordinates{
		X: f.Position.X,
		Y: newY,
	})

	return targets
}

func getLanceDirections(f *Figure) []Coordinates {
	var targets []Coordinates

	//newX := f.Position.X // same
	for y := f.Position.Y + 1; y < BoardSize; y++ {
		targets = append(targets, Coordinates{
			X: f.Position.X,
			Y: y,
		})
	}

	return targets
}

func getKnightDirections(f *Figure) []Coordinates {
	var targets []Coordinates
	var newX, newY Axis

	// Left "L"
	newX = f.Position.X - 1
	newY = f.Position.Y + 2
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Right "L"
	newX = f.Position.X + 1
	newY = f.Position.Y + 2
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	return targets
}

func GetFigureDirections(f *Figure) []Coordinates {
	// TODO: ensure that directions are stopped on "collision"
	// TODO: coordinates are opposite/inverted for different figure 'color'

	switch f.Type {
	case Pawn:
		return getPawnDirections(f)

	case Lance:
		return getLanceDirections(f)

	case Knight:
		return getKnightDirections(f)

	default:
		return nil
	}
}
