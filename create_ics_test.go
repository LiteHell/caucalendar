package main

import (
	"bytes"
	"testing"
	"time"

	"github.com/apognu/gocal"
)

func TestGenerateIcs(t *testing.T) {
	year := time.Now().Year()
	schedules, _ := GetCAUSchedules(year)
	ics := GenerateIcs(schedules)

	t.Log(ics)

	reader := bytes.NewReader([]byte(ics))
	ical := gocal.NewParser(reader)
	ical.SkipBounds = true
	err := ical.Parse()

	if err != nil {
		t.Error(err)
	} else if len(ical.Events) == 0 {
		t.Error("Empty ICal")
	}
}
