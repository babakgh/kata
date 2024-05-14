package oo_imp

import "github.com/babakgh/kata/finite-state-machine/go/pkg/core"

type TurnstileFSM struct {
	CurrentState State
	Turnstile    core.TurnstileMechanism
}

func NewTurnstileFSM(turnstile core.TurnstileMechanism) *TurnstileFSM {
	fsm := new(TurnstileFSM)
	fsm.CurrentState = Locked{}
	fsm.Turnstile = turnstile
	return fsm
}
