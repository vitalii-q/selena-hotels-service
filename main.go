package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // регистрация импорта для побочных эффектов
	"github.com/vitali-q/hotels-service/internal/database"
	"github.com/vitali-q/hotels-service/internal/handlers"
	//"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"github.com/sirupsen/logrus"
)

var DB *gorm.DB

/*func InitDB() error {
    dsn := "host=hotels-db user=hotels_user password=hotels_pass dbname=hotels_db port=26257 " +
       "sslmode=verify-full sslrootcert=/certs/ca.crt " +
       "sslcert=/certs/client.root.crt sslkey=/certs/client.root.key"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }
    DB = db
    return nil
}*/

func main() {
	// Создаем новый экземпляр Gin
	r := gin.Default()

	//logrus.SetLevel(logrus.DebugLevel)           // Установка уровня логирования
	//logrus.SetFormatter(&logrus.TextFormatter{ 
	//	FullTimestamp: true,                     // Красивый вывод логов
	//})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, hotels-service!")
	})

	// Инициализация базы данных через GORM
	log.Println("Initializing DB...")
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("DB initialized successfully")

	//logrus.Debug("qwer1")

	// Проверка подключения к БД
	r.GET("/health/db", func(c *gin.Context) {
		sqlDB, err := database.DB.DB() // получаем *sql.DB из GORM
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get raw DB"})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Database ping failed"})
			return
		}
		c.String(http.StatusOK, "Hotels-service: database connection OK")
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
