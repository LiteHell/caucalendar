package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func generateUid(schedule *CAUSchedule) string {
	return fmt.Sprintf("%x@caucalendar.online",
		md5.Sum([]byte(fmt.Sprintf("%d_%d_%s",
			schedule.StartDate.Unix(),
			schedule.EndDate.Unix(),
			schedule.Title))),
	)
}

func GenerateIcs(schedules *[]CAUSchedule) string {
	// Start VCALENDAR
	result := "BEGIN:VCALENDAR\n" +
		"VERSION:2.0\n" +
		"X-WR-CALNAME:중앙대학교 학사일정\n" +
		"X-WR-CALDESC:caucalendar.online에서 제공하는 중앙대학교 학사일정\n" +
		"CALSCALE:GREGORIAN\n" +
		"PRODID:adamgibbons/ics\n" +
		"METHOD:PUBLISH\n" +
		"X-PUBLISHED-TTL:PT1H\n"

	creationTimestamp := time.Now().Format("20060102T150405Z")

	for _, schedule := range *schedules {
		vEventEndData := ""
		if !schedule.EndDate.Equal(schedule.StartDate) {
			vEventEndData = fmt.Sprintf(
				"DTEND;VALUE=DATE:%s\n",
				schedule.EndDate.Format("20060102"),
			)
		}
		result +=
			fmt.Sprintf("BEGIN:VEVENT\n"+
				"UID:%s\n"+
				"SUMMARY:%s\n"+
				"DTSTAMP:%s\n"+
				"DTSTART;VALUE=DATE:%s\n"+
				vEventEndData+
				"END:VEVENT\n",
				generateUid(&schedule),
				schedule.Title,
				creationTimestamp,
				schedule.StartDate.Format("20060102"),
			)
	}

	// End VCALENDAR
	result += "END:VCALENDAR"
	return result
}
