package main

import (
	"TaskTracker/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTasks(w, r)
		case http.MethodPost:
			handler.CreateTask(w, r)
		case http.MethodPatch:
			handler.UpdateTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// // Database connection parameters
	// const (
	// 	host     = "localhost"
	// 	port     = 5432
	// 	user     = "postgres"
	// 	password = "test123" // Replace with your actual password
	// 	dbname   = "mydatabase"    // Replace with your actual database name
	// )

	// // Create a connection string
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// // Open a connection to the database
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	log.Fatal(err) // Use log.Fatal for connection errors
	// }
	// defer db.Close() // Ensure the connection is closed when the function exits

	// // Verify the connection is active using Ping()
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Successfully connected to the database!")

	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
