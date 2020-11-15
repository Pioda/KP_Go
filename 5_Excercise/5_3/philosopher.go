package main

import (
	"time"
	"fmt"
)

func main(){
	var f []bool
	for i := 0; i < 5; i++ {
		f = append(f, true)
	}
	var c []chan bool
	for i := 0; i < 5; i++ {
		c = append(c, make(chan bool))
	}
	var t = Table{
		forks: f,
		reqFork: make(chan int),
		putFork: make(chan int),
		commChan: c,
	}
	for i := 0; i < 5; i++ {
		philosopher := new(Philosopher)
		philosopher.id = i
		philosopher.table = &t
		go philosopher.run()
	}
	go t.run()

	time.Sleep(20000) 
	// es isst nur einer immer
}

type Philosopher struct{
	id int
	table *Table
}

type Table struct{
	forks []bool
	reqFork chan int
	putFork chan int
	commChan []chan bool
}

func (t *Table) run(){
	select{
	case id:= <-t.reqFork:
		var secondIndex = id - 1
		if(id == 0){
			secondIndex = len(t.forks) - 1
		}
		if(t.forks[secondIndex] && t.forks[id]){
			t.forks[secondIndex] = false
			t.forks[id] = false
			t.commChan[id] <- true
		} else{
			t.commChan[id] <- false
		}
	case id:= <- t.putFork:
		var secondIndex = id - 1
		if(id == 0){
			secondIndex = len(t.forks) - 1
		}
		t.forks[secondIndex] = true
		t.forks[id] = true
	}
}

func (p *Philosopher) run(){
	for {
		p.getFork()
		p.eat()
		p.putFork()
		p.think()
	}
}

func (p *Philosopher) eat(){
	fmt.Println("%v eats...", p.id)
	time.Sleep(2000)
}

func (p *Philosopher) think(){
	fmt.Println("%v thinks...", p.id)
	time.Sleep(2000)
}

func (p *Philosopher) getFork(){
	gotFork := false
	for !gotFork {
		p.table.reqFork <- p.id
	  	gotFork = <-p.table.commChan[p.id]
	}
}

func (p *Philosopher) putFork(){
	p.table.putFork <- p.id
}

