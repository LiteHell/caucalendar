package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type CAUSchedule struct {
	StartDate time.Time
	EndDate   time.Time
	Title     string
}

type cauScheduleRawItem struct {
	SUBJECT string
	START_Y string
	START_M int
	START_D int
	END_Y   string
	END_M   int
	END_D   int
}

type cauScheduleRawResponse struct {
	Data []cauScheduleRawItem `json:"data"`
}

type cauScheduleRequest struct {
	SCH_SITE_NO int
	SCH_YEAR    int
}

func GetCAUSchedules(year int) (*[]CAUSchedule, error) {
	const apiUrl string = "https://www.cau.ac.kr/ajax/FR_SCH_SVC/ScheduleListData.do"

	// Generate request body
	reqBody := []byte(fmt.Sprintf("SCH_SITE_NO=2&SCH_YEAR=%d", year))

	// Fetch api
	httpResp, err := http.Post(apiUrl, "application/x-www-form-urlencoded", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Convert api response into string
	var apiResp cauScheduleRawResponse
	json.NewDecoder(httpResp.Body).Decode(&apiResp)

	// Load timezone
	tz, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return nil, err
	}

	// Make objects
	result := make([]CAUSchedule, len(apiResp.Data))
	for idx, i := range apiResp.Data {
		startYear, _ := strconv.Atoi(i.START_Y)
		endYear, _ := strconv.Atoi(i.END_Y)

		result[idx].StartDate = time.Date(startYear, time.Month(i.START_M), i.START_D, 0, 0, 0, 0, tz)
		result[idx].EndDate = time.Date(endYear, time.Month(i.END_M), i.END_D, 0, 0, 0, 0, tz)
		result[idx].Title = i.SUBJECT
	}

	return &result, nil
}
