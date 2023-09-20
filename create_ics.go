package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func generateUid(schedule *CAUSchedule) string {
	kst, _ := time.LoadLocation("Asia/Seoul")
	return fmt.Sprintf("%x@calendar.puang.network",
		sha1.Sum([]byte(fmt.Sprintf("%d_%d_%d%d_%d_%d%s",
			schedule.StartDate.In(kst).Year(),
			schedule.StartDate.In(kst).Month(),
			schedule.StartDate.In(kst).Day(),
			schedule.EndDate.In(kst).Year(),
			schedule.EndDate.In(kst).Month(),
			schedule.EndDate.In(kst).Day(),
			schedule.Title))),
	)
}

func GenerateIcs(schedules *[]CAUSchedule) string {
	result :=
		// Start VCALENDAR
		"BEGIN:VCALENDAR\n" +
			"VERSION:2.0\n" +
			"TIMEZONE-ID:Asia/Seoul\n" +
			"X-WR-TIMEZONE:Asia/Seoul\n" +
			"X-WR-CALNAME:중앙대학교 학사일정\n" +
			"X-WR-CALDESC:calendar.puang.network에서 제공하는 중앙대학교 학사일정\n" +
			"CALSCALE:GREGORIAN\n" +
			"PRODID:adamgibbons/ics\n" +
			"METHOD:PUBLISH\n" +
			"X-PUBLISHED-TTL:PT1H\n" +
			// Start VTIMEZONE
			"BEGIN:VTIMEZONE\n" +
			"TZID:Asia/Seoul\n" +
			"TZURL:http://tzurl.org/zoneinfo-outlook/Asia/Seoul\n" +
			"X-LIC-LOCATION:Asia/Seoul\n" +
			"BEGIN:STANDARD\n" +
			"TZOFFSETFROM:+0900\n" +
			"TZOFFSETTO:+0900\n" +
			"TZNAME:KST\n" +
			"DTSTART:19700101T000000\n" +
			"END:STANDARD\n" +
			"END:VTIMEZONE\n"

	creationTimestamp := time.Now().Format("20060102T150405Z")
	kst, _ := time.LoadLocation("Asia/Seoul")

	for _, schedule := range *schedules {
		vEventEndData := ""
		if !schedule.EndDate.Equal(schedule.StartDate) {
			vEventEndData = fmt.Sprintf(
				"DTEND;TZID=Asia/Seoul;VALUE=DATE:%s\n",
				schedule.EndDate.In(kst).Format("20060102"),
			)
		}
		result +=
			fmt.Sprintf("BEGIN:VEVENT\n"+
				"UID:%s\n"+
				"SUMMARY:%s\n"+
				"DTSTAMP:%s\n"+
				"DTSTART;TZID=Asia/Seoul;VALUE=DATE:%s\n"+
				vEventEndData+
				"END:VEVENT\n",
				generateUid(&schedule),
				schedule.Title,
				creationTimestamp,
				schedule.StartDate.In(kst).Format("20060102"),
			)
	}

	// End VCALENDAR
	result += "END:VCALENDAR"
	return result
}
