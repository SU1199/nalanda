package worker

import (
	"log"
	"time"
)

func bot_(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Println("worker", id, "started the job ", j, " at", time.Now().Second())
		time.Sleep(2 * time.Second)
		log.Println("worker", id, "finished the job ", j, " at", time.Now().Second())
		results <- j
	}
}

func test_() {
	inp := []int{102199001, 102199440, 32429945, 423324323, 4234234, 423423423, 23423423, 324234234, 234234, 23423, 4234234, 2345, 34, 5656, 345, 3455, 2345, 345234, 5345435, 345, 345, 345, 345, 34, 534, 534, 45, 634, 564, 6345, 6745, 634, 645, 634, 645, 634, 56453, 634, 534, 45, 68, 567, 56}
	jobs := make(chan int, len(inp))
	results := make(chan int, len(inp))
	for w := 1; w <= 3; w++ {
		go bot_(w, jobs, results)
	}
	for _, j := range inp {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= len(inp); a++ {
		<-results
	}
}
