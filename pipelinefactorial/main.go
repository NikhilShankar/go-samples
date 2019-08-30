package main

import (
	"log"
	"sync"
)

func main() {
	ch := make(chan int)
	go factorial(7, ch)
	var x = 1
	for num := range ch {
		x *= num
	}


	rch := make(chan int)
	go allFactorials(10, rch)
	for num := range rch {
		log.Println("The factorial is ", num)
	}


}

func allFactorials(n int, resultChan chan int) {
	sum := 0
	for i := 1; i <= n; i++ {
		lock := sync.Mutex{}
		go func(i int) {
			ch := make(chan int)
			go factorial(i, ch)
			var x = 1
			for n := range ch {
				x *= n
			}
			resultChan <- x
			lock.Lock()
			sum += 1
			log.Println(sum)
			if sum == n {
				close(resultChan)
			}
			lock.Unlock()
		}(i)
	}

}

func factorial(n int, myChan chan int) {
	if n < 2 {
		myChan <- 1
		close(myChan)
		return
	}
	myChan <- n
	factorial(n-1, myChan)
}

