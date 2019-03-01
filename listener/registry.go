package event_listener

//Registry is a container for message proxies and registry server
type Registry struct {
	proxies []*MessageProxy
}

//AddNewProxy adds a new proxy to registry
func (registry *Registry) AddNewProxy(handler *Handler) error {
	proxy := NewMessageProxy(handler)
	err := proxy.Start()
	registry.proxies = append(registry.proxies, proxy)
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
	return &registry
}
