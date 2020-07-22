package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := []int{5, 4, 1232, 312}

	work := make(chan int)
	results := make(chan int, len(jobs))

	for i := 0; i < 3; i++ {
		go worker(i+1, work, results)
	}

	for _, job := range jobs {
		fmt.Printf("Sending work %d\n", job)
		work <- job
	}

	for i := 0; i < len(jobs); i++ {
		fmt.Printf("Received result: %d\n", <-results)
	}
}

func worker(id int, work chan int, results chan int) {
	for {
		fmt.Printf("(worker %d) Waiting....\n", id)

		// blocking
		newWork := <-work
		fmt.Printf("(worker %d) Received work: %d\n", id, newWork)

		// calculations
		result := newWork + 10

		fmt.Printf("(worker %d) Did some work: %d\n", id, newWork)

		// won't block until buffer is reached
		results <- result

		// rest
		time.Sleep(time.Second)
	}
}
