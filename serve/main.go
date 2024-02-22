package main

import (
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"litehell.info/caucalendar/crawl/crawl"
)

func filterSchedule(allSchedules *[]crawl.CAUSchedule, fromParam string, toParam string) *[]crawl.CAUSchedule {
	// Get default parameters
	defaultYearFrom, defaultYearTo := crawl.DefaultYear()

	// Use default parameter value if not provided or invalid
	yearFrom, yearTo := defaultYearFrom, defaultYearTo
	if fromParam != "" {
		var err error
		yearFrom, err = strconv.Atoi(fromParam)
		if err != nil || yearFrom < defaultYearFrom {
			yearFrom = defaultYearFrom
		}
	}
	if toParam != "" {
		var err error
		yearTo, err = strconv.Atoi(toParam)
		if err != nil || yearTo > defaultYearTo {
			yearTo = defaultYearTo
		}
	}

	// Swap parameters if they're somewaht wrong
	if yearFrom > yearTo {
		yearFrom, yearTo = yearTo, yearFrom
	}

	// Filter schedules
	tz, _ := time.LoadLocation("Asia/Seoul")
	from := time.Date(yearFrom, 1, 1, 0, 0, 0, 0, tz)
	to := time.Date(yearTo, 12, 31, 23, 59, 59, 59, tz)

	filtered := make([]crawl.CAUSchedule, 0)
	for _, i := range *allSchedules {
		if i.StartDate.Compare(from) >= 0 && i.EndDate.Compare(to) <= 0 {
			filtered = append(filtered, i)
		}
	}

	return &filtered
}

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	allSchedules := fetchData()
	filtered := filterSchedule(allSchedules, request.QueryStringParameters["from"], request.QueryStringParameters["to"])
	ics := GenerateIcs(filtered)

	return events.APIGatewayProxyResponse{
		Body:       ics,
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "text/calendar",
		},
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
