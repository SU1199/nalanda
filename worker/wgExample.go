package worker

import (
	"log"
	"sync"
	"time"
)

//  Wait groups ->

func wgEg() {
	now := time.Now()
	var wg sync.WaitGroup
	wg.Add(4)
	go t4(&wg)
	go t3(&wg)
	go t2(&wg)
	go t1(&wg)
	wg.Wait()
	log.Println(time.Since(now).Milliseconds())
}

func t1(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	log.Println("t1")
}

func t2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(200 * time.Millisecond)
	log.Println("t2")
}

func t3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(300 * time.Millisecond)
	log.Println("t3")
}

func t4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(400 * time.Millisecond)
	log.Println("t4")
}

//Channels ->

// func chEg() {
// 	now := time.Now()
// 	done := make(chan struct{})
// 	go t4(done)
// 	go t3(done)
// 	go t2(done)
// 	go t1(done)
// 	<-done
// 	<-done
// 	<-done
// 	<-done
// 	log.Println(time.Since(now).Milliseconds())
// }

// func t1(done chan struct{}) {
// 	time.Sleep(100 * time.Millisecond)
// 	log.Println("t1")
// 	done <- struct{}{}
// }

// func t2(done chan struct{}) {
// 	time.Sleep(200 * time.Millisecond)
// 	log.Println("t2")
// 	done <- struct{}{}
// }

// func t3(done chan struct{}) {
// 	time.Sleep(300 * time.Millisecond)
// 	log.Println("t3")
// 	done <- struct{}{}
// }

// func t4(done chan struct{}) {
// 	time.Sleep(400 * time.Millisecond)
// 	log.Println("t4")
// 	done <- struct{}{}
// }
