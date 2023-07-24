package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type cauCalendarHandler struct {
}

func handleIntervalError(resp http.ResponseWriter, req *http.Request, status int, msg string) {
	resp.WriteHeader(status)
	resp.Write([]byte(msg))
}

// ServeHTTP implements http.Handler.
func (handler cauCalendarHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/cau.ics" {
		defaultYearFrom, defaultYearTo := DefaultYear()
		yearFrom, yearTo := defaultYearFrom, defaultYearTo
		if req.URL.Query().Get("from") != "" {
			var err error
			yearFrom, err = strconv.Atoi(req.URL.Query().Get("from"))
			if err != nil || yearFrom < defaultYearFrom {
				yearFrom = defaultYearFrom
			}
		}

		if req.URL.Query().Get("to") != "" {
			var err error
			yearTo, err = strconv.Atoi(req.URL.Query().Get("to"))
			if err != nil || yearTo > defaultYearTo {
				yearTo = defaultYearTo
			}
		}

		if yearFrom > yearTo {
			yearFrom, yearTo = yearTo, yearFrom
		}

		tz, _ := time.LoadLocation("Asia/Seoul")
		from := time.Date(yearFrom, 1, 1, 0, 0, 0, 0, tz)
		to := time.Date(yearTo, 12, 31, 23, 59, 59, 59, tz)

		schedules, err := readRows(from, to)
		if err != nil {
			handleIntervalError(resp, req, 500, "Internal server error")
			fmt.Fprintf(os.Stderr, "Error on reading database: %s\n", err)
			return
		}

		ics := GenerateIcs(schedules)
		resp.WriteHeader(200)
		resp.Header().Add("Content-Type", "text/calendar")
		resp.Write([]byte(ics))
	} else {
		wd, err := os.Getwd()
		if err != nil {
			handleIntervalError(resp, req, 500, "Internal server error")
			return
		}

		staticDir := path.Join(wd, "static")

		targetFile := path.Join(staticDir, "index.html")
		if req.URL.Path != "/" && req.URL.Path != "/index.html" {
			targetFile = path.Join(staticDir, req.URL.Path)
			targetFile = path.Clean(targetFile)
		}
		if !strings.HasPrefix(targetFile, staticDir+"/") {
			handleIntervalError(resp, req, 500, "Fuck you")
		} else {
			file, err := os.Open(targetFile)
			if err != nil {
				handleIntervalError(resp, req, 404, "Not Found")
				return
			}

			resp.Header().Add("Content-Type", mime.TypeByExtension(path.Ext(targetFile)))
			io.Copy(resp, file)
			file.Close()
		}
	}
}

func main() {
	var server http.Server
	server.Handler = cauCalendarHandler{}
	server.Addr = ":8080"

	err := initializeDB()

	if err != nil {
		panic(
			fmt.Errorf("Database initialization failure: %s", err))
	}

	fmt.Println("Performing initial crawlling...")
	start, end := DefaultYear()
	events := []CAUSchedule{}
	for i := start; i <= end; i++ {
		fmt.Printf("Working on year %d\n", i)
		schedules, err := GetCAUSchedules(i)
		if err != nil {
			panic(fmt.Errorf("Initial crawlling failure on year %d: %s", i, err))
		}

		events = append(events, *schedules...)
	}

	fmt.Println("Inserting into database...")
	unique := getUniqueOnly(&events)
	err = insertRows(&unique)
	if err != nil {
		panic(fmt.Errorf("Initial database insertion failure: %s", err))
	}

	fmt.Println("Initial preparation Complete!")

	fmt.Printf("Listening on %s\n", server.Addr)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Fprintln(os.Stderr, err)
	}
}
