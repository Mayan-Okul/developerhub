package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// This is what a Task struct might look like in Go
type Task struct {
	ID  	int    `json:"id"`
   Title    string  `json:"title"`
}

func main() {
	// 1. Endpoint to get tasks
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		// Allow cross-origin requests from our frontend
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		// Sample tasks list
		tasks := []Task{
			{ID: 1, Title: "Master Go fundamentals"},
			{ID: 2, Title: "Build full-stack web apps"},
		}

		// Convert our Go data into JSON format and send it back
		json.NewEncoder(w).Encode(tasks)
	})

	// 2. Start the server
	fmt.Println("Go server running on http://localhost:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}