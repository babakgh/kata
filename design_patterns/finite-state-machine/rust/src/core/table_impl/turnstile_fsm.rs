use crate::core::TurnstileMechanism;
use super::transitions::{State, Event, Transition, create_transitions};
use std::collections::HashMap;

pub struct TurnstileFSM {
    current_state: State,
    turnstile: Box<dyn TurnstileMechanism>,
    transitions: HashMap<State, HashMap<Event, Transition>>,
}

impl TurnstileFSM {
    pub fn new(turnstile: Box<dyn TurnstileMechanism>) -> Self {
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
        if let Some(state_transitions) = self.transitions.get(&self.current_state) {
            if let Some(transition) = state_transitions.get(&event) {
                self.current_state = transition.new_state;
                (transition.action)(self.turnstile.as_mut());
            } else {
                // don't panic but notify!
                // TODO: Write a test for me
                print!("I am the one who reduced the coverage ;)!");
            }
        }
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
}
