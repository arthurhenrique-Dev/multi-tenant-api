package dependency

import (
	"sync"

	"github.com/arthurhenrique-Dev/multi-tenant-api/internal/platform/config"
	"github.com/arthurhenrique-Dev/multi-tenant-api/pkg"
)

type InitConfig struct {
	name  string
	start func(*config.Config) (any, error)
	after pkg.Set[string]
}

func NewInitConfig(name string, start func(*config.Config) (any, error), after pkg.Set[string]) *InitConfig {
	return &InitConfig{
		name:  name,
		start: start,
		after: after,
	}
}

type Initializer struct {
	mu        sync.RWMutex
	items     map[string]InitConfig
	started   pkg.Set[string]
	completed pkg.Set[string]
	config    *config.Config
}

func NewInitializer() *Initializer {
	return &Initializer{
		items:     make(map[string]InitConfig),
		started:   pkg.NewSet[string](),
		completed: pkg.NewSet[string](),
		config:    nil,
	}
}

func (i *Initializer) Register(config InitConfig) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.items[config.name] = config
}

func (i *Initializer) StartDependencies(config *config.Config) {
	i.config = config
	i.executor(
		func(config InitConfig) bool {
			return config.after.Len() == 0
		},
	)
}
func (i *Initializer) schedule(item InitConfig) {
	i.mu.Lock()

	if i.started.Contains(item.name) {
		i.mu.Unlock()
		return
	}

	i.started.Add(item.name)

	i.mu.Unlock()

	go i.run(item)
}
func (i *Initializer) run(item InitConfig) {
	item.start(i.config)
	i.complete(item.name)
}

func (i *Initializer) complete(name string) {
	i.mu.Lock()
	i.completed.Add(name)
	i.mu.Unlock()

	i.executor(
		func(config InitConfig) bool {
			return config.after.HaveAllSet(i.completed)
		},
	)
}

func (i *Initializer) executor(
	check func(InitConfig) bool,
) {
	i.mu.RLock()

	var ready []InitConfig

	for _, config := range i.items {
		if i.started.Contains(config.name) {
			continue
		}

		if check(config) {
			ready = append(ready, config)
		}
	}

	i.mu.RUnlock()

	for _, config := range ready {
		i.schedule(config)
	}
}
