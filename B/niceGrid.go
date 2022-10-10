// 問題文
// 次の図に示す、各マスが黒または白に塗られた縦 15 行 × 横 15 列のグリッドにおいて、 上から R 行目、左から C 列目のマスが何色かを出力して下さい。
package main

import (
	"fmt"

	"github.com/tom-023/AtCoder/module"
)

func niceGrid() {
	var row, column int
	fmt.Scanf("%d %d", &row, &column)
	if module.Max(module.Abs(row-8), module.Abs(column-8))%2 == 0 {
		fmt.Print("white")
	} else {
		fmt.Print("black")
	}
}
