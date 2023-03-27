package game

import (
	"errors"
	"fmt"
)

const BoardSize = 9    // 9x9
const FigureCount = 40 // 2x 20

type Game struct {
	Board            [BoardSize][BoardSize]FigureId
	Figures          [FigureCount]*Figure
	CurrentTurnOwner FigureOwner
}

func NewGame() *Game {
	var x, y Axis
	game := &Game{
		CurrentTurnOwner: White,
	}

	for x = 0; x < BoardSize; x++ {
		for y = 0; y < BoardSize; y++ {
			game.Board[x][y] = NotFigureId
		}
	}

	for i := range game.Figures {
		game.Figures[i] = nil
	}

	game.initFigures()
	return game
}

func (g *Game) createFigure(figureId FigureId, owner FigureOwner, figureType FigureType) {
	f := NewFigure(figureId, owner, figureType)
	g.Figures[figureId] = f
}

func (g *Game) PlaceFigure(x, y Axis, figureId FigureId) {
	// TODO: RemoveFigure() or check there is another figure
	f := g.Figures[figureId]
	f.SetCoordinates(x, y)
	g.Figures[figureId] = f
	g.Board[x][y] = figureId
}

func (g *Game) GetFigure(x, y Axis) (*Figure, error) {
	var err error

	if err = validateRange(x); errors.Is(err, ErrAxisOutOfRange) {
		panic(err)
	}

	if err = validateRange(y); errors.Is(err, ErrAxisOutOfRange) {
		panic(err)
	}

	figureId := g.Board[x][y]

	if figureId == NotFigureId {
		return nil, nil // should be nil anyway, but meh
	}

	f := g.Figures[figureId]
	return f, nil
}

func (g *Game) CleanBoard() {
	var x, y Axis

	for x = 0; x < BoardSize; x++ {
		for y = 0; y < BoardSize; y++ {
			g.Board[x][y] = NotFigureId
		}
	}

	for i := range g.Figures {
		g.Figures[i].MoveOut()
	}
}

// Should be called only once when new instance of Game is created
func (g *Game) initFigures() {
	// create all figures
	var figureId FigureId = 0

	// Black (at bottom) moves first.

	// Black Pawn front
	for x := 0; x < BoardSize; x++ {
		g.createFigure(figureId, Black, Pawn)
		figureId++
	}

	// Black Bishop
	g.createFigure(figureId, Black, Bishop)
	figureId++

	// Black Rook
	g.createFigure(figureId, Black, Rook)
	figureId++

	// L N S G K G S N L

	// Black Lance
	g.createFigure(figureId, Black, Lance)
	figureId++

	// Black Knight
	g.createFigure(figureId, Black, Knight)
	figureId++

	// Black Silver General
	g.createFigure(figureId, Black, SilverGeneral)
	figureId++

	// Black Golden General
	g.createFigure(figureId, Black, GoldenGeneral)
	figureId++

	// Black King
	g.createFigure(figureId, Black, King)
	figureId++

	// Black Golden General
	g.createFigure(figureId, Black, GoldenGeneral)
	figureId++

	// Black Silver General
	g.createFigure(figureId, Black, SilverGeneral)
	figureId++

	// Black Knight
	g.createFigure(figureId, Black, Knight)
	figureId++

	// Black Lance
	g.createFigure(figureId, Black, Lance)
	figureId++

	// White Pawn front
	for x := 0; x < BoardSize; x++ {
		g.createFigure(figureId, White, Pawn)
		figureId++
	}

	// White Rook
	g.createFigure(figureId, White, Rook)
	figureId++

	// White Bishop
	g.createFigure(figureId, White, Bishop)
	figureId++

	// L N S G K G S N L

	// White Lance
	g.createFigure(figureId, White, Lance)
	figureId++

	// White Knight
	g.createFigure(figureId, White, Knight)
	figureId++

	// White Silver General
	g.createFigure(figureId, White, SilverGeneral)
	figureId++

	// White Golden General
	g.createFigure(figureId, White, GoldenGeneral)
	figureId++

	// White King
	g.createFigure(figureId, White, King)
	figureId++

	// White Golden General
	g.createFigure(figureId, White, GoldenGeneral)
	figureId++

	// White Silver General
	g.createFigure(figureId, White, SilverGeneral)
	figureId++

	// White Knight
	g.createFigure(figureId, White, Knight)
	figureId++

	// White Lance
	g.createFigure(figureId, White, Lance)
	figureId++
}

func (g *Game) SetDefaultBoard() {
	// create all figures
	var figureId FigureId
	var x, y Axis

	g.CleanBoard()

	figureId = 0
	x, y = 0, 2

	// Black Pawn front
	for ; x < BoardSize; x++ {
		g.PlaceFigure(x, y, figureId)
		figureId++
	}

	// Black Bishop
	x, y = 1, 1
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Rook
	x, y = 7, 1
	g.PlaceFigure(x, y, figureId)
	figureId++

	// L N S G K G S N L

	// Black Lance
	x, y = 0, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Knight
	x, y = 1, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Silver General
	x, y = 2, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Golden General
	x, y = 3, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black King
	x, y = 4, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Golden General
	x, y = 5, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Silver General
	x, y = 6, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Knight
	x, y = 7, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// Black Lance
	x, y = 8, 0
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Pawn front
	x, y = 0, 6
	for ; x < BoardSize; x++ {
		g.PlaceFigure(x, y, figureId)
		figureId++
	}

	// White Rook
	x, y = 1, 7
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Bishop
	x, y = 7, 7
	g.PlaceFigure(x, y, figureId)
	figureId++

	// L N S G K G S N L

	// White Lance
	x, y = 0, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Knight
	x, y = 1, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Silver General
	x, y = 2, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Golden General
	x, y = 3, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White King
	x, y = 4, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Golden General
	x, y = 5, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Silver General
	x, y = 6, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Knight
	x, y = 7, 8
	g.PlaceFigure(x, y, figureId)
	figureId++

	// White Lance
	x, y = 8, 8
	g.PlaceFigure(x, y, figureId)
	figureId++
}

func (g *Game) Move(oldX, oldY, newX, newY Axis) error {
	var err error
	var f1, f2 *Figure

	fmt.Printf("# (%d,%d) -> (%d,%d)\n", oldX, oldY, newX, newY)
	m := NewMove(oldX, oldY, newX, newY)

	f1, f2, err = g.validateMove(m)
	if err != nil {
		return err
	}

	if f1 == nil {
		fmt.Printf("No figure\n")
		return ErrInvalidMove
	}

	fmt.Printf("# %s\n", f1)
	fmt.Printf("# %s\n", f2)

	// "Flip" turn owner
	if f1.Owner == White {
		g.CurrentTurnOwner = Black
	} else {
		g.CurrentTurnOwner = White
	}

	// Update figure
	f1.SetCoordinates(newX, newY)

	// Update board
	g.Board[oldX][oldY] = NotFigureId
	g.Board[newX][newY] = f1.Id

	return nil
}
