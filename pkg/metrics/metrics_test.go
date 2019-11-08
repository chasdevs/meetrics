package metrics

import (
	"google.golang.org/api/calendar/v3"
	"testing"
)

func TestParseEventTime(t *testing.T) {
	d := mockEventDateTime()
	parseEventDateTime(d)
}

func mockEventDateTime() *calendar.EventDateTime {
	return &calendar.EventDateTime{
		"",
		"2019-10-30T13:00:00Z",
		"UTC",
		[]string{},
		[]string{},
	}
}