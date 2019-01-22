package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

// var transforms = []string{
// 	otherWord,
// 	otherWord,
// 	otherWord,
// 	otherWord,
// 	otherWord + "app",
// 	otherWord + "site",
// 	otherWord + "time",
// 	"get" + otherWord,
// 	"go" + otherWord,
// 	"lets" + otherWord,
// }

func main() {
	lines := readFile("test.txt")
	// 乱数のシードを生成
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := lines[rand.Intn(len(lines))]
		// 普通に出力すると文字列連結されない
		//fmt.Println(strings.Replace(t, "otherWord", s.Text(), -1))
		output := strings.Replace(t, "otherWord", s.Text(), -1)
		// 行に＋があれば文字列を連結して表示
		if strings.Contains(output, "+") {
			strargs := strings.Split(output, "+")
			fmt.Println(strargs[0] + strargs[1])
		} else {
			fmt.Println(output)
		}
	}
}

func readFile(filepath string) []string {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("ファイル読み込みに失敗しました。")
		os.Exit(1)
	}

	// 関数が終了した際に確実に閉じるようにする
	// defer　は処理を遅延実行させる
	defer f.Close()

	//Scanerで読み込み
	// とりあえず２０確保
	lines := make([]string, 0, 20)
	scanner := bufio.NewScanner(f)
	// ファイルの末尾に達してfalseを返すまでループ
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanError := scanner.Err(); scanError != nil {
		fmt.Fprintf(os.Stderr, "ファイルスキャンに失敗しました。")
	}
	return lines
}
