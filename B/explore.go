// 問題文
// 高橋君はゲームの中で洞窟を探索しています。

// 洞窟は N 個の部屋が一列に並んだ構造であり、入り口から順に部屋 1,2,…,N と番号がついています。

// 最初、高橋君は部屋 1 におり、持ち時間 は T です。
// 各 1≤i≤N−1 について、持ち時間を Ai消費することで、部屋 i から部屋 i+1 へ移動することができます。これ以外に部屋を移動する方法はありません。
// また、持ち時間が 0 以下になるような移動は行うことができません。

// 洞窟内には M 個のボーナス部屋があります。i 番目のボーナス部屋は部屋 Xiであり、この部屋に到達すると持ち時間が Yi増加します。
// 高橋君は部屋 N にたどりつくことができますか？

package main

import (
	"bufio"
	"fmt"
	"os"
)

func explore() {
	in := bufio.NewReader(os.Stdin)

	var n, m, t, x, y int
	fmt.Fscan(in, &n, &m, &t)
	rooms := make([]int, n-1)
	bonusRooms := make([]int, n-1)

	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &rooms[i])
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &x, &y)
		bonusRooms[x-1] = y
	}

	for i, v := range rooms {
		t = t - v + bonusRooms[i]
		if t <= 0 {
			fmt.Print("No")
			return
		}
	}

	fmt.Print("Yes")
}
