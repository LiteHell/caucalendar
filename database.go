package main

import (
	"container/list"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

var db *sql.DB

func initializeDB() error {
	var err error
	db, err = sql.Open("sqlite", ":memory:")

	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS schedules (title TEXT, start INTEGER, end TEXT, UNIQUE (title, start, end))")

	if err != nil {
		return err
	}

	return nil
}

func readRows(from, to time.Time) (*[]CAUSchedule, error) {
	rows, err := db.Query("SELECT * FROM schedules WHERE start >= ? AND end <= ?", from.Unix(), to.Unix())
	if err != nil {
		return nil, err
	}

	list := list.New()
	for rows.Next() {
		schedule := CAUSchedule{}
		var startUnix, endUnix int64
		rows.Scan(&schedule.Title, &startUnix, &endUnix)
		schedule.StartDate = time.Unix(startUnix, 0)
		schedule.EndDate = time.Unix(endUnix, 0)
		list.PushBack(&schedule)
	}

	result := make([]CAUSchedule, list.Len())
	idx := 0
	for i := list.Front(); i != nil; i = i.Next() {
		result[idx] = *i.Value.(*CAUSchedule)
		idx += 1
	}

	return &result, nil
}

func insertRows(schedules *[]CAUSchedule) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	tx.Exec("DELETE FROM schedules")

	sqlPlaceholders := make([]string, len(*schedules))
	sqlArgs := make([]interface{}, len(sqlPlaceholders)*3)
	for i, schedule := range *schedules {
		sqlPlaceholders[i] = "(?, ?, ?)"
		sqlArgs[i*3] = schedule.Title
		sqlArgs[i*3+1] = schedule.StartDate.Unix()
		sqlArgs[i*3+2] = schedule.EndDate.Unix()
	}

	_, err = tx.Exec("INSERT INTO schedules (title, start, end) VALUES "+strings.Join(sqlPlaceholders, ","), sqlArgs...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s! Rollbacking...", err)
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
