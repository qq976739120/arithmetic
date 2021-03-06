package main

import "fmt"
func climbStairs(n int) int {
	if 1 == n {
		return 1
	}

	f := make([]int, n+1)
	f[1], f[2] = 1, 2

	for i := 3; i < len(f); i++ {
		f[i] = f[i-1] + f[i-2]
	}
	fmt.Println(f[n])
	return f[n]
}


func climbStairs2(n int) int {
	a, b := 1, 1

	for ; n > 0; n-- {
		a, b = b, a+b
	}
	fmt.Println(a)
	return a
}
func main(){
	climbStairs(10)
}