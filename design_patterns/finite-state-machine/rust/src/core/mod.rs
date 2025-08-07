pub mod table_impl;

/// The TurnstileMechanism trait defines the actions that can be performed on a turnstile
pub trait TurnstileMechanism {
    fn lock(&mut self);
    fn unlock(&mut self);
    fn alarm(&mut self);
    fn thanks(&mut self);
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
