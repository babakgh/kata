#[derive(Debug)]
pub struct Game {
    rolls: Vec<u32>,
}

impl Default for Game {
    fn default() -> Self {
        Game::new()
    }
}

impl Game {
    pub fn new() -> Self {
        Game { rolls: Vec::new() }
    }

    pub fn roll(&mut self, pins: u32) {
        self.rolls.push(pins);
    }

    pub fn score(&self) -> u32 {
        self.score_in_details().iter().sum()
    }

    fn score_in_details(&self) -> Vec<u32> {
        let mut detailed_scores = Vec::new();
        (1..=10).fold(0, |first_in_frame, _frame| {
            detailed_scores.push(self.calculated_score_in_frame(first_in_frame));

            first_in_frame + (if self.is_strike(first_in_frame) { 1 } else { 2 })
        });
        detailed_scores
    }

    fn calculated_score_in_frame(&self, first_in_frame: usize) -> u32 {
        if self.is_strike(first_in_frame) {
            10 + self.rolls[first_in_frame + 1] + self.rolls[first_in_frame + 2]
        } else if self.is_spare(first_in_frame) {
            10 + self.rolls[first_in_frame + 2]
        } else {
            self.rolls[first_in_frame] + self.rolls[first_in_frame + 1]
        }
    }

    fn is_strike(&self, first_in_frame: usize) -> bool {
        self.rolls[first_in_frame] == 10
    }

    fn is_spare(&self, first_in_frame: usize) -> bool {
        self.rolls[first_in_frame] + self.rolls[first_in_frame + 1] == 10
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_when_a_gutter_game_played() {
        let mut game = Game::new();

        for _ in 1..=20 {
            game.roll(0);
        }

        assert_eq!(game.score(), 0);
    }

    #[test]
    fn test_when_all_rolls_are_ones_played() {
        let mut game = Game::new();

        for _ in 1..=20 {
            game.roll(1);
        }

        assert_eq!(game.score(), 20);
    }

    #[test]
    fn test_when_a_spare_is_played() {
        let mut game = Game::new();

        // spare
        game.roll(5);
        game.roll(5);

        game.roll(3); // the roll after spare
        for _ in 1..=17 {
            game.roll(0);
        }

        assert_eq!(game.score(), 16);
    }

    #[test]
    fn test_when_a_strike_is_played() {
        let mut game = Game::new();

        // strike
        game.roll(10);

        // the frame after strike
        game.roll(3);
        game.roll(4);

        for _ in 1..=16 {
            game.roll(0);
        }

        assert_eq!(game.score(), 24);
    }

    #[test]
    fn test_when_a_perfect_game_is_played() {
        let mut game = Game::new();

        for _ in 1..=12 {
            game.roll(10);
        }

        assert_eq!(game.score(), 300);
    }

    #[test]
    fn test_when_a_game_is_played() {
        let mut game = Game::new();

        // strike
        game.roll(10);

        // the frame after strike
        game.roll(3);
        game.roll(4);

        for _ in 1..=14 {
            game.roll(0);
        }

        // spare
        game.roll(5);
        game.roll(5);

        game.roll(5); // roll extra ball

        assert_eq!(game.score(), 39);
    }
}
