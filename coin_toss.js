// coin_toss.js — JavaScript версия

const readline = require('readline');

const coinFrames = [
    `
   _____
  /     \\
 |  (●)  |
 |  (●)  |
  \\_____/`,
    `
   _____
  /     \\
 |  ( )  |
 |  ( )  |
  \\_____/`
];

class CoinToss3D {
    constructor() {
        this.heads = 0;
        this.tails = 0;
        this.total = 0;
    }

    clearScreen() {
        console.clear();
    }

    animateCoin(result) {
        for (let i = 0; i < 6; i++) {
            this.clearScreen();
            console.log('🪙 3D Coin Toss (JavaScript)');
            console.log('Бросок...');
            console.log(coinFrames[i % 2]);
            // Синхронная задержка (блокирующая)
            const wait = Date.now() + 80;
            while (Date.now() < wait) {}
        }

        this.clearScreen();
        console.log('🪙 3D Coin Toss (JavaScript)');
        if (result === 'heads') {
            console.log('\x1b[32m');
            console.log('   _____');
            console.log('  /     \\');
            console.log(' |  (●)  |');
            console.log(' |  (●)  |');
            console.log('  \\_____/');
            console.log('\n✅ Результат: ОРЁЛ!\x1b[0m');
            this.heads++;
        } else {
            console.log('\x1b[31m');
            console.log('   _____');
            console.log('  /     \\');
            console.log(' |  ( )  |');
            console.log(' |  ( )  |');
            console.log('  \\_____/');
            console.log('\n❌ Результат: РЕШКА!\x1b[0m');
            this.tails++;
        }
        this.total++;
        console.log(`\n📊 Статистика: Орёл: ${this.heads}, Решка: ${this.tails}`);
    }

    flip() {
        const result = Math.random() < 0.5 ? 'heads' : 'tails';
        this.animateCoin(result);
    }

    run() {
        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout
        });

        console.log('🪙 3D Coin Toss (JavaScript)');
        console.log('Нажмите Enter, чтобы бросить монету (q для выхода)');

        rl.on('line', (input) => {
            if (input.trim().toLowerCase() === 'q') {
                console.log('До свидания!');
                rl.close();
                return;
            }
            this.flip();
        });
    }
}

const game = new CoinToss3D();
game.run();
