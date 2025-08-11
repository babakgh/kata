pub mod table_impl;

/// The TurnstileMechanism trait defines the actions that can be performed on a turnstile
pub trait TurnstileMechanism {
    fn lock(&mut self);
    fn unlock(&mut self);
    fn alarm(&mut self);
    fn thanks(&mut self);
}

/// Implement TurnstileMechanism for Box<dyn TurnstileMechanism> to enable generic FSM with trait objects
impl TurnstileMechanism for Box<dyn TurnstileMechanism> {
    fn lock(&mut self) {
        self.as_mut().lock();
    }

    fn unlock(&mut self) {
        self.as_mut().unlock();
    }

    fn alarm(&mut self) {
        self.as_mut().alarm();
    }

    fn thanks(&mut self) {
        self.as_mut().thanks();
    }
}

/// Concrete implementation of a Turnstile
pub struct Turnstile {
    pub name: String,
}

impl Turnstile {
    pub fn new(name: String) -> Self {
        Self { name }
    }
}

impl TurnstileMechanism for Turnstile {
    fn lock(&mut self) {
        println!("{}: Lock", self.name);
    }

    fn unlock(&mut self) {
        println!("{}: Unlock", self.name);
    }

    fn alarm(&mut self) {
        println!("{}: Alarm", self.name);
    }

    fn thanks(&mut self) {
        println!("{}: Thanks", self.name);
    }
}
