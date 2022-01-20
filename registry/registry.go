package registry

import (
	"sync"
	"time"
)

// Registry registe the key-value pair
type Registry struct {
	lock sync.Mutex
	apps map[string]Application
}

// Application application is the server detail information
type Application struct {
	appid          string
	latestTimestap int64
	instances      map[string]Instance
}

// Instance instance is contract information of server
type Instance struct {
	Env     string `json:"env"`
	AppID   string `json:"appID"`
	Address string `json:"address"`
	Status  string `json:"status"`
}

var registry *Registry = NewRegistry()

func NewRegistry() *Registry {
	return &Registry{
		apps: make(map[string]Application),
	}
}

func OnceRegistry() *Registry {
	if registry == nil {
		return NewRegistry()
	}
	return registry
}

func Register(i *Instance) {
	if registry == nil {
		return
	}
	if _, ok := registry.apps[i.AppID]; ok {
		return
	}
	registry.register(i)
}

func (r *Registry) register(i *Instance) {
	r.lock.Lock()
	defer r.lock.Unlock()
	// double check lock
	if _, ok := r.apps[i.AppID]; ok {
		return
	}
	a := Application{
		appid:          i.AppID,
		latestTimestap: time.Now().Unix(),
		instances:      make(map[string]Instance),
	}
	a.instances[i.AppID] = *i
	r.apps[i.AppID] = a
}
