# coin_toss.rb — Ruby версия

COIN_FRAMES = [
  "
   _____
  /     \\
 |  (●)  |
 |  (●)  |
  \\_____/",
  "
   _____
  /     \\
 |  ( )  |
 |  ( )  |
  \\_____/"
]

class CoinToss3D
  def initialize
    @heads = 0
    @tails = 0
    @total = 0
  end

  def clear_screen
    system('clear') || system('cls')
  end

  def animate_coin(result)
    6.times do |i|
      clear_screen
      puts "🪙 3D Coin Toss (Ruby)"
      puts "Бросок..."
      puts COIN_FRAMES[i % 2]
      sleep 0.08
    end

    clear_screen
    puts "🪙 3D Coin Toss (Ruby)"
    if result == 'heads'
      puts "\e[32m"
      puts "   _____"
      puts "  /     \\"
      puts " |  (●)  |"
      puts " |  (●)  |"
      puts "  \\_____/"
      puts "\n✅ Результат: ОРЁЛ!\e[0m"
      @heads += 1
    else
      puts "\e[31m"
      puts "   _____"
      puts "  /     \\"
      puts " |  ( )  |"
      puts " |  ( )  |"
      puts "  \\_____/"
      puts "\n❌ Результат: РЕШКА!\e[0m"
      @tails += 1
    end
    @total += 1
    puts "\n📊 Статистика: Орёл: #{@heads}, Решка: #{@tails}"
  end

  def flip
    result = rand(2) == 0 ? 'heads' : 'tails'
    animate_coin(result)
  end

  def run
    puts "🪙 3D Coin Toss (Ruby)"
    puts "Нажмите Enter, чтобы бросить монету (q для выхода)"

    loop do
      print "> "
      input = gets.chomp.strip.downcase
      break if input == 'q'
      flip
    end
    puts "До свидания!"
  end
end

game = CoinToss3D.new
game.run
