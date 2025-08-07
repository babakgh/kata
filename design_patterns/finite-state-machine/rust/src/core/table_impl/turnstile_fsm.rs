use crate::core::TurnstileMechanism;
use super::transitions::{State, Event, Transition, create_transitions};
use std::collections::HashMap;

/// Generic finite state machine for turnstiles
pub struct TurnstileFSM<T: TurnstileMechanism> {
    current_state: State,
    turnstile: T,
    transitions: HashMap<State, HashMap<Event, Transition<T>>>,
}

impl<T: TurnstileMechanism> TurnstileFSM<T> {
    pub fn new(turnstile: T) -> Self {
        Self {
            current_state: State::Locked,
            turnstile,
            transitions: create_transitions(),
        }
    }
    
    pub fn current_state(&self) -> State {
        self.current_state
    }
    
    pub fn handle_event(&mut self, event: Event) {
        // First check if transition exists and get the new state and action
        if let Some(transition) = self.get_transition(event) {
            let new_state = transition.new_state;
            let action = transition.action;
            
            // Now perform the transition with the captured values
            self.transition_to(new_state, action);
        } else {
            self.handle_invalid_transition(event);
        }
    }

    /// Get the transition for the current state and event
    fn get_transition(&self, event: Event) -> Option<&Transition<T>> {
        self.transitions
            .get(&self.current_state)?
            .get(&event)
    }

    /// Execute a state transition
    fn transition_to(&mut self, new_state: State, action: fn(&mut T)) {
        let _old_state = self.current_state;
        self.current_state = new_state;
        action(&mut self.turnstile);
        
        // Could add logging here for debugging:
        // log::debug!("FSM transition: {:?} -> {:?}", old_state, new_state);
    }

    /// Handle invalid state/event combinations
    fn handle_invalid_transition(&self, _event: Event) {
        // don't panic but notify!
        // TODO: Write a test for me
        print!("I am the one who reduced the coverage ;)!");
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::core::TurnstileMechanism;
    use std::cell::RefCell;
    use std::rc::Rc;

    #[derive(Debug, Default)]
    struct TestCallsRecord {
        locked_called: bool,
        unlocked_called: bool,
        alarmed_called: bool,
        thanks_called: bool,
    }

    impl TestCallsRecord {
        fn reset(&mut self) {
            self.locked_called = false;
            self.unlocked_called = false;
            self.alarmed_called = false;
            self.thanks_called = false;
        }
    }

    struct FakeTurnstile {
        calls: Rc<RefCell<TestCallsRecord>>,
    }

    impl FakeTurnstile {
        fn new() -> (Self, Rc<RefCell<TestCallsRecord>>) {
            let calls = Rc::new(RefCell::new(TestCallsRecord::default()));
            (
                Self {
                    calls: calls.clone(),
                },
                calls,
            )
        }
    }

    impl TurnstileMechanism for FakeTurnstile {
        fn lock(&mut self) {
            self.calls.borrow_mut().locked_called = true;
        }

        fn unlock(&mut self) {
            self.calls.borrow_mut().unlocked_called = true;
        }

        fn alarm(&mut self) {
            self.calls.borrow_mut().alarmed_called = true;
        }

        fn thanks(&mut self) {
            self.calls.borrow_mut().thanks_called = true;
        }
    }

    #[test]
    fn test_generic_fsm_locked_state_coin_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(turnstile);
        
        assert_eq!(fsm.current_state(), State::Locked);
        
        fsm.handle_event(Event::Coin);
        
        assert_eq!(fsm.current_state(), State::Unlocked);
        
        let calls_ref = calls.borrow();
        assert!(calls_ref.unlocked_called);
        assert!(!calls_ref.locked_called);
        assert!(!calls_ref.alarmed_called);
        assert!(!calls_ref.thanks_called);
    }

    #[test]
    fn test_generic_fsm_locked_state_pass_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(turnstile);
        
        assert_eq!(fsm.current_state(), State::Locked);
        
