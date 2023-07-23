package main

import (
	"testing"
	"time"
)

func TestGetCAUSchedules(t *testing.T) {
	year := time.Now().Year()
	schedules, err := GetCAUSchedules(year)

	if err != nil {
		t.Error(err)
		return
	} else if len(*schedules) == 0 {
		t.Error("Empty schedules")
		return
	}

	for _, i := range *schedules {
		t.Logf("Schedule: %s from %s to %s", i.Title, i.StartDate.String(), i.EndDate.String())
	}
}
