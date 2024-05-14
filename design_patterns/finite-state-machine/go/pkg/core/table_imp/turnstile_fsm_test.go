package table_imp_test

import (
	"testing"

	"github.com/babakgh/kata/finite-state-machine/go/pkg/core/table_imp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTableImp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Table Implementation Suite")
}

type fakeTurnstile struct {
	lockedCalled, unlockedCalled, alarmedCalled, thanksCalled bool
}

func NewFakeTurnstile() *fakeTurnstile {
	f := new(fakeTurnstile)
	f.lockedCalled, f.unlockedCalled, f.alarmedCalled, f.thanksCalled = false, false, false, false
	return f
}

func (f *fakeTurnstile) Lock() {
	f.lockedCalled = true
}

func (f *fakeTurnstile) Unlock() {
	f.unlockedCalled = true
}

func (f *fakeTurnstile) Alarm() {
	f.alarmedCalled = true
}

func (f *fakeTurnstile) Thanks() {
	f.thanksCalled = true
}

var _ = Describe("FSM", func() {
	var (
		turnstile *fakeTurnstile
		fsm       *table_imp.TurnstileFSM
	)

	BeforeEach(func() {
		turnstile = NewFakeTurnstile()
		fsm = table_imp.NewTurnstileFSM(turnstile)
	})

	Describe("handling coins and passes", func() {
		Context("when the state is Locked", func() {
			It("should unlock when a coin is inserted", func() {
				fsm.HandleEvent(table_imp.Coin)
				Expect(fsm.CurrentState).To(Equal(table_imp.Unlocked))
				Expect(turnstile.unlockedCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.lockedCalled).To(Equal(false))
				Expect(turnstile.alarmedCalled).To(Equal(false))
				Expect(turnstile.thanksCalled).To(Equal(false))
			})

			It("should remain locked and raise an alarm when a pass is used", func() {
				fsm.HandleEvent(table_imp.Pass)
				Expect(fsm.CurrentState).To(Equal(table_imp.Locked))
				Expect(turnstile.alarmedCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.lockedCalled).To(Equal(false))
				Expect(turnstile.unlockedCalled).To(Equal(false))
				Expect(turnstile.thanksCalled).To(Equal(false))
			})
		})

		Context("when the state is Unlocked", func() {
			BeforeEach(func() {
				fsm.HandleEvent(table_imp.Coin) // Unlock the FSM
			})

			It("should remain unlocked and thank when a coin is inserted", func() {
				Expect(turnstile.unlockedCalled).To(Equal(true))

				fsm.HandleEvent(table_imp.Coin)
				Expect(fsm.CurrentState).To(Equal(table_imp.Unlocked))
				Expect(turnstile.thanksCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.lockedCalled).To(Equal(false))
				Expect(turnstile.alarmedCalled).To(Equal(false))
			})

			It("should lock when a pass is used", func() {
				Expect(turnstile.unlockedCalled).To(Equal(true))

				fsm.HandleEvent(table_imp.Pass)
				Expect(fsm.CurrentState).To(Equal(table_imp.Locked))
				Expect(turnstile.lockedCalled).To(Equal(true))

				// Ensure other actions haven't been called
				Expect(turnstile.alarmedCalled).To(Equal(false))
				Expect(turnstile.thanksCalled).To(Equal(false))
			})
		})
	})
})
