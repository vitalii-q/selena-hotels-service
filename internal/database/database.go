package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {

	user := os.Getenv("HOTELS_COCKROACH_USER")

	certsDir := "/certs" // default for dev
	if os.Getenv("PROJECT_SUFFIX") == "prod" {
		certsDir = "/certs-cloud"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s sslcert=%s sslkey=%s",
		os.Getenv("HOTELS_COCKROACH_HOST"),
		user,
		os.Getenv("HOTELS_COCKROACH_PASSWORD"),
		os.Getenv("HOTELS_COCKROACH_DB_NAME"),
		os.Getenv("HOTELS_COCKROACH_PORT_INNER"),
		os.Getenv("DB_SSLMODE"),
		fmt.Sprintf("%s/ca.crt", certsDir),
		fmt.Sprintf("%s/client.%s.crt", certsDir, user),
		fmt.Sprintf("%s/client.%s.key", certsDir, user),
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

