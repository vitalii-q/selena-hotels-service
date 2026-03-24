package database

import (
	"fmt"
	"os"

	"github.com/vitali-q/hotels-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) (*gorm.DB, error) {
	certsDir := "/certs-cloud" // default for prod
	if os.Getenv("PROJECT_SUFFIX") == "dev" {
		certsDir = "/certs"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s sslcert=%s sslkey=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
		fmt.Sprintf("%s/ca.crt", certsDir),
		fmt.Sprintf("%s/client.%s.crt", certsDir, cfg.DBUser),
		fmt.Sprintf("%s/client.%s.key", certsDir, cfg.DBUser),
	)

	fmt.Println("DB HOST:", os.Getenv("HOTELS_COCKROACH_HOST"))
	fmt.Println("CERTS DIR:", certsDir)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open failed: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}
	
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return db, nil
}

