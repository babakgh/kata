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

/// Action is a function that takes a mutable reference to a TurnstileMechanism
pub type Action = fn(&mut dyn TurnstileMechanism);

#[derive(Clone)]
pub struct Transition {
    pub current_state: State,
    pub event: Event,
    pub new_state: State,
    pub action: Action,
}

impl Transition {
    pub fn new(current_state: State, event: Event, new_state: State, action: Action) -> Self {
        Self {
            current_state,
            event,
            new_state,
            action,
        }
    }
}

/// Create the transition table
/// Given State - When Event - Then Next State - Then Action
/// Locked        Coin         Unlocked          unlock
/// Locked        Pass         Locked            alarm
/// Unlocked      Coin         Unlocked          thanks
/// Unlocked      Pass         Locked            lock
pub fn create_transitions() -> HashMap<State, HashMap<Event, Transition>> {
    let mut transitions = HashMap::new();
    
    // Locked state transitions
    let mut locked_transitions = HashMap::new();
    locked_transitions.insert(
        Event::Coin,
        Transition::new(State::Locked, Event::Coin, State::Unlocked, |t| t.unlock()),
    );
    locked_transitions.insert(
        Event::Pass,
        Transition::new(State::Locked, Event::Pass, State::Locked, |t| t.alarm()),
    );
    transitions.insert(State::Locked, locked_transitions);
    
    // Unlocked state transitions
    let mut unlocked_transitions = HashMap::new();
    unlocked_transitions.insert(
        Event::Coin,
        Transition::new(State::Unlocked, Event::Coin, State::Unlocked, |t| t.thanks()),
    );
    unlocked_transitions.insert(
        Event::Pass,
        Transition::new(State::Unlocked, Event::Pass, State::Locked, |t| t.lock()),
    );
    transitions.insert(State::Unlocked, unlocked_transitions);
    
    transitions
}
