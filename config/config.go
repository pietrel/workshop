package config

import (
	"fmt"
	"os"
)

func GetDbConfig() string {
	return fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)
}

func GetSrvConfig() string {
	return fmt.Sprintf(":%s", os.Getenv("PORT"))
}
