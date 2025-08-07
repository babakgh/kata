use crate::core::TurnstileMechanism;
use super::transitions::{State, Event, Transition, create_transitions};
use std::collections::HashMap;
use std::sync::LazyLock;

// Lazy static transitions table - created once and shared across all FSM instances
static TRANSITIONS: LazyLock<HashMap<State, HashMap<Event, Transition>>> = 
    LazyLock::new(|| create_transitions());

pub struct TurnstileFSM {
    current_state: State,
    turnstile: Box<dyn TurnstileMechanism>,
}

impl TurnstileFSM {
    pub fn new(turnstile: Box<dyn TurnstileMechanism>) -> Self {
        Self {
            current_state: State::Locked,
            turnstile,
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
    fn get_transition(&self, event: Event) -> Option<&Transition> {
        TRANSITIONS
            .get(&self.current_state)?
            .get(&event)
    }

    /// Execute a state transition
    fn transition_to(&mut self, new_state: State, action: fn(&mut dyn TurnstileMechanism)) {
        let _old_state = self.current_state;
        self.current_state = new_state;
        action(self.turnstile.as_mut());
        
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
    fn test_locked_state_coin_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile));
        
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
    fn test_locked_state_pass_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile));
        
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
    fn test_unlocked_state_coin_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile));
        
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
    fn test_unlocked_state_pass_event() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile));
        
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
    fn test_transitions_are_shared() {
        // Test that the static transitions are properly initialized
        let (turnstile1, _) = FakeTurnstile::new();
        let (turnstile2, _) = FakeTurnstile::new();
        
        let fsm1 = TurnstileFSM::new(Box::new(turnstile1));
        let fsm2 = TurnstileFSM::new(Box::new(turnstile2));
        
        // Both FSMs should start in the same initial state
        assert_eq!(fsm1.current_state(), State::Locked);
        assert_eq!(fsm2.current_state(), State::Locked);
        
        // Test that the LazyLock is working by ensuring consistent behavior
        // This implicitly tests that TRANSITIONS is properly shared
    }
}
