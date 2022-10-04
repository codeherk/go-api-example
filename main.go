package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// TODO: Define Task Struct
type task struct {
	Description      string `json:"description"`
	CreatedTimestamp string `json:"createdTimestamp"`
}

// TODO: Define tasks variable globally
var tasks = []task{
	{Description: "Get 10k followers", CreatedTimestamp: time.Now().Format(time.RFC3339)},
	{Description: "Get paid from content", CreatedTimestamp: time.Now().Format(time.RFC3339)},
}

// TODO: Define getTasks Handler
func getTasks(w http.ResponseWriter, r *http.Request) {
	// Marshal tasks into json
	resp, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Error marshalling tasks to json: %s", err.Error())
	}

	// Return response
	w.Write(resp)
}

func main() {
	r := mux.NewRouter()

	// TODO: Register new route
	r.HandleFunc("/tasks", getTasks).Methods("GET")

	log.Print("Now serving on port 8090")
	err := http.ListenAndServe(":8090", r)

	if err != nil {
		log.Fatal(err)
	}
}
