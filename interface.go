package mcts

type Move interface{}

type State interface {
	PlayerJustMoved() int

	Clone() State

	DoMove(Move)

	GetMoves() []Move

	// GetResult returns a value int {0, 0.5, 1}
	GetResult(playerJustMoved int) float64
}
