package core

import "fmt"

type TurnstileMechanism interface {
	// Actions
	Lock()
	Unlock()
	Alarm()
	Thanks()
}

type Turnstile struct {
	name string
}

func NewTurnstile(name string) *Turnstile {
	t := new(Turnstile)
	t.name = name
	return t
}

func (t *Turnstile) Lock() {
	fmt.Printf("%s: Lock\n", t.name)
}

func (t *Turnstile) Unlock() {
	fmt.Printf("%s: Unlock\n", t.name)
}

func (t *Turnstile) Alarm() {
	fmt.Printf("%s: Alarm\n", t.name)
}

func (t *Turnstile) Thanks() {
	fmt.Printf("%s: Thanks\n", t.name)
}
