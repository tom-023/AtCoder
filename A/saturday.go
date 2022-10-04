// 問題文
// ある日、学校へ行くのに疲れてしまった高橋くんは、土曜日まであと何日あるかを知りたくなりました。
// その日は平日で、曜日を英語で表すと S だったことが分かっています。その日より後の直近の土曜日は何日後かを求めてください。

// なお、月曜日、火曜日、水曜日、木曜日、金曜日はそれぞれ英語で Monday, Tuesday, Wednesday, Thursday, Friday です。

// 制約
// S は Monday, Tuesday, Wednesday, Thursday, Friday のいずれかである
// 入力
// 入力は以下の形式で標準入力から与えられる。

// S
// 出力
// 答えを整数として出力せよ。

// 入力例 1
// Copy
// Wednesday
// 出力例 1
// Copy
// 3
// この日は水曜日なので、3 日後に土曜日になります。

// 入力例 2
// Copy
// Monday
// 出力例 2
// Copy
// 5

package main

import "fmt"

func saturday() {
	var dayOfWeek string
	fmt.Scan(&dayOfWeek)
	var afterNum int
	switch dayOfWeek {
	case "Monday":
		afterNum = 5
	case "Tuesday":
		afterNum = 4
	case "Wednesday":
		afterNum = 3
	case "Thursday":
		afterNum = 2
	case "Friday":
		afterNum = 1
	}

	fmt.Println(afterNum)
}
