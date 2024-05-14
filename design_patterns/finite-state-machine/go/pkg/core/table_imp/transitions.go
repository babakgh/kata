package table_imp

import "github.com/babakgh/kata/finite-state-machine/go/pkg/core"

type Action func(t core.TurnstileMechanism)

type Transition struct {
	currentState State
	event        Event
	newState     State
	action       Action
}

// Given State - When Event - Then Next State - Then Action
// Locked        Coin         Unlocked          unlock
// Locked        Pass         Locked            alarm
// Unlocked      Coin         Unlocked          thanks
// Unlocked      Pass         Locked            lock
var Transitions map[State]map[Event]Transition = map[State]map[Event]Transition{
	Locked: {
		Coin: Transition{Locked, Coin, Unlocked, func(t core.TurnstileMechanism) { t.Unlock() }},
		Pass: Transition{Locked, Pass, Locked, func(t core.TurnstileMechanism) { t.Alarm() }},
	},
	Unlocked: {
		Coin: Transition{Unlocked, Coin, Unlocked, func(t core.TurnstileMechanism) { t.Thanks() }},
		Pass: Transition{Unlocked, Pass, Locked, func(t core.TurnstileMechanism) { t.Lock() }},
	},
}
