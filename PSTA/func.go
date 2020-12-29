package main

import (
	"fmt"
)

func main() {
	input := []int{1,2,3,4,5,6,7}
	odd_func := func(i int) bool {return i%2 == 1}
	even_func := func(i int) bool {return i%2 == 0}

	fmt.Println(filter(input, odd_func))
	fmt.Println(filter(input, even_func))

	fibNum := 10

	fmt.Println(fib_it(fibNum))

	fmt.Println(fib_rec(fibNum))

	counterIncrease3 := counter_closure(3)
	fmt.Println(counterIncrease3())
	fmt.Println(counterIncrease3())

	pow := func(i int) int {return mult(i)(i)}
	fmt.Println(pow(3))

	fmt.Println(curried_add(3)(2))
}

func filter(input []int, filterFunction func(int) bool)(ret []int){
	for _, i := range input{
		if filterFunction(i){
			ret = append(ret, i)
		}
	}
	return
}

func fib_it(n int) int{
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		temp := a
        a = b
        b = temp + a
	}
	return a
}

func fib_rec(n int) int{
	if n == 0 || n == 1{
		return n
	} else{
		return fib_rec(n-1) + fib_rec(n-2)
	} 
}

func counter_closure(n int) func() int {
	counter:= 0
	return func() int{
		counter += n
		return counter
	}
}

func curried_add(n int) func(int) int {
	return func(i int) int{
		 return n + i
	}
}

func mult(a int) func(int) int{
	return func(i int) int{
		return a * i
	}
}