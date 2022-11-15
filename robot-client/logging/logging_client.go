package logging

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	log_server *string
)

func Init(server *string) {
	log_server = server
}

func Log(entry LogEntry) {
	go logThread(entry)
}

func logThread(entry LogEntry) {
	// FIXME Add retry logic
	// FIXME spawn one goroutine to service queue of requests instead of 1 thread per
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(entry)
	if err != nil {
		fmt.Println("Failed to encode log entry:", entry)
		return
	}
	resp, err := http.Post(*log_server+"/logs/entries", "application/json", &buf)
	if err != nil {
		fmt.Println("Failed to push to cloud:", entry)
	}
	defer resp.Body.Close()
}
