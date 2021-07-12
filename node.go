package mcts

import "math"

var SliceInitialCapacity = 8

const minFloat64 float64 = -1061109568

type Node struct {
	Move            Move
	Parent          *Node
	ChildNodes      []*Node
	Wins            float64
	Visits          int
	UntriedMoves    []Move
	PlayerJustMoved int
}

func NewNode(move Move, parent *Node, state State) *Node {
	return &Node{
		Move:            move,
		Parent:          parent,
		ChildNodes:      make([]*Node, 0, SliceInitialCapacity),
		Wins:            0,
		Visits:          0,
		UntriedMoves:    state.GetMoves(),
		PlayerJustMoved: state.PlayerJustMoved(),
	}
}

func (n *Node) UCTSelectChild() *Node {
	max, idx := minFloat64, -1
	for i, c := range n.ChildNodes {
		s := c.Wins/float64(c.Visits) + math.Sqrt(2*math.Log(float64(n.Visits))/float64(c.Visits))
		if s > max {
			max = s
			idx = i
		}
	}
	return n.ChildNodes[idx]
}

func (n *Node) AddChild(move Move, state State) *Node {
	node := NewNode(move, n, state)
	l := len(n.UntriedMoves)
	for i := 0; i < l-1; i++ {
		m := n.UntriedMoves[i]
		if m == move {
			n.UntriedMoves[i] = n.UntriedMoves[l-1]
			break
		}
	}
	n.UntriedMoves = n.UntriedMoves[:l-1]
	n.ChildNodes = append(n.ChildNodes, node)
	return node
}

func (n *Node) Update(result float64) {
	n.Visits++
	n.Wins += result
}
