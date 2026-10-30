package boot

import (
	"log"

	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/bootstrap/dependency"
	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/config"
	"github.com/joho/godotenv"
)

var Dependencies = dependency.NewInitializer()

func Startup() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	config, _ := config.Load()
	Dependencies.StartDependencies(&config)
}