        fsm.handle_event(Event::Pass);
        
        assert_eq!(fsm.current_state(), State::Locked);
        
        let calls_ref = calls.borrow();
        assert!(calls_ref.alarmed_called);
        assert!(!calls_ref.locked_called);
        assert!(!calls_ref.unlocked_called);
        assert!(!calls_ref.thanks_called);
    }

    #[test]
    fn test_generic_fsm_unlocked_state_coin_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(turnstile);
        
        // First unlock the FSM
        fsm.handle_event(Event::Coin);
        assert_eq!(fsm.current_state(), State::Unlocked);
        
        // Reset the call tracking
        calls.borrow_mut().reset();
        
        fsm.handle_event(Event::Coin);
        
        assert_eq!(fsm.current_state(), State::Unlocked);
        
        let calls_ref = calls.borrow();
        assert!(calls_ref.thanks_called);
        assert!(!calls_ref.locked_called);
        assert!(!calls_ref.unlocked_called);
        assert!(!calls_ref.alarmed_called);
    }

    #[test]
    fn test_generic_fsm_unlocked_state_pass_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(turnstile);
        
        // First unlock the FSM
        fsm.handle_event(Event::Coin);
        assert_eq!(fsm.current_state(), State::Unlocked);
        
        // Reset the call tracking
        calls.borrow_mut().reset();
        
        fsm.handle_event(Event::Pass);
        
        assert_eq!(fsm.current_state(), State::Locked);
        
        let calls_ref = calls.borrow();
        assert!(calls_ref.locked_called);
        assert!(!calls_ref.unlocked_called);
        assert!(!calls_ref.alarmed_called);
        assert!(!calls_ref.thanks_called);
    }

    #[test]
    fn test_zero_cost_abstraction() {
        // Test that generics work with concrete types (zero-cost abstraction)
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm: TurnstileFSM<FakeTurnstile> = TurnstileFSM::new(turnstile);
        
        // This should compile to direct function calls, no vtable lookup
        fsm.handle_event(Event::Coin);
        fsm.handle_event(Event::Pass);
        
        let calls_ref = calls.borrow();
        assert!(calls_ref.unlocked_called);
        assert!(calls_ref.locked_called);
    }

    #[test]
    fn test_backward_compatibility_with_trait_objects() {
        // Test that we can still use trait objects for dynamic dispatch
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile) as Box<dyn TurnstileMechanism>);
        
        assert_eq!(fsm.current_state(), State::Locked);
        
        fsm.handle_event(Event::Coin);
        
        assert_eq!(fsm.current_state(), State::Unlocked);
        
        let calls_ref = calls.borrow();
        assert!(calls_ref.unlocked_called);
    }

    #[test]
    fn test_multiple_concrete_types() {
        // Test that different concrete types generate different specialized implementations
        
        struct SimpleTurnstile {
            state: String,
        }
        
        impl SimpleTurnstile {
            fn new() -> Self {
                Self { state: "initial".to_string() }
            }
        }
        
        impl TurnstileMechanism for SimpleTurnstile {
            fn lock(&mut self) { self.state = "locked".to_string(); }
            fn unlock(&mut self) { self.state = "unlocked".to_string(); }
            fn alarm(&mut self) { self.state = "alarm".to_string(); }
            fn thanks(&mut self) { self.state = "thanks".to_string(); }
        }
        
        // Each of these will get its own monomorphized implementation
        let (fake_turnstile, _) = FakeTurnstile::new();
        let simple_turnstile = SimpleTurnstile::new();
        
        let mut fake_fsm = TurnstileFSM::new(fake_turnstile);
        let mut simple_fsm = TurnstileFSM::new(simple_turnstile);
        
        fake_fsm.handle_event(Event::Coin);
        simple_fsm.handle_event(Event::Coin);
        
        assert_eq!(fake_fsm.current_state(), State::Unlocked);
        assert_eq!(simple_fsm.current_state(), State::Unlocked);
    }
}
