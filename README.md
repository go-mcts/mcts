# mcts

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/go-mcts/mcts/Go?logo=github)](https://github.com/go-mcts/mcts/actions?query=workflow%3AGo)
[![codecov](https://img.shields.io/codecov/c/github/go-mcts/mcts/main?logo=codecov)](https://codecov.io/gh/go-mcts/mcts)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-mcts/mcts)](https://goreportcard.com/report/github.com/go-mcts/mcts)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-mcts/mcts.svg)](https://pkg.go.dev/github.com/go-mcts/mcts)
[![Release](https://img.shields.io/github/release/go-mcts/mcts)](https://github.com/go-mcts/mcts/releases/latest)

Package mcts provides parallel monte-carlo tree search for your Go applications.

## Installation

Go latest is recommended.

```bash
go get -u github.com/go-mcts/mcts
```

## Getting started

See [examples](examples) directory.

Implements `mcts.Move` and `mcts.State`:

```go
// examples/nim/nim.go
type Move int

type State struct {
	playerToMove int
	chips        int
}

func (s *State) PlayerToMove() int {
	return s.playerToMove
}

func (s *State) HasMoves() bool {
	s.checkInvariant()
	return s.chips > 0
}

func (s *State) GetMoves() []mcts.Move {
	s.checkInvariant()

	var moves []mcts.Move
	for i := 1; i <= min(3, s.chips); i++ {
		moves = append(moves, i)
	}
	return moves
}

func (s *State) DoMove(move mcts.Move) {
	m := move.(int)
	if m < 1 || m > 3 {
		panic("illegal move")
	}
	s.checkInvariant()

	s.chips -= m
	s.playerToMove = 3 - s.playerToMove

	s.checkInvariant()
}

func (s *State) DoRandomMove(rd *rand.Rand) {
	if s.chips <= 0 {
		panic("invalid chips")
	}
	s.checkInvariant()

	max := min(3, s.chips)
	s.DoMove(rd.Intn(max) + 1)

	s.checkInvariant()
}

func (s *State) GetResult(currentPlayerToMove int) float64 {
	if s.chips != 0 {
		panic("game not over")
	}
	s.checkInvariant()

	if s.playerToMove == currentPlayerToMove {
		return 1.0
	}
	return 0.0
}

func (s *State) Clone() mcts.State {
	return &State{
		playerToMove: s.playerToMove,
		chips:        s.chips,
	}
}
```

Run `mcts.ComputeMove`:

```go
state := &State{
	playerToMove: 1,
	chips:        chips,
}
move := mcts.ComputeMove(state, mcts.MaxIterations(100000), mcts.Verbose(true))
```

## License

This project is under the MIT License. See the [LICENSE](LICENSE) file for the full license text.
