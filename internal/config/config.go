package config

import (
	"os"
)

type APICofig struct {
	Port string
}

func LoadConfig() *APICofig {
	return &APICofig{
		Port: os.Getenv("PORT"),
	}
}
