package main

import (
	"ais/internal/app/handler"
	"ais/internal/app/repository"
	"ais/internal/app/service"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Драйвер для PostgreSQL
)

func main() {
	// Подключение к базе данных PostgreSQL
	//connStr := "host=host.docker.internal port=5432 user=postgres dbname=tb password=123434 sslmode=disable"
	connStr := "host=db port=5432 user=postgres dbname=tb password=123434 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot reach the database: ", err)
	}
	fmt.Println("Connected to the database")

	meetingRepo := repository.NewMeetingRepository(db)

	meetingService := service.NewMeetingService(meetingRepo)

	meetingHandler := handler.NewMeetingHandler(meetingService)

	http.HandleFunc("/meetings", meetingHandler.MeetingsHandler)
	http.HandleFunc("/meetings/", meetingHandler.DeleteMeetingHandler)
	fmt.Println("Server is listening on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
