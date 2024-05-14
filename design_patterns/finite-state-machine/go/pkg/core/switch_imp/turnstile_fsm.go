package switch_imp

import "github.com/babakgh/kata/finite-state-machine/go/pkg/core"

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
	switch fsm.CurrentState {
	case Locked:
		switch event {
		case Coin:
			fsm.CurrentState = Unlocked
			fsm.turnstile.Unlock()
		case Pass:
			fsm.turnstile.Alarm()
		}
	case Unlocked:
		switch event {
		case Coin:
			fsm.turnstile.Thanks()
		case Pass:
			fsm.CurrentState = Locked
			fsm.turnstile.Lock()
		}
	}
}
