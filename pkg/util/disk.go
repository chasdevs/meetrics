package util

import (
	"bufio"
	"encoding/json"
	"google.golang.org/api/calendar/v3"
	"os"
	"path"
)

func SaveEvents(events []*calendar.Event) error {
	// open file
	filePath := path.Join(RootPath(), "data", "dump")

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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