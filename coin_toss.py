
### 1. `coin_toss.py` (Python)

```python
# coin_toss.py — Python версия

import random
import time
import os
import sys
from colorama import init, Fore, Style

init(autoreset=True)

# 3D-представление монеты (вращающиеся стороны)
COIN_FRAMES = [
    """
   _____
  /     \\
 |  (●)  |
 |  (●)  |
  \\_____/
""",
    """
   _____
  /     \\
 |  ( )  |
 |  ( )  |
  \\_____/
""",
    """
   _____
  /     \\
 |  (●)  |
 |  (●)  |
  \\_____/
""",
    """
   _____
  /     \\
 |  ( )  |
 |  ( )  |
  \\_____/
""",
    """
   _____
  /     \\
 |  (●)  |
 |  (●)  |
  \\_____/
"""
]

class CoinToss3D:
    def __init__(self):
        self.heads = 0
        self.tails = 0
        self.total = 0

    def clear_screen(self):
        """Очищает экран."""
        os.system('cls' if os.name == 'nt' else 'clear')

    def animate_coin(self, result):
        """Анимирует вращение монеты в 3D."""
        frames = COIN_FRAMES * 2  # повтор для эффекта
        for frame in frames:
            self.clear_screen()
            print("🪙 3D Coin Toss (Python)")
            print("Бросок...")
            print(frame)
            time.sleep(0.08)

        # Показываем результат с цветом
        self.clear_screen()
        print("🪙 3D Coin Toss (Python)")
        if result == 'heads':
            print(Fore.GREEN + "\n   _____")
            print("  /     \\")
            print(" |  (●)  |")
            print(" |  (●)  |")
            print("  \\_____/")
            print("\n✅ Результат: ОРЁЛ!" + Style.RESET_ALL)
            self.heads += 1
        else:
            print(Fore.RED + "\n   _____")
            print("  /     \\")
            print(" |  ( )  |")
            print(" |  ( )  |")
            print("  \\_____/")
            print("\n❌ Результат: РЕШКА!" + Style.RESET_ALL)
            self.tails += 1
        self.total += 1

        print(f"\n📊 Статистика: Орёл: {self.heads}, Решка: {self.tails}")

    def flip(self):
        """Основной метод броска."""
        result = random.choice(['heads', 'tails'])
        self.animate_coin(result)

    def run(self):
        """Запуск интерактивного режима."""
        print("🪙 3D Coin Toss (Python)")
        print("Нажмите Enter, чтобы бросить монету (q для выхода)")

        while True:
            cmd = input("\n> ").strip().lower()
            if cmd == 'q':
                print("До свидания!")
                break
            self.flip()

def main():
    game = CoinToss3D()
    game.run()

if __name__ == "__main__":
    main()
