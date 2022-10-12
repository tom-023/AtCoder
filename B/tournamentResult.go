// 問題文
// N 人の人が総当り戦の試合をしました。

// N 行 N 列からなる試合の結果の表 A が与えられます。A の i 行目 j 列目の要素を Ai,jと表します。
// Ai,jは i=j のとき - であり、それ以外のとき W, L, D のいずれかです。
// Ai,jが W, L, D であることは、人 i が人 j との試合に勝った、負けた、引き分けたことをそれぞれ表します。

// 与えられた表に矛盾があるかどうかを判定してください。

// 次のいずれかが成り立つとき、与えられた表には矛盾があるといいます。

// ある組 (i,j) が存在して、人 i が人 j に勝ったが、人 j が人 i に負けていない
// ある組 (i,j) が存在して、人 i が人 j に負けたが、人 j が人 i に勝っていない
// ある組 (i,j) が存在して、人 i が人 j に引き分けたが、人 j が人 i に引き分けていない

package main

import (
	"fmt"
	"strings"
)

func tournamentResult() {
	var n int
	fmt.Scanln(&n)
	var rowStrings string
	var result string
	table := make([][]string, n)
	individualPerformance := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&rowStrings)
		individualPerformance = strings.Split(rowStrings, "")
		table[i] = individualPerformance
	}

	result = "correct"
LOOP:
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if (table[i][j] == "W" && table[j][i] == "L") || (table[i][j] == "L" && table[j][i] == "W") || (table[i][j] == "D" && table[j][i] == "D") {
			} else {
				result = "incorrect"
				break LOOP
			}
		}
	}
	fmt.Print(result)
}
