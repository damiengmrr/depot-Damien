package main

import "fmt"

func zero() int {
	return 0
}

func plus_one(n int) int {
	return n + 1
}

func identity(n int) int {
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func difference(a, b int) int {
	if a > b {

		return a - b
	} else {

		return b - a
	}
}

func moyenne(x, y, z, t int) int {

	return (x + y + z + t) / 4
}

func invert(a int) int {
	return -a
}

func sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(zero())
	//	fmt.Println(max3(2,3,4))
	fmt.Println(max(2, 3))
	fmt.Println(difference(1, 2))
	fmt.Println(plus_one(3))
	fmt.Println(identity(3))
	fmt.Println(moyenne(2, 2, 4, 4))
	fmt.Println(invert(5))
	fmt.Println(sum(4, 8))

}
