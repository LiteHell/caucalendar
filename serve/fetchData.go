package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"litehell.info/caucalendar/crawl/crawl"
)

func fetchData() *[]crawl.CAUSchedule {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	client := s3.NewFromConfig(cfg)
	obj, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String("calendar.puang.network"),
		Key:    aws.String("events.json"),
	})

	result := make([]crawl.CAUSchedule, 0)

	jsonDecoder := json.NewDecoder(obj.Body)
	jsonDecoder.Decode(&result)

	return &result
}
