package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// 1. gh pr list コマンドの実行
	cmdList := exec.Command("gh", "pr", "list")
	cmdList.Stdout = os.Stdout
	cmdList.Stderr = os.Stderr

	if err := cmdList.Run(); err != nil {
		fmt.Println("Error running gh pr list:", err)
		os.Exit(1)
	}

	// 2. ユーザーにプルリクエスト番号を入力してもらう
	fmt.Print("\n\033[32m?\033[0m Type the number of PR to checkout? ... ")
	reader := bufio.NewReader(os.Stdin)
	prNumber, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	prNumber = strings.TrimSpace(prNumber) // 入力の改行や空白を削除

	// 3. 入力されたプルリクエスト番号のチェックアウト処理
	cmdCheckout := exec.Command("gh", "pr", "checkout", prNumber)
	cmdCheckout.Stdout = os.Stdout
	cmdCheckout.Stderr = os.Stderr

	if err := cmdCheckout.Run(); err != nil {
		fmt.Println("Error checking out PR:", err)
		os.Exit(1)
	}
}
