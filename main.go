package main

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
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
		_, yearTo := DefaultYear()
		schedules, err := GetCAUSchedules(yearTo)
		if err != nil {
			handleIntervalError(resp, req, 500, "Internal server error")
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

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Fprintln(os.Stderr, err)
	}
}
