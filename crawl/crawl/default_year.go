package crawl

import "time"

func DefaultYear() (int, int) {
	now := time.Now()
	if now.Month() == time.December {
		return 2004, now.Year() + 1
	} else {
		return 2004, now.Year()
	}
}
