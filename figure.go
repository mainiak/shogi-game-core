package game

import (
	"fmt"
)

type FigureType uint8

const (
	Pawn          FigureType = iota + 1 // start with 1, not 0
	Lance                               //FigureType
	Knight                              //FigureType
	Rook                                //FigureType
	Bishop                              //FigureType
	SilverGeneral                       //FigureType
	GoldenGeneral                       //FigureType
	King                                //FigureType
)

func (ft FigureType) String() string {
	return fmt.Sprintf("%s | %s | %s", ft.NameShort(), ft.NameKanji(), ft.NameEn())
}

func (ft FigureType) NameShort() string {
	switch ft {

	case Pawn:
		return "P"

	case Lance:
		return "L"

	case Knight:
		return "N"

	case Rook:
		return "R"

	case Bishop:
		return "B"

	case SilverGeneral:
		return "S"

	case GoldenGeneral:
		return "G"

	case King:
		return "K"

	default:
		return "Error"
	}
}

func (ft FigureType) NameKanji() string {
	switch ft {

	case Pawn:
		return "歩兵"

	case Lance:
		return "香車"

	case Knight:
		return "桂馬"

	case Rook:
		return "飛車"

	case Bishop:
		return "角行"

	case SilverGeneral:
		return "銀将"

	case GoldenGeneral:
		return "金将"

	case King:
		return "王将"

	default:
		return "Error"
	}
}

func (ft FigureType) NameEn() string {
	switch ft {

	case Pawn:
		return "Pawn"

	case Lance:
		return "Lance"

	case Knight:
		return "Knight"

	case Rook:
		return "Rook"

	case Bishop:
		return "Bishop"

	case SilverGeneral:
		return "Silver General"

	case GoldenGeneral:
		return "Golden General"

	case King:
		return "King"

	default:
		return "Error"
	}
}

// 0..39 for 2x 20 figures
type FigureId uint8

const (
	NotFigureId FigureId = 0xFF // > (0..39)
)

func (fId FigureId) String() string {
	return fmt.Sprintf("%02d", fId)
}

type FigureOwner uint8

const (
	White FigureOwner = 1
	Black FigureOwner = 2
)

func (fo FigureOwner) String() string {
	if fo == White {
		return "White"
	}

	if fo == Black {
		return "Black"
	}

	return "Error"
}

type Figure struct {
	Id       FigureId
	Type     FigureType
	Owner    FigureOwner
	Position Coordinates
}

func NewFigure(id FigureId, owner FigureOwner, ft FigureType) *Figure {
	f := &Figure{
		Id:    id,
		Type:  ft,
		Owner: owner,
		Position: Coordinates{
			// Not on board by default
			X: AxisOut,
			Y: AxisOut,
		},
	}
	return f
}

// TODO: return error and use testCoordinates()
func (f *Figure) SetCoordinates(x, y Axis) {
	f.Position.X = x
	f.Position.Y = y
}

func (f *Figure) MoveOut() {
	f.Position.X = AxisOut
	f.Position.Y = AxisOut
}

func (f *Figure) String() string {
	return fmt.Sprintf("%s:%s:(%d,%d):%s", f.Id, f.Owner, f.Position.X, f.Position.Y, f.Type)
}

func (f *Figure) Short() string {
	if f.Owner == White {
		return fmt.Sprintf("w%s", f.Type.NameShort())
	}

	if f.Owner == Black {
		return fmt.Sprintf("b%s", f.Type.NameShort())
	}

	return fmt.Sprintf("?%s", f.Type.NameShort())
}
