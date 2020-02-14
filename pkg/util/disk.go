package util

import (
	"bufio"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/calendar/v3"
	"os"
	"path"
)

var eventStore = path.Join(RootPath(), "data", "dump")

func SaveEvents(events []*calendar.Event) error {

	file, err := os.OpenFile(eventStore, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// new writer w/ default 4096 buffer size
	w := bufio.NewWriter(file)

	// turn events into json
	data, _ := json.Marshal(events)
	_, _ = w.Write(data)
	_, _ = w.WriteString("\n")

	// flush outstanding data
	return w.Flush()
}

func StreamEvents() []calendar.Event {
	file, _ := os.Open(eventStore)
	defer file.Close()

	events := make([]calendar.Event, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		eventSlice := make([]calendar.Event, 0)
		err := json.Unmarshal(scanner.Bytes(), &eventSlice)
		if err != nil {
			log.Error("Could not unmarshal json", err)
		}

		log.Info(eventSlice)
		_ = append(events, eventSlice...)
	}

	if err := scanner.Err(); err != nil {
		log.Error("Error scanning file.", err)
	}

	return events
}