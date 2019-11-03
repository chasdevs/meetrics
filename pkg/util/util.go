package util

import (
	"path"
	"runtime"
	"strings"
	"time"
)

func BeginningOfYesterday() time.Time {
	return BeginningOfDay(1)
}

func BeginningOfDay(daysAgo int) time.Time {
	now := time.Now()
	year, month, yesterday := now.AddDate(0, 0, -1*daysAgo).Date()
	return time.Date(year, month, yesterday, 0, 0, 0, 0, now.Location())
}

func IsWeekday(date time.Time) bool {
	return date.Weekday() > 0 && date.Weekday() < 6
}

// https://rosettacode.org/wiki/Strip_control_codes_and_extended_characters_from_a_string#Go
func StripCtlAndExtFromUTF8(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= 32 && r < 127 {
			return r
		}
		return -1
	}, str)
}

// Filepath
func ThisFilePath() string {
	_, filename, _, _ := runtime.Caller(1)
	filename = path.Join(filename, "..")
	return filename
}
