package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // регистрация импорта для побочных эффектов
	"github.com/sirupsen/logrus"
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

	logrus.Debug("qwer1")

    // Подключение к БД
	var err error
	connStr := "postgresql://hotels_user:hotels_pass@hotels-db:26257/hotels?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	logrus.Debug("qwer2")

	r.GET("/health/db", func(c *gin.Context) {
		if err := db.Ping(); err != nil {
			log.Printf("DB Ping error: %v", err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database connection failed"})
			return
		}
		c.String(http.StatusOK, "Database connection OK")
	})

	//logrus.Error("ests")
	//logrus.Debug("sfds")
	//logrus.Debug("Hotel service started")

	// Настроим сервер на прослушивание порта 8080
	if err := r.Run(":9064"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}