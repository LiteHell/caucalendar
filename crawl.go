package main

import "time"

func crawlWorker() {
	for {
		time.Sleep(time.Hour * 1)
		fetchAllYears()
	}
}

func setupCrawller() {
	go crawlWorker()
}
