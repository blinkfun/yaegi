package main

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	println(a[1]) // 2
	for k := 0; k < 2; k++ {
		for i, v := range a {
			println(v)
			if i == 3 {
				break
			}
		}
	}
}

// Output:
// 2
// 1
// 2
// 3
// 4
// 1
// 2
// 3
// 4
