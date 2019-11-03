package main

import (
	"github.com/chasdevs/meetrics/pkg/data"
)

func init() {
	data.Init()
}

func main() {
	data.Mgr.ClearMeetings()
	data.Mgr.ClearUserMeetingMins()
}
