package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/sirupsen/logrus"
)

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

	//logrus.Error("ests")
	//logrus.Debug("sfds")
	//logrus.Debug("Hotel service started")

	// Настроим сервер на прослушивание порта 8080
	if err := r.Run(":9064"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
