package oo_imp

type State interface {
	Coin(fsm *TurnstileFSM)
	Pass(fsm *TurnstileFSM)
}

type Locked struct{}

func (l Locked) Coin(fsm *TurnstileFSM) {
	fsm.CurrentState = Unlocked{}
	fsm.Turnstile.Unlock()
}

func (l Locked) Pass(fsm *TurnstileFSM) {
	fsm.Turnstile.Alarm()
}

type Unlocked struct{}

func (l Unlocked) Coin(fsm *TurnstileFSM) {
	fsm.Turnstile.Thanks()
}

func (l Unlocked) Pass(fsm *TurnstileFSM) {
	fsm.CurrentState = Locked{}
	fsm.Turnstile.Lock()
}
