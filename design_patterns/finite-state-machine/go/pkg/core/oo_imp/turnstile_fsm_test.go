package oo_imp_test

import (
	"testing"

	"github.com/babakgh/kata/finite-state-machine/go/pkg/core/oo_imp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestOOImp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Object Oriented Implementation Suite")
}

type turnstileSpy struct {
	lockedCalled, unlockedCalled, alarmedCalled, thanksCalled bool
}

func NewTurnstileSpy() *turnstileSpy {
	f := new(turnstileSpy)
	f.lockedCalled, f.unlockedCalled, f.alarmedCalled, f.thanksCalled = false, false, false, false
	return f
}

func (f *turnstileSpy) Lock() {
	f.lockedCalled = true
}

func (f *turnstileSpy) Unlock() {
	f.unlockedCalled = true
}

func (f *turnstileSpy) Alarm() {
	f.alarmedCalled = true
}

func (f *turnstileSpy) Thanks() {
	f.thanksCalled = true
}

var _ = Describe("FSM", func() {
	var (
		turnstile *turnstileSpy
		fsm       *oo_imp.TurnstileFSM
	)

	BeforeEach(func() {
		turnstile = NewTurnstileSpy()
		fsm = oo_imp.NewTurnstileFSM(turnstile)
	})

	Describe("handling coins and passes", func() {
		Context("when the state is Locked", func() {
			It("should unlock when a coin is inserted", func() {
				fsm.CurrentState.Coin(fsm)
				Expect(fsm.CurrentState).To(Equal(oo_imp.Unlocked{}))
				Expect(turnstile.unlockedCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.lockedCalled).To(Equal(false))
				Expect(turnstile.alarmedCalled).To(Equal(false))
				Expect(turnstile.thanksCalled).To(Equal(false))
			})

			It("should remain locked and raise an alarm when a pass is used", func() {
				fsm.CurrentState.Pass(fsm)
				Expect(fsm.CurrentState).To(Equal(oo_imp.Locked{}))
				Expect(turnstile.alarmedCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.lockedCalled).To(Equal(false))
				Expect(turnstile.unlockedCalled).To(Equal(false))
				Expect(turnstile.thanksCalled).To(Equal(false))
			})
		})

		Context("when the state is Unlocked", func() {
			BeforeEach(func() {
				fsm.CurrentState.Coin(fsm) // Unlock the FSM
			})

			It("should remain unlocked and thank when a coin is inserted", func() {
				Expect(turnstile.unlockedCalled).To(Equal(true))

				fsm.CurrentState.Coin(fsm)
				Expect(fsm.CurrentState).To(Equal(oo_imp.Unlocked{}))
				Expect(turnstile.thanksCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.lockedCalled).To(Equal(false))
				Expect(turnstile.alarmedCalled).To(Equal(false))
			})

			It("should lock when a pass is used", func() {
				Expect(turnstile.unlockedCalled).To(Equal(true))

				fsm.CurrentState.Pass(fsm)
				Expect(fsm.CurrentState).To(Equal(oo_imp.Locked{}))
				Expect(turnstile.lockedCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.alarmedCalled).To(Equal(false))
				Expect(turnstile.thanksCalled).To(Equal(false))
			})
		})
	})
})
