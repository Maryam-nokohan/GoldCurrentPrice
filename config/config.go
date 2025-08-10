package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init(){
	log.Println("Loading .env")
	Load(".env")
}

func GetFromEnv(key string) string {
	return os.Getenv(key)
}

func Load(file string){
	err := godotenv.Load(file)
	if err != nil {
		log.Printf("Error in loading .env : %s\n", err)
	}
}