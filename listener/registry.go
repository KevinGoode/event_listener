package event_listener

import (
	"sync"
)

//Registry is a container for message proxies and registry server
type Registry struct {
	proxies []ProxyAPI
	lock    *sync.Mutex
}

//AlreadyRegisteredClientAndMessage returns any existing proxy that supports clientID and messageID provided
func (registry *Registry) AlreadyRegisteredClientAndMessage(clientID string, messageID string) ProxyAPI {
	for _, proxy := range registry.proxies {
		if proxy.IsClientAndMessage(clientID, messageID) {
			return proxy
		}
	}
	return nil
}

//DeleteProxy returns any existing proxy that supports clientID and messageID provided
func (registry *Registry) DeleteProxy(clientID string, messageID string) error {
	registry.lock.Lock()
	for pos, proxy := range registry.proxies {
		if proxy.IsClientAndMessage(clientID, messageID) {
			//Stop proxy then delete it
			proxy.Stop()
			registry.proxies = append(registry.proxies[:pos], registry.proxies[pos+1:]...)
		}
	}
	registry.lock.Unlock()
	return nil
}

//AddNewProxy adds a new proxy to registry
func (registry *Registry) AddNewProxy(handler *Handler) error {
	registry.lock.Lock()
	proxy := NewMessageProxy(handler)
	err := proxy.Start()
	registry.proxies = append(registry.proxies, proxy)
	registry.lock.Unlock()
	return err
}

//Stop stops all proxies and stops registry server
func (registry *Registry) Stop() error {
	var err error
	for _, proxy := range registry.proxies {
		errorProxy := proxy.Stop()
		if errorProxy != nil && err == nil {
			//Just return first error found
			//Need to improve this
			err = errorProxy
		}
	}
	return err
}

//NewRegistry creates a new Registry component
func NewRegistry() *Registry {
	registry := Registry{}
	registry.lock = &sync.Mutex{}
	return &registry
}
