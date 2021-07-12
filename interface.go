package mcts

type Move interface{}

type State interface {
	// PlayerJustMoved current step
	PlayerJustMoved() Player

	Clone() State

	DoMove(Move)

	GetMoves() []Move

	// GetResult returns a value in {0, 0.5, 1}
	// if game is not over, you should panic an error
	GetResult(playerJustMoved Player) Result
}
