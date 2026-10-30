package main

import (
	"log"

	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/bootstrap/dependency"
	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/config"
	"github.com/joho/godotenv"
)

func main() {

	AppInitializr := dependency.NewInitializer()
	AppInitializr.StartDependencies()
}
