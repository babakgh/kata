mod core;

use core::Turnstile;
use core::table_impl::{TurnstileFSM, Event};

fn main() {
    let turnstile = Box::new(Turnstile::new("Main Gate".to_string()));
    let mut fsm = TurnstileFSM::new(turnstile);
    
    println!("=== Turnstile FSM Demo ===");
    println!("Initial state: {:?}", fsm.current_state());
    
    println!("\n1. Inserting coin while locked:");
    fsm.handle_event(Event::Coin);
    println!("State after coin: {:?}", fsm.current_state());
    
    println!("\n2. Inserting another coin while unlocked:");
    fsm.handle_event(Event::Coin);
    println!("State after second coin: {:?}", fsm.current_state());
    
    println!("\n3. Passing through while unlocked:");
    fsm.handle_event(Event::Pass);
    println!("State after pass: {:?}", fsm.current_state());
    
    println!("\n4. Trying to pass while locked:");
    fsm.handle_event(Event::Pass);
    println!("State after pass (locked): {:?}", fsm.current_state());
}
