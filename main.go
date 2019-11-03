package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/chasdevs/meetrics/pkg/data"
	"github.com/chasdevs/meetrics/pkg/metrics"
	"github.com/chasdevs/meetrics/pkg/util"
	log "github.com/sirupsen/logrus"
)

/*
 * TODO: Identify and save rooms.
 * TODO: Organize the code.
 * TODO: Add unit tests!!!
 * TODO: Link users to meetings in database via the EventMaps.
 * TODO: Handle overlapping events for users.

 * TODO: Background reading on NPS to inform the meeting nps question: "was the meeting too long or too short?" (too short => high value!)
 * Look for "remnant" windows of 30min or so between meetings. "Swiss cheese factor". Look for "dead space". A single hour is sorta "low value" time. Above 90 min you get higher value.
 */

// Initialization

func init() {
	log.SetLevel(log.DebugLevel)
	data.Init()
}

// Main

func main() {
	computeYesterday()
}

func computeYesterday() {
	date := util.BeginningOfYesterday()
	if util.IsWeekday(date) {
		metrics.CompileMetrics(date)
	}
}
