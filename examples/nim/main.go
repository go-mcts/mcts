package main

import (
	"fmt"
	"github.com/go-mcts/mcts"
	"math/rand"
	"time"
)

func UCTPlayGame() {
	state := &NimState{
		playerJustMoved: mcts.Player2,
		chips:           15,
	}
	for len(state.GetMoves()) != 0 {
		m := mcts.MultiGoroutineUCT(state, 5*time.Second, 5)
		fmt.Printf("chips: %d, player: %d, count: %d\n", state.chips, 3-state.playerJustMoved, m)
		state.DoMove(m)
	}
	if state.GetResult(state.playerJustMoved) == 1.0 {
		fmt.Printf("Player %d wins!\n", state.playerJustMoved)
	} else if state.GetResult(state.playerJustMoved) == 0.0 {
		fmt.Printf("Player %d wins!\n", 3-state.playerJustMoved)
	} else {
		fmt.Println("Nobody wins!")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	UCTPlayGame()
}
