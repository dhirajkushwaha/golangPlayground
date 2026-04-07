package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)

	ch <- 1
	go func() {
		ch <- 2
	}()

	time.Sleep(1 * time.Second)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// goroutine is parked in sendq (channel send queue) since buffer is full
//  sendq : trying to send 2




// ---------------




package main

import (
	"fmt"
	"time"
)



 





// the actual good design:

func log(msg string, logs chan string){
	select {
	case logs <- msg:

	default:
		handleBackpressure(msg)
	}
}


func writer(ch chan string){
	batch := make([]string, 0, 100)


	for{
		select {
		case log := <-ch:
			batch = append(batch,log)

			if len(batch) >= 100 {
				flush(batch)
				batch = batch[:0]
			}


		case <-time.After(10*time.Millisecond):
			if len(batch) > 0 {
				flush(batch) 
				batch = batch[:0]
			}
		}
	}
}




func main(){
	ch := make(chan int, 100)

	
	for i := 0; i < 10000000; i++{
		go func(i int){
			ch <- i
		}(i)
	}



	for i := 0; i < 10000000; i++{
		<-ch
	}
	
}


// ----------


const workers = 1000 

jobs := make(chan int, 1000)


for w := 0; w < workers; w++ {
	go func(){
		for job := range jobs{
			process(job)
		}
	}()
}


// -----


// semaphore pattern


sem := make(chan struct{}, 100)
// what does this all even mean???
for i := 0; i < 10000000; i++{
	sem <- struct{}{}


	go func(i int){
		defer func() { <- sem }()
		process(i)
	}
}








 
