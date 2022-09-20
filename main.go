package main

import (
	"github.com/AlexCollin/CryptoComparer0/models"
	"log"
	"net/http"
	"time"
)

const tasksPerSecond = 110000000

func main() {

	source := &models.InchSources{}
	var file = source.parseFile()
	// bitrue.Url = ""
	// totalTasks := 10000000000000000
	// concurrency := 5

	// var wg sync.WaitGroup
	// wg.Add(concurrency)

	// log.Println("Starting ...")

	// for i := 0; i < concurrency; i++ {
	// 	go func(count int) {
	// 		runWorker(count, totalTasks/concurrency)
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()
	// log.Println("... Done")
}

func runWorker(tn, n int) {
	res, _ := http.Get("http://localhost:8080/tasksPerSecond?fromTokenAddress=0xdac17f958d2ee523a2206206994597c13d831ec7")
	if res == nil {

	}

	log.Printf("task: %d n: %d\n", tn, n)
	var throttle <-chan time.Time
	if tasksPerSecond > 0 {
		throttle = time.Tick(time.Duration(tasksPerSecond) * time.Microsecond)
	}

	for i := 0; i < n; i++ {
		if tasksPerSecond > 0 {
			<-throttle
		}
		log.Printf("doing task %d\n", tn)
	}
}
