// 問題文
// 整数 a,b,c,d が与えられるので、以下の指示に従って 2 行出力してください。

// 1 行目は (a+b)×(c−d) の計算結果を整数として出力してください。
// 2 行目は入力にかかわらず Takahashi と出力してください。

// 制約
// −100≤a,b,c,d≤100
// a,b,c,d は整数
// 入力
// 入力は以下の形式で標準入力から与えられる。

// a b c d
// 出力
// 問題文の指示に従って 2 行出力せよ。

// 入力例 1
// Copy
// 1 2 5 3
// 出力例 1
// Copy
// 6
// Takahashi
// (1+2)×(5−3)=3×2=6 です。よって 1 行目は 6 を出力します。
// 2 行目は Takahashi を出力します。1 文字目を小文字にしたりスペルを誤ったりすると誤答となるので注意してください。

// 入力例 2
// Copy
// 10 -20 30 -40
// 出力例 2
// Copy
// -700
// Takahashi
// 入出力に負の数が含まれる場合もあります。

// 入力例 3
// Copy
// 100 100 100 -100
// 出力例 3
// Copy
// 40000
// Takahashi

package main

import "fmt"

func anywayTakahashi() {
	var a, b, c, d int
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	fmt.Println((a + b) * (c - d))
	fmt.Println("Takahashi")
}
