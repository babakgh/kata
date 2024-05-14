package table_imp

import (
	"fmt"

	"github.com/babakgh/kata/finite-state-machine/go/pkg/core"
)

type State int

const (
	Locked State = iota + 1
	Unlocked
)

type Event int

const (
	Coin Event = iota + 1
	Pass
)

type TurnstileFSM struct {
	CurrentState State
	turnstile    core.TurnstileMechanism
}

func NewTurnstileFSM(turnstile core.TurnstileMechanism) *TurnstileFSM {
	fsm := new(TurnstileFSM)
	fsm.CurrentState = Locked
	fsm.turnstile = turnstile
	return fsm
}

func (fsm *TurnstileFSM) HandleEvent(event Event) {
	if tr, ok := Transitions[fsm.CurrentState][event]; ok {
		fsm.CurrentState = tr.newState
		tr.action(fsm.turnstile)
	} else {
		// don't panic but notify!
		// TODO: Write a test for me
		fmt.Print("I am the one who reduced the coverage ;)!")
	}
}
