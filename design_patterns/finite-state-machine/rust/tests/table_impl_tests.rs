use finite_state_machine::core::table_impl::{Event, State, TurnstileFSM};
use finite_state_machine::core::TurnstileMechanism;
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

#[cfg(test)]
mod turnstile_fsm_integration_tests {
    use super::*;

    #[test]
    fn test_locked_state_with_coin_should_unlock() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile));

        assert_eq!(fsm.current_state(), State::Locked);

        fsm.handle_event(Event::Coin);

        assert_eq!(fsm.current_state(), State::Unlocked);

        let calls_ref = calls.borrow();
        assert!(calls_ref.unlocked_called, "unlock should be called");
        assert!(!calls_ref.locked_called, "lock should not be called");
        assert!(!calls_ref.alarmed_called, "alarm should not be called");
        assert!(!calls_ref.thanks_called, "thanks should not be called");
    }

    #[test]
    fn test_locked_state_with_pass_should_remain_locked_and_alarm() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile));

        assert_eq!(fsm.current_state(), State::Locked);

        fsm.handle_event(Event::Pass);

        assert_eq!(fsm.current_state(), State::Locked);

        let calls_ref = calls.borrow();
        assert!(calls_ref.alarmed_called, "alarm should be called");
        assert!(!calls_ref.locked_called, "lock should not be called");
        assert!(!calls_ref.unlocked_called, "unlock should not be called");
        assert!(!calls_ref.thanks_called, "thanks should not be called");
    }

    #[test]
    fn test_unlocked_state_with_coin_should_remain_unlocked_and_thank() {
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
        assert!(calls_ref.thanks_called, "thanks should be called");
        assert!(!calls_ref.locked_called, "lock should not be called");
        assert!(!calls_ref.unlocked_called, "unlock should not be called");
        assert!(!calls_ref.alarmed_called, "alarm should not be called");
    }

    #[test]
    fn test_unlocked_state_with_pass_should_lock() {
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
        assert!(calls_ref.locked_called, "lock should be called");
        assert!(!calls_ref.unlocked_called, "unlock should not be called");
        assert!(!calls_ref.alarmed_called, "alarm should not be called");
        assert!(!calls_ref.thanks_called, "thanks should not be called");
    }

    #[test]
    fn test_complete_workflow() {
        let (turnstile, calls) = FakeTurnstile::new();
        let mut fsm = TurnstileFSM::new(Box::new(turnstile));

        // Initial state should be locked
        assert_eq!(fsm.current_state(), State::Locked);

        // 1. Insert coin while locked -> should unlock
        fsm.handle_event(Event::Coin);
        assert_eq!(fsm.current_state(), State::Unlocked);
        assert!(calls.borrow().unlocked_called);

        // 2. Insert another coin while unlocked -> should thank
        calls.borrow_mut().reset();
        fsm.handle_event(Event::Coin);
        assert_eq!(fsm.current_state(), State::Unlocked);
        assert!(calls.borrow().thanks_called);

        // 3. Pass through while unlocked -> should lock
        calls.borrow_mut().reset();
        fsm.handle_event(Event::Pass);
        assert_eq!(fsm.current_state(), State::Locked);
        assert!(calls.borrow().locked_called);

        // 4. Try to pass while locked -> should alarm
        calls.borrow_mut().reset();
        fsm.handle_event(Event::Pass);
        assert_eq!(fsm.current_state(), State::Locked);
        assert!(calls.borrow().alarmed_called);
    }
}
