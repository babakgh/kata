mod game;

fn main() {
    let mut game = game::Game::new();

    game.roll(0);

    println!("Score: {}", game.score())
}
