package mcts

type Player int

const (
	Player1 Player = 1
	Player2 Player = 2
)

func ChangePlayer(currentPlayer Player) Player {
	return 3 - currentPlayer
}

type Result float64

const (
	ResultNone Result = 2.0

	ResultWin  Result = 1.0
	ResultDraw Result = 0.5
	ResultLose Result = 0.0
)
