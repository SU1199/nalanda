package main

import (
	"flag"
	"nalanda/db"
	"nalanda/worker"
	// "nalanda/permutationGen"
)

func main() {
	// permutationGen.Generate(21)

	//flag for number of concurrent gophers
	num := flag.Int("w", 30, "number of workers")

	//connect to db
	db.ConnectDB()

	//generate enum slice
	inp := db.JobsGen()

	//make channels
	jobs := make(chan int, len(inp))
	results := make(chan int, len(inp))

	//deploy workers
	for w := 1; w <= *num; w++ {
		go worker.Bot(w, jobs, results)
	}

	//add to jobs channel
	for _, j := range inp {
		jobs <- j
	}
	close(jobs)

	//listen to result channel
	for a := 1; a <= len(inp); a++ {
		<-results
	}
}
