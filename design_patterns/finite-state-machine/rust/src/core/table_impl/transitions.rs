use crate::core::TurnstileMechanism;
use std::collections::HashMap;

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum State {
    Locked,
    Unlocked,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
pub enum Event {
    Coin,
    Pass,
}

/// Generic action type that works with any TurnstileMechanism implementation
pub type Action<T> = fn(&mut T);

/// Transition struct for generic turnstile mechanisms
#[derive(Clone)]
pub struct Transition<T: TurnstileMechanism> {
    pub new_state: State,
    pub action: Action<T>,
}

impl<T: TurnstileMechanism> Transition<T> {
    pub fn new(new_state: State, action: Action<T>) -> Self {
        Self { new_state, action }
    }
}

/// Create the transition table for a specific turnstile type
/// Given State - When Event - Then Next State - Then Action
/// Locked        Coin         Unlocked          unlock
/// Locked        Pass         Locked            alarm
/// Unlocked      Coin         Unlocked          thanks
/// Unlocked      Pass         Locked            lock
pub fn create_transitions<T: TurnstileMechanism>() -> HashMap<State, HashMap<Event, Transition<T>>> {
    let mut transitions = HashMap::new();

    // Locked state transitions
    let mut locked_transitions = HashMap::new();
    locked_transitions.insert(
        Event::Coin,
        Transition::new(State::Unlocked, |t: &mut T| t.unlock()),
    );
    locked_transitions.insert(
        Event::Pass, 
        Transition::new(State::Locked, |t: &mut T| t.alarm())
    );
    transitions.insert(State::Locked, locked_transitions);

    // Unlocked state transitions
    let mut unlocked_transitions = HashMap::new();
    unlocked_transitions.insert(
        Event::Coin,
        Transition::new(State::Unlocked, |t: &mut T| t.thanks()),
    );
    unlocked_transitions.insert(
        Event::Pass, 
        Transition::new(State::Locked, |t: &mut T| t.lock())
    );
    transitions.insert(State::Unlocked, unlocked_transitions);

    transitions
}
