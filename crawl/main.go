package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"litehell.info/caucalendar/crawl/crawl"
)

func HandleLambdaEvent() error {
	schedules := crawl.FetchAllYears()
	jsonBytes, err := json.Marshal(*schedules)

	if err != nil {
		panic(err)
	}

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String("calendar.puang.network"),
		Key:         aws.String("events.json"),
		ContentType: aws.String("application/json"),
		Body:        bytes.NewReader(jsonBytes),
		Metadata: map[string]string{
			"Content-Type": "application/json",
		},
	})

	return nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
