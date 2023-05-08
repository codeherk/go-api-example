package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// global Database pointer for handlers to use
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

func getHealth(c *gin.Context) {
	err := db.Ping()

	if err != nil {
		log.Printf("Error pinging MySQL database: %s\n", err.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println("Successful ping")
	c.Writer.WriteHeader(http.StatusOK)
}

func getTasks(c *gin.Context) {
	log.Println("Attempting query to Task table")

	// Query Task Table
	rows, err := db.Query("SELECT description, created_timestamp FROM Tasks")
	if err != nil {
		log.Printf("Error querying Task table: %s\n", err.Error())
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Description, &task.CreatedTimestamp)
		if err != nil {
			log.Printf("Error converting rows: %s\n", err.Error())
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)
	}

	// if err := rows.Err(); err != nil {
	// 	panic(err.Error())
	// }

	// for _, task := range tasks {
	// 	// fmt.Printf("Task ID: %d\n", task.ID)
	// 	fmt.Printf("Description: %s\n", task.Description)
	// 	fmt.Printf("Created Timestamp: %s\n", task.CreatedTimestamp)
	// }

	// Marshal tasks into json
	// resp, err := json.Marshal(tasks)

	// if err != nil {
	// 	log.Printf("Error marshalling tasks to json: %s\n", err.Error())
	// 	c.Writer.WriteHeader(http.StatusInternalServerError)
	// }

	// Return response
	c.JSON(http.StatusOK, tasks)
}

func main() {
	router := gin.Default()

	// Get MYSQL variables
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")

	log.Println("Attempting to open MySQL connection")

	// Establish DB connection.
	// Format: USER:PASSWORD@tcp(HOST:PORT)/DATABASE
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database))

	if err != nil {
		log.Fatalf("Error opening MySQL connection: %s", err.Error())
	}

	db = conn

	router.GET("/health", getHealth)
	router.GET("/tasks", getTasks)

	// log.Println("Now serving on port 8090")
	err = router.Run(":8090")

	if err != nil {
		log.Fatal(err)
	}
}
