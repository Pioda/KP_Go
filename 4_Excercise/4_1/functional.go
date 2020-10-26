package main

import "fmt"

func main() {
	doMath(Pow, 13)
	doMath(func(n int) int {
		if n == 1 {
			return 1
		}
		return n * (n - 1)
	}, 5)
	comp()
}

func comp() {
	square := func(x any) any { return x.(int) * x.(int) }
	compose := func(f, g function) function {
		return func(x any) any {
			return (f(g(x)))
		}
	}
	result := compose(square, square)(4)
	fmt.Println(result)
}

type any interface{}
type function func(any) any

func doMath(f func(int) int, num int) {
	result := f(num)
	fmt.Println(result)
}

func Pow(num int) int {
	return num * num
}
