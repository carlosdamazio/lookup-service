package main

import (
	"os"

	"github.com/carlosdamazio/lookup-service/internal/rest/handler/server"
)

func main() {
	server.StartServer(os.Getenv("APP_HOST"), "3000")
}
