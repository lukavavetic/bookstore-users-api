package date

import "time"

const (
	dateFormat = "2006-01-02T15:05:05Z"
)

func GetNowString() string {
	return time.Now().UTC().Format(dateFormat)
}