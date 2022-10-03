// 問題文
// 与えられる 5 つの整数 A,B,C,D,E の中に何種類の整数があるかを出力してください。

// 制約
// 0≤A,B,C,D,E≤100
// 入力はすべて整数
// 入力
// 入力は以下の形式で標準入力から与えられる。

// A B C D E
// 出力
// 答えを出力せよ。

// 入力例 1
// Copy
// 31 9 24 31 24
// 出力例 1
// Copy
// 3
// 与えられる 5 つの整数 31,9,24,31,24 の中には、9,24,31 という 3 種類の整数があります。 よって、3 を出力します。

// 入力例 2
// Copy
// 0 0 0 0 0
// 出力例 2
// Copy
// 1

package main

import "fmt"

func fiveIntergers() {
	var a, b, c, d, e int
	fmt.Scanf("%d %d %d %d %d", &a, &b, &c, &d, &e)
	list := []int{a, b, c, d, e}
	set := make(map[int]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}
	fmt.Println(len(set))
}
