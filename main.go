package main

import (
	"bootcamp-api-hmsi/connectDB"
	"bootcamp-api-hmsi/query"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_DRIVER := os.Getenv("DB_DRIVER")
	PORT := os.Getenv("PORT")

	log.Info().Msg(DB_HOST)
	log.Info().Msg(DB_PORT)
	log.Info().Msg(DB_USER)
	log.Info().Msg(DB_PASSWORD)
	log.Info().Msg(DB_NAME)
	log.Info().Msg(DB_DRIVER)
	log.Info().Msg(PORT)

	db, err := connectDB.GetConnPostgres(DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME, DB_DRIVER)

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	fmt.Println("Succesfully connected")

	// DB struct initializes
	DB := query.DB{Conn: db}

	// update
	err = DB.Delete(3)

	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}

	fmt.Println("succeeded")
}
