// 問題文
// 数直線の原点に高橋君がいます。高橋君は座標 X にあるゴールに移動しようとしています。

// 座標 Y には壁があり、最初、高橋君は壁を超えて移動することができません。
// 座標 Z にあるハンマーを拾った後でなら、壁を破壊して通過できるようになります。

// 高橋君がゴールに到達することが可能か判定し、可能であれば移動距離の最小値を求めてください。

// 制約
// −1000≤X,Y,Z≤1000
// X,Y,Z は相異なり、いずれも 0 でない
// 入力に含まれる値は全て整数である
// 入力
// 入力は以下の形式で標準入力から与えられる。

// X Y Z
// 出力
// 高橋君がゴールに到達することが可能であれば、移動距離の最小値を出力せよ。不可能であれば、かわりに -1 と出力せよ。

// 入力例 1
// Copy
// 10 -10 1
// 出力例 1
// Copy
// 10
// 高橋君はまっすぐゴールに向かうことができます。

// 入力例 2
// Copy
// 20 10 -10
// 出力例 2
// Copy
// 40
// ゴールは壁の向こう側にあります。まずハンマーを拾い、壁を壊すことでゴールに到達することができます。

// 入力例 3
// Copy
// 100 1 1000
// 出力例 3
// Copy
// -1

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tom-023/AtCoder/module"
)

func hammer() {
	rand.Seed(time.Now().UnixNano())
	// x := rand.Intn(100)
	// y := rand.Intn(100)
	// z := rand.Intn(100)

	// fmt.Print(x)
	// fmt.Print("\n")
	// fmt.Print(y)
	// fmt.Print("\n")
	// fmt.Print(z)
	// fmt.Print("\n")
	var x, y, z int
	fmt.Scanf("%d %d %d", &x, &y, &z)

	if y < 0 {
		x *= -1
		y *= -1
		z *= -1
	}

	if x < y {
		fmt.Print(module.Abs(x))
	} else {
		if z > y {
			fmt.Print(-1)
		} else {
			fmt.Print(calculate(x, z))
		}
	}
}

func calculate(x, z int) int {
	return module.Abs(z) + module.Abs(x-z)
}

// func abs(i int) int {
// 	return int(math.Abs(float64(i)))
// }
