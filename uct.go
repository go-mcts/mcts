package mcts

import (
	"math/rand"
	"sort"
)

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
