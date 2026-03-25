package database

import (
	"fmt"

	"github.com/vitali-q/hotels-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(env *config.Env) (*gorm.DB, error) {
	certsDir := "/certs-cloud" // default for prod
	//fmt.Println("env.AppEnv:", env.AppEnv)
	//fmt.Println("env.DBHost:", env.DBHost)
	if env.AppEnv == "dev" {
		//fmt.Println("env.AppEnv2:", env.AppEnv)
		certsDir = "/certs"
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert=%s sslcert=%s sslkey=%s",
		env.DBHost,
		env.DBUser,
		env.DBPassword,
		env.DBName,
		env.DBPort,
		env.DBSSLMode,
		fmt.Sprintf("%s/ca.crt", certsDir),
		fmt.Sprintf("%s/client.%s.crt", certsDir, env.DBUser),
		fmt.Sprintf("%s/client.%s.key", certsDir, env.DBUser),
	)

	fmt.Println("DB HOST:", env.DBHost)
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

