package mcts

import (
	"context"
	"log"
	"math/rand"
	"sort"
)

func MultiGoroutineUCT(ctx context.Context, rootState State, numberOfGoroutines int) Move {
	nodeChan := make(chan *Node, numberOfGoroutines)

	for i := 0; i < numberOfGoroutines; i++ {
		go func() {
			nodeChan <- internalUCTWithContext(ctx, rootState)
		}()
	}

	visits := make(map[Move]int)
	wins := make(map[Move]float64)
	for i := 0; i < numberOfGoroutines; i++ {
		node := <-nodeChan
		for _, child := range node.ChildNodes {
			visits[child.Move] += child.Visits
			wins[child.Move] += child.Wins
		}
	}

	bestScore := minFloat64
	var bestMove Move

	for m, v := range visits {
		w := wins[m]
		score := (w + 1) / float64(v+2)
		if score > bestScore {
			bestScore = score
			bestMove = m
		}
	}

	log.Printf("player %d score: %.2f", 3-rootState.PlayerJustMoved(), bestScore)

	return bestMove
}

func internalUCTWithContext(ctx context.Context, rootState State) *Node {
	rootNode := NewNode(nil, nil, rootState)

LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			node := rootNode
			state := rootState.Clone()

			for len(node.UntriedMoves) == 0 && len(node.ChildNodes) != 0 {
				node = node.UCTSelectChild()
				state.DoMove(node.Move)
			}

			if len(node.UntriedMoves) != 0 {
				m := randomChoice(node.UntriedMoves)
				state.DoMove(m)
				node = node.AddChild(m, state)
			}

			for len(state.GetMoves()) != 0 {
				state.DoMove(randomChoice(state.GetMoves()))
			}

			for node != nil {
				node.Update(state.GetResult(node.PlayerJustMoved))
				node = node.Parent
			}
		}
	}

	return rootNode
}

func UCT(rootState State, maxIterations int) Move {
	rootNode := NewNode(nil, nil, rootState)

	for i := 0; i < maxIterations; i++ {
		node := rootNode
		state := rootState.Clone()

		for len(node.UntriedMoves) == 0 && len(node.ChildNodes) != 0 {
			node = node.UCTSelectChild()
			state.DoMove(node.Move)
		}

		if len(node.UntriedMoves) != 0 {
			m := randomChoice(node.UntriedMoves)
			state.DoMove(m)
			node = node.AddChild(m, state)
		}

		for len(state.GetMoves()) != 0 {
			state.DoMove(randomChoice(state.GetMoves()))
		}

		for node != nil {
			node.Update(state.GetResult(node.PlayerJustMoved))
			node = node.Parent
		}
	}

	sort.Slice(rootNode.ChildNodes, func(i, j int) bool {
		return rootNode.ChildNodes[i].Visits > rootNode.ChildNodes[j].Visits
	})
	return rootNode.ChildNodes[len(rootNode.ChildNodes)-1].Move
}

func randomChoice(moves []Move) Move {
	return moves[rand.Intn(len(moves))]
}
