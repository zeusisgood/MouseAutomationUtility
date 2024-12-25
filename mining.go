package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("長押ししたいキーを選択してください。")
	fmt.Println("1: 左クリック, 2: 右クリック")
	r := bufio.NewReader(os.Stdin)
	for {
		input, _ := r.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "1" {
			leftclick()
		} else if input == "2" {
			rightClick()
		} else {
			fmt.Println("1か2を入力してください。")
			continue
		}
	}
}

func leftclick() {
	fmt.Println("10秒後に左クリック長押しを開始します。")
	fmt.Println("プログラムを終了するには 'ctrl + c' を入力してください。")
	time.Sleep(10 * time.Second)
	robotgo.Toggle("left")
	waitforinput()
}

func rightClick() {
	fmt.Println("10秒後に右クリック長押しを開始します。")
	fmt.Println("プログラムを終了するには 'ctrl + c' を入力してください。")
	time.Sleep(10 * time.Second)
	robotgo.Toggle("right")
	waitforinput()
}

func waitforinput() {
	r := bufio.NewReader(os.Stdin)
	_, _ = r.ReadString('\n')
}
