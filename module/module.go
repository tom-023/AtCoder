package module

func Abs(a int) int {
	if a > 0 {
		return a
	} else {
		return a * -1
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
