// 問題文
// 英小文字からなる長さが奇数の文字列 S が与えられます。

// S の中央の文字を出力してください。

// 中央の文字とは
// 制約
// S は英小文字からなる長さが奇数の文字列
// S の長さは 1 以上 99 以下
// 入力
// 入力は以下の形式で標準入力から与えられる。

// S
// 出力
// 答えを出力せよ。

// 入力例 1
// Copy
// atcoder
// 出力例 1
// Copy
// o
// atcoder の中央の文字は o です。

// 入力例 2
// Copy
// a
// 出力例 2
// Copy
// a

package main

import "fmt"

func middleLetter() {
	var alphabet string
	fmt.Scan(&alphabet)
	centerNum := len(alphabet) / 2
	fmt.Print(alphabet[centerNum : centerNum+1])
}
