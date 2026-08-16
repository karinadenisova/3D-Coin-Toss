// coin_toss.rs — Rust версия

use rand::Rng;
use std::io::{self, Write};
use std::thread;
use std::time::Duration;

const COIN_FRAMES: [&str; 2] = [
    r#"
   _____
  /     \
 |  (●)  |
 |  (●)  |
  \_____/"#,
    r#"
   _____
  /     \
 |  ( )  |
 |  ( )  |
  \_____/"#,
];

struct CoinToss3D {
    heads: u32,
    tails: u32,
    total: u32,
}

impl CoinToss3D {
    fn new() -> Self {
        CoinToss3D { heads: 0, tails: 0, total: 0 }
    }

    fn clear_screen(&self) {
        print!("\x1B[2J\x1B[1;1H");
        io::stdout().flush().unwrap();
    }

    fn animate_coin(&mut self, result: &str) {
        for i in 0..6 {
            self.clear_screen();
            println!("🪙 3D Coin Toss (Rust)");
            println!("Бросок...");
            println!("{}", COIN_FRAMES[i % 2]);
            thread::sleep(Duration::from_millis(80));
        }

        self.clear_screen();
        println!("🪙 3D Coin Toss (Rust)");
        if result == "heads" {
            println!("\x1b[32m");
            println!("   _____");
            println!("  /     \\");
            println!(" |  (●)  |");
            println!(" |  (●)  |");
            println!("  \\_____/");
            println!("\n✅ Результат: ОРЁЛ!\x1b[0m");
            self.heads += 1;
        } else {
            println!("\x1b[31m");
            println!("   _____");
            println!("  /     \\");
            println!(" |  ( )  |");
            println!(" |  ( )  |");
            println!("  \\_____/");
            println!("\n❌ Результат: РЕШКА!\x1b[0m");
            self.tails += 1;
        }
        self.total += 1;
        println!("\n📊 Статистика: Орёл: {}, Решка: {}", self.heads, self.tails);
    }

    fn flip(&mut self) {
        let result = if rand::thread_rng().gen_bool(0.5) { "heads" } else { "tails" };
        self.animate_coin(result);
    }

    fn run(&mut self) {
        println!("🪙 3D Coin Toss (Rust)");
        println!("Нажмите Enter, чтобы бросить монету (q для выхода)");

        loop {
            print!("> ");
            io::stdout().flush().unwrap();
            let mut input = String::new();
            io::stdin().read_line(&mut input).unwrap();
            let input = input.trim().to_lowercase();
            if input == "q" {
                println!("До свидания!");
                break;
            }
            self.flip();
        }
    }
}

fn main() {
    let mut game = CoinToss3D::new();
    game.run();
}
