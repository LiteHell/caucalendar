package crawl

import "fmt"

func fetchYear(year int, resultCh chan<- *[]CAUSchedule) {

	fmt.Printf("Working on year %d\n", year)
	schedules, err := GetCAUSchedules(year)
	if err != nil {
		panic(fmt.Errorf("Initial crawlling failure on year %d: %s", year, err))
	}

	resultCh <- schedules
}

func FetchAllYears() *[]CAUSchedule {
	start, end := DefaultYear()
	events := []CAUSchedule{}
	resultsCh := make(chan *[]CAUSchedule, end-start+1)
	for i := start; i <= end; i++ {
		go fetchYear(i, resultsCh)
	}

	for i := start; i <= end; i++ {
		result := <-resultsCh
		events = append(events, *result...)
	}

	unique := getUniqueOnly(&events)
	return &unique
}
