package bootstrap

import (
	"testing"
	"time"

	"github.com/arthurhenrique-Dev/multi-tenant-api/pkg"
)

func TestInitializerDependencies(t *testing.T) {
	initializer := NewInitializer()

	executed := make(chan string, 3)

	initializer.Register(InitConfig{
		name: "database",
		start: func() {
			executed <- "database"
		},
		after: pkg.NewSet[string](),
	})

	initializer.Register(InitConfig{
		name: "cache",
		start: func() {
			executed <- "cache"
		},
		after: pkg.NewSet("database"),
	})

	initializer.Register(InitConfig{
		name: "api",
		start: func() {
			executed <- "api"
		},
		after: pkg.NewSet("cache"),
	})

	initializer.StartDependencies()

	expected := []string{
		"database",
		"cache",
		"api",
	}

	for _, name := range expected {
		select {
		case result := <-executed:
			if result != name {
				t.Fatalf("expected %s, got %s", name, result)
			}

		case <-time.After(5 * time.Second):
			t.Fatalf("timeout waiting for %s", name)
		}
	}
}
