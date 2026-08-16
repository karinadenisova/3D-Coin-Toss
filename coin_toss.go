// coin_toss.go — Go версия

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var coinFrames = []string{
	`
   _____
  /     \
 |  (●)  |
 |  (●)  |
  \_____/`,
	`
   _____
  /     \
 |  ( )  |
 |  ( )  |
  \_____/`,
}

type CoinToss3D struct {
	heads int
	tails int
	total int
}

func (c *CoinToss3D) clearScreen() {
	cmd := exec.Command("clear")
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cls")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (c *CoinToss3D) animateCoin(result string) {
	for i := 0; i < 6; i++ {
		c.clearScreen()
		fmt.Println("🪙 3D Coin Toss (Go)")
		fmt.Println("Бросок...")
		frame := coinFrames[i%2]
		fmt.Println(frame)
		time.Sleep(80 * time.Millisecond)
	}

	c.clearScreen()
	fmt.Println("🪙 3D Coin Toss (Go)")
	if result == "heads" {
		fmt.Println("\x1b[32m")
		fmt.Println("   _____")
		fmt.Println("  /     \\")
		fmt.Println(" |  (●)  |")
		fmt.Println(" |  (●)  |")
		fmt.Println("  \\_____/")
		fmt.Println("\n✅ Результат: ОРЁЛ!\x1b[0m")
		c.heads++
	} else {
		fmt.Println("\x1b[31m")
		fmt.Println("   _____")
		fmt.Println("  /     \\")
		fmt.Println(" |  ( )  |")
		fmt.Println(" |  ( )  |")
		fmt.Println("  \\_____/")
		fmt.Println("\n❌ Результат: РЕШКА!\x1b[0m")
		c.tails++
	}
	c.total++
	fmt.Printf("\n📊 Статистика: Орёл: %d, Решка: %d\n", c.heads, c.tails)
}

func (c *CoinToss3D) flip() {
	rand.Seed(time.Now().UnixNano())
	result := "heads"
	if rand.Intn(2) == 0 {
		result = "tails"
	}
	c.animateCoin(result)
}

func (c *CoinToss3D) run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🪙 3D Coin Toss (Go)")
	fmt.Println("Нажмите Enter, чтобы бросить монету (q для выхода)")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "q" {
			fmt.Println("До свидания!")
			break
		}
		c.flip()
	}
}

func main() {
	game := &CoinToss3D{}
	game.run()
}
