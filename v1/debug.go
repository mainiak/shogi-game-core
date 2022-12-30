package game

func (g *Game) SetTestBoard(ft FigureType) FigureId {
	var f *Figure = nil

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

	return f.Id
}

func (g *Game) TestMove(targetX, targetY Axis, fId FigureId, expectedValid bool) {
	// reset test board
	var origX, origY Axis = 4, 4
	g.PlaceFigure(origX, origY, fId)

	// It is always test figure move turn
	g.CurrentTurnOwner = White

	// move test
	err := g.Move(origX, origY, targetX, targetY)

	// board cleanup - HACK FIXME
	if err == nil {
		g.Board[targetX][targetY] = NotFigureId
	}
}
