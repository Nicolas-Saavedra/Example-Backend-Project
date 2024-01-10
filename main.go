package main

import (
	"log"

	"github.com/Nicolas-Saavedra/Example-Backend-Project/internal/utils"
)

const VERSION = "v1.0"

func main() {
	log.Printf("Running Example-Backend-Project %s", VERSION)
	utils.LogVersionWarnings(VERSION)
}
