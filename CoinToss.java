// CoinToss.java — Java версия

import java.util.Scanner;
import java.util.Random;

public class CoinToss {
    private static final String[] COIN_FRAMES = {
        "\n   _____\n  /     \\\n |  (●)  |\n |  (●)  |\n  \\_____/",
        "\n   _____\n  /     \\\n |  ( )  |\n |  ( )  |\n  \\_____/"
    };

    private int heads = 0;
    private int tails = 0;
    private int total = 0;

    private void clearScreen() {
        System.out.print("\033[H\033[2J");
        System.out.flush();
    }

    private void animateCoin(String result) throws InterruptedException {
        for (int i = 0; i < 6; i++) {
            clearScreen();
            System.out.println("🪙 3D Coin Toss (Java)");
            System.out.println("Бросок...");
            System.out.println(COIN_FRAMES[i % 2]);
            Thread.sleep(80);
        }

        clearScreen();
        System.out.println("🪙 3D Coin Toss (Java)");
        if (result.equals("heads")) {
            System.out.println("\u001B[32m");
            System.out.println("   _____");
            System.out.println("  /     \\");
            System.out.println(" |  (●)  |");
            System.out.println(" |  (●)  |");
            System.out.println("  \\_____/");
            System.out.println("\n✅ Результат: ОРЁЛ!\u001B[0m");
            heads++;
        } else {
            System.out.println("\u001B[31m");
            System.out.println("   _____");
            System.out.println("  /     \\");
            System.out.println(" |  ( )  |");
            System.out.println(" |  ( )  |");
            System.out.println("  \\_____/");
            System.out.println("\n❌ Результат: РЕШКА!\u001B[0m");
            tails++;
        }
        total++;
        System.out.printf("\n📊 Статистика: Орёл: %d, Решка: %d\n", heads, tails);
    }

    private void flip() throws InterruptedException {
        Random rand = new Random();
        String result = rand.nextBoolean() ? "heads" : "tails";
        animateCoin(result);
    }

    public void run() throws InterruptedException {
        Scanner scanner = new Scanner(System.in);
        System.out.println("🪙 3D Coin Toss (Java)");
        System.out.println("Нажмите Enter, чтобы бросить монету (q для выхода)");

        while (true) {
            System.out.print("> ");
            String input = scanner.nextLine().trim().toLowerCase();
            if (input.equals("q")) {
                System.out.println("До свидания!");
                break;
            }
            flip();
        }
        scanner.close();
    }

    public static void main(String[] args) throws InterruptedException {
        new CoinToss().run();
    }
}
