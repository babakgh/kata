mod core;

use core::Turnstile;
use core::table_impl::{TurnstileFSM, Event};

fn main() {
    println!("=== Turnstile FSM Demo ===");
    
    // Demonstrate the new generic approach (zero-cost abstraction)
    println!("\n--- Generic FSM (Static Dispatch) ---");
    let turnstile = Turnstile::new("Hardware Gate".to_string());
    let mut generic_fsm = TurnstileFSM::new(turnstile);
    
    println!("Initial state: {:?}", generic_fsm.current_state());
    
    println!("\n1. Inserting coin while locked:");
    generic_fsm.handle_event(Event::Coin);
    println!("State after coin: {:?}", generic_fsm.current_state());
    
    println!("\n2. Inserting another coin while unlocked:");
    generic_fsm.handle_event(Event::Coin);
    println!("State after second coin: {:?}", generic_fsm.current_state());
    
    println!("\n3. Passing through while unlocked:");
    generic_fsm.handle_event(Event::Pass);
    println!("State after pass: {:?}", generic_fsm.current_state());
    
    println!("\n4. Trying to pass while locked:");
    generic_fsm.handle_event(Event::Pass);
    println!("State after pass (locked): {:?}", generic_fsm.current_state());
    
    // Demonstrate trait objects with explicit generic syntax
    println!("\n--- Dynamic FSM (Trait Objects) ---");
    let dynamic_turnstile = Box::new(Turnstile::new("Software Gate".to_string())) as Box<dyn core::TurnstileMechanism>;
    let mut dynamic_fsm = TurnstileFSM::new(dynamic_turnstile);
    
    println!("Initial state: {:?}", dynamic_fsm.current_state());
    
    println!("\n1. Inserting coin while locked:");
    dynamic_fsm.handle_event(Event::Coin);
    println!("State after coin: {:?}", dynamic_fsm.current_state());
    
    println!("\n2. Passing through while unlocked:");
    dynamic_fsm.handle_event(Event::Pass);
    println!("State after pass: {:?}", dynamic_fsm.current_state());
    
    // Demonstrate collection of different turnstile types using explicit generic syntax
    println!("\n--- Mixed Collection (Dynamic Dispatch) ---");
    let turnstiles: Vec<TurnstileFSM<Box<dyn core::TurnstileMechanism>>> = vec![
        TurnstileFSM::new(Box::new(Turnstile::new("Gate 1".to_string())) as Box<dyn core::TurnstileMechanism>),
        TurnstileFSM::new(Box::new(Turnstile::new("Gate 2".to_string())) as Box<dyn core::TurnstileMechanism>),
        TurnstileFSM::new(Box::new(Turnstile::new("Gate 3".to_string())) as Box<dyn core::TurnstileMechanism>),
    ];
    
    println!("Processing {} turnstiles in collection:", turnstiles.len());
    for (i, mut fsm) in turnstiles.into_iter().enumerate() {
        println!("  Turnstile {}: {:?} -> ", i + 1, fsm.current_state());
        fsm.handle_event(Event::Coin);
        println!("{:?}", fsm.current_state());
    }
}
