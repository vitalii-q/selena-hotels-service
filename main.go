package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/sirupsen/logrus"
)

var db *sql.DB

func main() {
	// Создаем новый экземпляр Gin
	r := gin.Default()

	// Пример маршрута для API
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Hotel Service is running",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	var err error
    // Подключение к БД (пример)
    connStr := "postgresql://user:password@hotels-db:26257/hotels?sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Cannot connect to DB: %v", err)
    }

	http.HandleFunc("/health/db", dbHealthCheckHandler)

	//logrus.Error("ests")
	//logrus.Debug("sfds")
	//logrus.Debug("Hotel service started")

	// Настроим сервер на прослушивание порта 8080
	if err := r.Run(":9064"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func dbHealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    err := db.Ping()
    if err != nil {
        http.Error(w, "Database connection failed", http.StatusServiceUnavailable)
        return
    }
    fmt.Fprintln(w, "Database connection OK")
}