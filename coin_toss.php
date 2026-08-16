<?php
// coin_toss.php — PHP версия

$COIN_FRAMES = [
    '
   _____
  /     \
 |  (●)  |
 |  (●)  |
  \_____/',
    '
   _____
  /     \
 |  ( )  |
 |  ( )  |
  \_____/'
];

class CoinToss3D {
    private $heads = 0;
    private $tails = 0;
    private $total = 0;

    private function clearScreen() {
        echo "\033[2J\033[1;1H";
    }

    private function animateCoin($result) {
        for ($i = 0; $i < 6; $i++) {
            $this->clearScreen();
            echo "🪙 3D Coin Toss (PHP)\n";
            echo "Бросок...\n";
            echo $COIN_FRAMES[$i % 2] . "\n";
            usleep(80000);
        }

        $this->clearScreen();
        echo "🪙 3D Coin Toss (PHP)\n";
        if ($result == 'heads') {
            echo "\033[32m";
            echo "   _____\n";
            echo "  /     \\\n";
            echo " |  (●)  |\n";
            echo " |  (●)  |\n";
            echo "  \\_____/\n";
            echo "\n✅ Результат: ОРЁЛ!\033[0m\n";
            $this->heads++;
        } else {
            echo "\033[31m";
            echo "   _____\n";
            echo "  /     \\\n";
            echo " |  ( )  |\n";
            echo " |  ( )  |\n";
            echo "  \\_____/\n";
            echo "\n❌ Результат: РЕШКА!\033[0m\n";
            $this->tails++;
        }
        $this->total++;
        echo "\n📊 Статистика: Орёл: {$this->heads}, Решка: {$this->tails}\n";
    }

    private function flip() {
        $result = rand(0, 1) == 0 ? 'heads' : 'tails';
        $this->animateCoin($result);
    }

    public function run() {
        echo "🪙 3D Coin Toss (PHP)\n";
        echo "Нажмите Enter, чтобы бросить монету (q для выхода)\n";

        while (true) {
            echo "> ";
            $input = trim(fgets(STDIN));
            $input = strtolower($input);
            if ($input == 'q') {
                echo "До свидания!\n";
                break;
            }
            $this->flip();
        }
    }
}

$game = new CoinToss3D();
$game->run();
?>
