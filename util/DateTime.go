package util

import "time"

func currentYear() string {
	dt := time.Now()
	return dt.Format("2006")
}

func currentDate() string {
	dt := time.Now()
	return dt.Format("2006-01-02")
}

func currentTime() string {
	dt := time.Now()
	return dt.Format("15꞉04꞉05")
}
