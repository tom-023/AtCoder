// 問題文
// 3 問の問題からなる試験があり、それぞれの問題の配点は 1 点、2 点、4 点でした。
// 高橋君、青木君、すぬけ君の 3 人がこの試験を受け、 高橋君は A 点、青木君は B 点を取りました。
// すぬけ君は、高橋君と青木君のうち少なくとも一方が解けた問題は解け、 2 人とも解けなかった問題は解けませんでした。
// すぬけ君の点数を求めてください。
// ただし、この問題の制約下で、すぬけ君の点数は一意に定まる事が証明できます。

// 制約
// 0≤A,B≤7
// A,B は整数
// 入力
// 入力は以下の形式で標準入力から与えられる。

// A B
// 出力
// すぬけ君の点数を整数で出力せよ。

// 入力例 1
// Copy
// 1 2
// 出力例 1
// Copy
// 3
// 高橋君は 1 点を取った事から、 1 点の問題のみを正解し、それ以外の 2 問は解けなかったことがわかります。
// 同様に、青木君は 2 点を取った事から、 2 点の問題のみを正解し、それ以外の 2 問は解けなかったことがわかります。

// よって、すぬけ君は 1 点の問題と 2 点の問題を正解し、高橋君と青木君がともに解けなかった 4 点の問題はすぬけ君も解けなかったことになるので、3 点を取ったことがわかります。よって、3 を出力します。

// 入力例 2
// Copy
// 5 3
// 出力例 2
// Copy
// 7
// 高橋君は 5 点を取った事から、 1 点の問題と 4 点の問題を正解し、 2 点の問題は解けなかったことがわかります。
// 同様に、青木君は 3 点を取った事から、 1 点の問題と 2 点の問題を正解し、 4 点の問題は解けなかったことがわかります。

// よって、3 問すべてについて、高橋君と青木君の少なくとも一方が正解しているため、すぬけ君はすベての問題に正解し、7 点を取ったことがわかります。 よって、7 を出力します。

// 入力例 3
// Copy
// 0 0
// 出力例 3
// Copy
// 0
// 高橋君と青木君は 2 人ともいずれの問題も解けていません。 よって、すぬけ君もいずれの問題も解けておらず、 0 を出力します。

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a | b) // 論理和によるビット演算
}