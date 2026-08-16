// coin_toss.cs — C# версия

using System;
using System.Threading;

class CoinToss3D
{
    private static readonly string[] COIN_FRAMES = {
        @"
   _____
  /     \
 |  (●)  |
 |  (●)  |
  \_____/",
        @"
   _____
  /     \
 |  ( )  |
 |  ( )  |
  \_____/"
    };

    private int heads = 0;
    private int tails = 0;
    private int total = 0;

    private void ClearScreen()
    {
        Console.Clear();
    }

    private void AnimateCoin(string result)
    {
        for (int i = 0; i < 6; i++)
        {
            ClearScreen();
            Console.WriteLine("🪙 3D Coin Toss (C#)");
            Console.WriteLine("Бросок...");
            Console.WriteLine(COIN_FRAMES[i % 2]);
            Thread.Sleep(80);
        }

        ClearScreen();
        Console.WriteLine("🪙 3D Coin Toss (C#)");
        if (result == "heads")
        {
            Console.ForegroundColor = ConsoleColor.Green;
            Console.WriteLine("   _____");
            Console.WriteLine("  /     \\");
            Console.WriteLine(" |  (●)  |");
            Console.WriteLine(" |  (●)  |");
            Console.WriteLine("  \\_____/");
            Console.WriteLine("\n✅ Результат: ОРЁЛ!");
            Console.ResetColor();
            heads++;
        }
        else
        {
            Console.ForegroundColor = ConsoleColor.Red;
            Console.WriteLine("   _____");
            Console.WriteLine("  /     \\");
            Console.WriteLine(" |  ( )  |");
            Console.WriteLine(" |  ( )  |");
            Console.WriteLine("  \\_____/");
            Console.WriteLine("\n❌ Результат: РЕШКА!");
            Console.ResetColor();
            tails++;
        }
        total++;
        Console.WriteLine($"\n📊 Статистика: Орёл: {heads}, Решка: {tails}");
    }

    private void Flip()
    {
        Random rand = new Random();
        string result = rand.Next(2) == 0 ? "heads" : "tails";
        AnimateCoin(result);
    }

    public void Run()
    {
        Console.WriteLine("🪙 3D Coin Toss (C#)");
        Console.WriteLine("Нажмите Enter, чтобы бросить монету (q для выхода)");

        while (true)
        {
            Console.Write("> ");
            string input = Console.ReadLine()?.Trim().ToLower();
            if (input == "q")
            {
                Console.WriteLine("До свидания!");
                break;
            }
            Flip();
        }
    }

    public static void Main()
    {
        new CoinToss3D().Run();
    }
}
