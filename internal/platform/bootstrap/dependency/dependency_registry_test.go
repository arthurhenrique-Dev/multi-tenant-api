package dependency

import (
	"testing"
	"time"

	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/config"
	"github.com/arthurhenrique-Dev/multi-tenant-api/pkg"
)

func TestInitializerDependencies(t *testing.T) {
	initializer := NewInitializer()
	appConfig := &config.Config{}

	type execution struct {
		name   string
		config *config.Config
	}
	executed := make(chan execution, 3)

	initializer.Register(InitConfig{
		name: "database",
		start: func(cfg *config.Config) (any, error) {
			executed <- execution{name: "database", config: cfg}
			return nil, nil
		},
		after: pkg.NewSet[string](),
	})

	initializer.Register(InitConfig{
		name: "cache",
		start: func(cfg *config.Config) (any, error) {
			executed <- execution{name: "cache", config: cfg}
			return nil, nil
		},
		after: pkg.NewSet("database"),
	})

	initializer.Register(InitConfig{
		name: "api",
		start: func(cfg *config.Config) (any, error) {
			executed <- execution{name: "api", config: cfg}
			return nil, nil
		},
		after: pkg.NewSet("database", "cache"),
	})

	initializer.StartDependencies(appConfig)

	expected := []string{
		"database",
		"cache",
		"api",
	}

	for _, name := range expected {
		select {
		case result := <-executed:
			if result.name != name {
				t.Fatalf("expected %s, got %s", name, result.name)
			}
			if result.config != appConfig {
				t.Errorf("%s received a different configuration", name)
			}

		case <-time.After(5 * time.Second):
			t.Fatalf("timeout waiting for %s", name)
		}
	}
}
