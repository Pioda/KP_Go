package main

import (
	"fmt"
	"errors"
	"time"
)

func main() {
    //fibChan := fib() // <- write func fib
    //for n := 1; n <= 10; n++ {
    //    fmt.Printf("The %dth Fibonacci number is %d\n", n, <-fibChan)
	//}
	main2()
}

func fib() <-chan int{
	intChan := make(chan int) 
	go func(){
		a,b := 0,1
		for{
			intChan <- b
			a,b = b, a+b
		}
	}()
	return intChan
}

func main2() {
    res, err := setTimeout(func() int {
        time.Sleep(2000 * time.Millisecond)
        return 1
    }, 3*time.Second)

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("operation returned %d", res)
    }
}

func setTimeout(funcToRun func() int, timeoutAfter time.Duration) (int, error) {
	timeOutChan := time.After(timeoutAfter)
	funcChan := make(chan int)
	go func(){
		funcChan <- funcToRun()
	}()
	select {
		case <- timeOutChan:
			return -1, errors.New("Function timed out")
		case v := <- funcChan:
			return v, nil
	}
}