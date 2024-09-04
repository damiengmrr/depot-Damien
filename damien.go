package main

import "fmt"

func zero() int {
	return 0
}


func max(a,b int) int {
	if a>b {
		return a
	} else {
		return b
	}
}

func max3(a int,b int,c int) int {
	if a>b {
		max_ab = a // donc a>b
	} else {
		max_ab = b // donc b>a
	}
	if max_ab > c {
		return max_ab // car "A" ou "B" est > a "C"
	} else {
		return c // c est > a "A" et "B"
	}

}

func main() {
	fmt.Println(zero())
	fmt.Println(max3(2,3,4))
	fmt.Println(max(2,3))
}