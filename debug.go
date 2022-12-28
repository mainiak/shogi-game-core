package game

import (
	"fmt"
)

func (g *Game) SetTestBoard(ft FigureType) FigureId {
	var f *Figure = nil

	//fmt.Printf("\n## Test Board: %s\n", ft) // DEBUG
	g.CleanBoard()

	// Grab next available figure from pile
	for i := range g.Figures {
		f = g.Figures[i]
		if f.Type == ft {
			break
		}
	}

	// Put figure in center of the board
	g.PlaceFigure(4, 4, f.Id)

/*
	fmt.Printf("# %s\n", f) // DEBUG
	g.Draw()                // DEBUG
*/

	return f.Id
}

func (g *Game) TestMove(targetX, targetY Axis, fId FigureId, expectedValid bool) {
	// reset test board
	var origX, origY Axis = 4, 4
	g.PlaceFigure(origX, origY, fId)

	// It is always test figure move turn
	g.CurrentTurnOwner = White

	fmt.Printf("# Expecting outcome: ")
	if expectedValid {
		fmt.Printf("VALID\n")
	} else {
		fmt.Printf("INVALID\n")
	}

	// move test
	err := g.Move(origX, origY, targetX, targetY)
	if err != nil {
		fmt.Printf("!! %s\n", err)
		if expectedValid {
			fmt.Printf("# Result: INVALID - WRONG\n")
		} else {
			fmt.Printf("# Result: INVALID - correct\n")
		}
	} else {
		if expectedValid {
			fmt.Printf("# Result: VALID - correct\n")
		} else {
			fmt.Printf("# Result: VALID - WRONG\n")
		}
	}
	g.Draw() // DEBUG

	// board cleanup - HACK FIXME
	if err == nil {
		g.Board[targetX][targetY] = NotFigureId
	}
}

func (g *Game) Status() {
	var err error
	var x, y Axis
	var f *Figure
	for y = 9; y > 0; y-- {
		for x = 0; x < BoardSize; x++ {
			f, err = g.GetFigure(x, (y - 1))
			if err != nil {
				panic(err)
			}
			fmt.Printf("(%d,%d) %s\n", x, (y - 1), f)
		}
	}
}

func (g *Game) FigureStatus() {
	for i := range g.Figures {
		f := g.Figures[i]
		fmt.Printf("%s\n", f)
	}
}

func (g *Game) Draw() {
	var err error
	var x, y Axis
	var f *Figure
	var s string

	// Ascii CLI ART !!
	for y = 9; y > 0; y-- {
		for x = 0; x < BoardSize; x++ {
			f, err = g.GetFigure(x, (y - 1))
			if err != nil {
				panic(err)
			}

			if f == nil {
				s = "       "
			} else {
				s = fmt.Sprintf("(%s) %s", f.Id, f.Short())
			}

			fmt.Printf("| %s ", s)

			if x == 8 {
				fmt.Printf("|\n")
			}
		}
	}
}

func (g *Game) DrawKanji() {
	var err error
	var x, y Axis
	var f *Figure
	var s string

	// Ascii CLI ART !!
	for y = 9; y > 0; y-- {
		for x = 0; x < BoardSize; x++ {
			f, err = g.GetFigure(x, (y - 1))
			if err != nil {
				panic(err)
			}

			if f == nil {
				s = "    "
			} else {
				s = f.Type.NameKanji()
			}

			fmt.Printf("| %s ", s)

			if x == 8 {
				fmt.Printf("|\n")
			}
		}
	}
}
