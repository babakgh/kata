package root

// * Dynamic - Definable and Loadable
// * Multi Mechanism
// * Tidy usage
// * Memory efficient
// * Distributable
// * Persistable
// * Generic?!

type Mechanism any

type State interface {
	// Has a name
	// Has a entry action
	// Has a exit action
}

type FiniteStateMachineInterface interface {
	// Should have Current State
}

type FiniteStateMachineFactory interface {
	CreateFiniteStateMachine() FiniteStateMachineInterface
}

type FiniteStateMachineContext struct {
	CurrentState State
}
