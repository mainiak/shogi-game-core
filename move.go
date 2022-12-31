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

func getPawnDirections(f *Figure, inverted bool) []Coordinates {
	var newY Axis
	var targets []Coordinates

	//newX := f.Position.X // same
	if inverted {
		newY = f.Position.Y - 1
	} else {
		newY = f.Position.Y + 1
	}

	if err := validateRange(newY); errors.Is(err, ErrAxisOutOfRange) {
		return nil // no new possible coordinates to move
	}

	targets = append(targets, Coordinates{
		X: f.Position.X,
		Y: newY,
	})

	return targets
}

func getLanceDirections(f *Figure, inverted bool) []Coordinates {
	var newY Axis
	var targets []Coordinates

	//newX := f.Position.X // same
	if inverted {
		for newY = f.Position.Y - 1; newY >= 0; newY-- {
			targets = append(targets, Coordinates{
				X: f.Position.X,
				Y: newY,
			})

			if newY == 0 {
				break
			}
		}
	} else {
		for newY = f.Position.Y + 1; newY < BoardSize; newY++ {
			targets = append(targets, Coordinates{
				X: f.Position.X,
				Y: newY,
			})
		}
	}

	return targets
}

func getKnightDirections(f *Figure, inverted bool) []Coordinates {
	var targets []Coordinates
	var newX, newY Axis

	// Left "L"
	newX = f.Position.X - 1
	if inverted {
		newY = f.Position.Y - 2
	} else {
		newY = f.Position.Y + 2
	}
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Right "L"
	newX = f.Position.X + 1
	if inverted {
		newY = f.Position.Y - 2
	} else {
		newY = f.Position.Y + 2
	}
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	return targets
}

func getRookDirections(f *Figure) []Coordinates {
	var targets []Coordinates

	// UP
	//newX := f.Position.X // same
	for y := f.Position.Y + 1; y < BoardSize; y++ {
		targets = append(targets, Coordinates{
			X: f.Position.X,
			Y: y,
		})
	}

	// Down
	//newX := f.Position.X // same
	for y := f.Position.Y - 1; y >= 0; y-- {
		targets = append(targets, Coordinates{
			X: f.Position.X,
			Y: y,
		})

		if y == 0 {
			break
		}
	}

	// Left
	//newY := f.Position.Y // same
	for x := f.Position.X - 1; x >= 0; x-- {
		targets = append(targets, Coordinates{
			X: x,
			Y: f.Position.Y,
		})

		if x == 0 {
			break
		}
	}

	// Right
	//newY := f.Position.Y // same
	for x := f.Position.X + 1; x < BoardSize; x++ {
		targets = append(targets, Coordinates{
			X: x,
			Y: f.Position.Y,
		})
	}

	return targets
}

func getBishopDirections(f *Figure) []Coordinates {
	var i, newX, newY Axis
	var targets []Coordinates

	// UP Right
	for i = 0; i < BoardSize; i++ {
		newX = f.Position.X + i
		newY = f.Position.Y + i
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// Down Right
	for i = 0; i < BoardSize; i++ {
		newX = f.Position.X + i
		newY = f.Position.Y - i
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// Down Left
	for i = 0; i < BoardSize; i++ {
		newX = f.Position.X - i
		newY = f.Position.Y - i
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// UP Left
	for i = 0; i < BoardSize; i++ {
		newX = f.Position.X - i
		newY = f.Position.Y + i
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	return targets
}

func getSilverGeneralDirections(f *Figure, inverted bool) []Coordinates {
	var targets []Coordinates
	var newX, newY Axis

	// UP
	if inverted == false {
		newX = f.Position.X
		newY = f.Position.Y + 1
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// UP Right
	newX = f.Position.X + 1
	newY = f.Position.Y + 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Down Right
	newX = f.Position.X + 1
	newY = f.Position.Y - 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Down
	if inverted {
		newX = f.Position.X
		newY = f.Position.Y - 1
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// Down Left
	newX = f.Position.X - 1
	newY = f.Position.Y - 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// UP Left
	newX = f.Position.X - 1
	newY = f.Position.Y + 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	return targets
}

func getGoldenGeneralDirections(f *Figure, inverted bool) []Coordinates {
	var targets []Coordinates
	var newX, newY Axis

	// UP
	newX = f.Position.X
	newY = f.Position.Y + 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// UP Right
	if inverted == false {
		newX = f.Position.X + 1
		newY = f.Position.Y + 1
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// Right
	newX = f.Position.X + 1
	newY = f.Position.Y
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Down Right
	if inverted {
		newX = f.Position.X + 1
		newY = f.Position.Y - 1
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// Down
	newX = f.Position.X
	newY = f.Position.Y - 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Down Left
	if inverted {
		newX = f.Position.X - 1
		newY = f.Position.Y - 1
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	// Left
	newX = f.Position.X - 1
	newY = f.Position.Y
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// UP Left
	if inverted == false {
		newX = f.Position.X - 1
		newY = f.Position.Y + 1
		if validateRange(newX) == nil && validateRange(newY) == nil {
			targets = append(targets, Coordinates{
				X: newX,
				Y: newY,
			})
		}
	}

	return targets
}

func getKingDirections(f *Figure) []Coordinates {
	var targets []Coordinates
	var newX, newY Axis

	// UP
	newX = f.Position.X
	newY = f.Position.Y + 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// UP Right
	newX = f.Position.X + 1
	newY = f.Position.Y + 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Right
	newX = f.Position.X + 1
	newY = f.Position.Y
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Down Right
	newX = f.Position.X + 1
	newY = f.Position.Y - 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Down
	newX = f.Position.X
	newY = f.Position.Y - 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Down Left
	newX = f.Position.X - 1
	newY = f.Position.Y - 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// Left
	newX = f.Position.X - 1
	newY = f.Position.Y
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	// UP Left
	newX = f.Position.X - 1
	newY = f.Position.Y + 1
	if validateRange(newX) == nil && validateRange(newY) == nil {
		targets = append(targets, Coordinates{
			X: newX,
			Y: newY,
		})
	}

	return targets
}

func GetFigureDirections(f *Figure, inverted bool) []Coordinates {
	// TODO: ensure that directions are stopped on "collision"
	// TODO: coordinates are opposite/inverted for different figure 'color'

	switch f.Type {
	case Pawn:
		return getPawnDirections(f, inverted)

	case Lance:
		return getLanceDirections(f, inverted)

	case Knight:
		return getKnightDirections(f, inverted)

	case Rook:
		return getRookDirections(f)

	case Bishop:
		return getBishopDirections(f)

	case SilverGeneral:
		return getSilverGeneralDirections(f, inverted)

	case GoldenGeneral:
		return getGoldenGeneralDirections(f, inverted)

	case King:
		return getKingDirections(f)

	default:
		return nil
	}
}
