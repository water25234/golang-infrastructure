package time

import (
	"time"
)

// GetCurrentDate means
func GetCurrentDate() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrentUnix means
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// GetCurrentMilliUnix means
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetCurrentNanoUnix means
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}
