package main

import "fmt"

func crawlYear(year int, resultCh chan<- *[]CAUSchedule) {

	fmt.Printf("Working on year %d\n", year)
	schedules, err := GetCAUSchedules(year)
	if err != nil {
		panic(fmt.Errorf("Initial crawlling failure on year %d: %s", year, err))
	}

	resultCh <- schedules
}

func crawlAllYears() {
	start, end := DefaultYear()
	events := []CAUSchedule{}
	resultsCh := make(chan *[]CAUSchedule, end-start+1)
	for i := start; i <= end; i++ {
		go crawlYear(i, resultsCh)
	}

	for i := start; i <= end; i++ {
		result := <-resultsCh
		events = append(events, *result...)
	}

	unique := getUniqueOnly(&events)
	fmt.Printf("Inserting into database (%d events)...\n", len(unique))
	err := insertRows(&unique)
	if err != nil {
		panic(fmt.Errorf("Initial database insertion failure: %s", err))
	}
}
