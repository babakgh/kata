# Finite State Machine - Rust Implementation

This is a Rust implementation of the Finite State Machine design pattern, specifically implementing a turnstile FSM using the table-driven approach.

## Structure

The project follows the same architectural patterns as the Go implementation:

- `src/core/mod.rs` - Core trait definition for `TurnstileMechanism`
- `src/core/table_impl/` - Table-driven FSM implementation
  - `transitions.rs` - State transition table and types
  - `turnstile_fsm.rs` - Main FSM implementation
- `src/main.rs` - Example application
- `tests/table_impl_tests.rs` - Integration tests

## Features

### States
- `Locked` - Initial state, turnstile is locked
- `Unlocked` - Turnstile is unlocked after coin insertion

### Events
- `Coin` - A coin is inserted
- `Pass` - Someone attempts to pass through

### Transition Table
| Current State | Event | New State | Action |
|---------------|-------|-----------|--------|
| Locked        | Coin  | Unlocked  | unlock |
| Locked        | Pass  | Locked    | alarm  |
| Unlocked      | Coin  | Unlocked  | thanks |
| Unlocked      | Pass  | Locked    | lock   |

## Usage

### Running the Demo
```bash
cargo run
```

### Running Tests
```bash
cargo test
```

### Running Specific Tests
```bash
# Unit tests
cargo test --lib

# Integration tests
cargo test --test table_impl_tests

# Binary tests
cargo test --bin finite-state-machine
```

## Architecture

### Core Components

#### TurnstileMechanism Trait
```rust
pub trait TurnstileMechanism {
    fn lock(&mut self);
    fn unlock(&mut self);
    fn alarm(&mut self);
    fn thanks(&mut self);
}
```

#### FSM Implementation
```rust
pub struct TurnstileFSM {
    current_state: State,
    turnstile: Box<dyn TurnstileMechanism>,
    transitions: HashMap<State, HashMap<Event, Transition>>,
}
```

### Key Design Decisions

1. **Ownership Model**: Uses `Box<dyn TurnstileMechanism>` for owned trait objects
2. **State Management**: Enums for compile-time safety of states and events
3. **Function Pointers**: Uses function pointers for actions in transition table
4. **Testing Strategy**: 
   - Unit tests in implementation module
   - Integration tests using `Rc<RefCell<>>` for shared state tracking
   - Fake objects pattern for testing

### Rust-Specific Features

- **Type Safety**: Enums prevent invalid state/event combinations
- **Memory Safety**: No manual memory management required
- **Pattern Matching**: Used for transition lookups
- **Trait Objects**: Enable polymorphism for different turnstile implementations

## Comparison with Go Implementation

| Aspect | Go | Rust |
|--------|----|----- |
| State/Event Types | `int` constants | Type-safe enums |
| Error Handling | Runtime panics | Compile-time prevention |
| Memory Management | GC | Ownership system |
| Polymorphism | Interfaces | Traits |
| Testing | Manual state tracking | `Rc<RefCell<>>` pattern |

## Building

### Prerequisites
- Rust 1.70+ (uses 2021 edition)

### Commands
```bash
# Check code
cargo check

# Build
cargo build

# Build optimized
cargo build --release

# Run with output
cargo run

# Test with verbose output
cargo test -- --nocapture
```

## Future Improvements

1. Add error handling for invalid transitions
2. Implement state entry/exit actions
3. Add state machine visualization
4. Support for hierarchical states
5. Add metrics/logging for state transitions
