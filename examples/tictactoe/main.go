package main

import (
	"fmt"
	"github.com/go-mcts/mcts"
	"runtime"
	"time"
)

func main() {
	state := &State{
		playerJustMoved: mcts.Player2, // next is player1 (ChessX)
		board: [3][3]int{
			{0, 0, 0},
			{0, ChessX, ChessO},
			{0, 0, 0},
		},
	}

	move := mcts.MultiGoroutineUCT(state, time.Second, runtime.NumCPU()+1)

	fmt.Printf("%#v\n", move)
}
