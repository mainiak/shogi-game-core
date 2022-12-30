package game

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

/*
 * Test suite
 */

var _ = Describe("Game v1", func() {
	var g *Game

	BeforeEach(func() {
		g = NewGame()
	})

	Describe("object", func() {
		It("NewGame()", func() {
			Expect(g).NotTo(BeNil())

			// Board size check
			Expect(len(g.Board)).Should(Equal(BoardSize))
			Expect(len(g.Board[0])).Should(Equal(BoardSize))

			// Initialized board
			Expect(g.Board[5][5]).Should(Equal(NotFigureId))

			// Game starts with 'White' player
			Expect(g.CurrentTurnOwner).Should(Equal(White))
		})
	})
})
