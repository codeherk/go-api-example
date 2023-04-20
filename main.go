package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Database object for handlers to use
var db *sql.DB

type Task struct {
	Description      string `json:"description"`
	CreatedTimestamp string `json:"createdTimestamp"`
}

// TODO: Implement API Error Handling
// type ErrorResponse struct {
// 	Error   string `json:"error"`
// 	Message string `json:"message"`
// 	Detail  string `json:"detail"`
// }

func getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("Attempting query to Task table")

	// Query Task Table
	rows, err := db.Query("SELECT description, created_timestamp FROM Tasks")
	if err != nil {
		log.Printf("Error querying Task table: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Description, &task.CreatedTimestamp)
		if err != nil {
			log.Printf("Error converting rows: %s\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	// if err := rows.Err(); err != nil {
	// 	panic(err.Error())
	// }

	for _, task := range tasks {
		// fmt.Printf("Task ID: %d\n", task.ID)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Created Timestamp: %s\n", task.CreatedTimestamp)
	}

	// Marshal tasks into json
	resp, err := json.Marshal(tasks)

	if err != nil {
		log.Printf("Error marshalling tasks to json: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Return response
	w.Write(resp)
}

func main() {
	r := mux.NewRouter()

	// Get MYSQL variables
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	log.Println("Attempting to open MySQL connection")

	// Establish DB connection. Format: USER:PASSWORD@tcp(HOST:PORT)/DATABASE
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database))

	if err != nil {
		log.Fatalf("Error opening MySQL connection: %s", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatalf("Error pinging MySQL database: %s", err.Error())
	} else {
		log.Println("Successful ping")
	}

	// TODO: Register new route
	r.HandleFunc("/health", getHealth).Methods("GET")
	r.HandleFunc("/tasks", getTasks).Methods("GET")

	log.Println("Now serving on port 8090")
	err = http.ListenAndServe(":8090", r)

	if err != nil {
		log.Fatal(err)
	}
}
