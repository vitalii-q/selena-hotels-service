package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // регистрация импорта для побочных эффектов
	"github.com/vitali-q/hotels-service/internal/handlers"
	//"github.com/sirupsen/logrus"
)

var db *sql.DB

func main() {
	// Создаем новый экземпляр Gin
	r := gin.Default()

	//logrus.SetLevel(logrus.DebugLevel)           // Установка уровня логирования
	//logrus.SetFormatter(&logrus.TextFormatter{ 
	//	FullTimestamp: true,                     // Красивый вывод логов
	//})

	// Пример маршрута для API
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Hotel Service is running",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})

	//logrus.Debug("qwer1")

    // Подключение к БД
	var err error
	connStr := "postgresql://hotels_user:hotels_pass@hotels_db:26257/hotels?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	r.GET("/health/db", func(c *gin.Context) {
		//logrus.Warn("qwer23")

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

	handlers.RegisterHotelRoutes(r)

	// Настроим сервер на прослушивание порта 8080
	if err := r.Run(":9064"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}