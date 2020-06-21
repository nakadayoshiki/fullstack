package api

import (
	"fmt"
	"log"
	"os"

	"github.com/nakadayoshiki/fullstack/github.com/joho/godotenv"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/controllers"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/seed"
)

var s = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	s.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(s.DB)

	s.Run(":8080")
}
