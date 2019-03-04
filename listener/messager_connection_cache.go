package event_listener

import (
	"sync"
)

type messageConnectionConfig struct {
	clientID   string
	connection MessagerConnectionAPI
}
type messagerConnectionCache struct {
	cache []messageConnectionConfig
}

//MessagerConnectionCacheAPI API that definesconnection cache service
type MessagerConnectionCacheAPI interface {
	GetConnection(clientID string) MessagerConnectionAPI
	AddNewConnection(clientID string, connection MessagerConnectionAPI)
}

var singleInstance MessagerConnectionCacheAPI
var once sync.Once

func (connectionCache *messagerConnectionCache) GetConnection(clientID string) MessagerConnectionAPI {
	for _, connection := range connectionCache.cache {
		if connection.clientID == clientID {
			return connection.connection
		}
	}
	return nil
}
func (connectionCache *messagerConnectionCache) AddNewConnection(clientID string, connection MessagerConnectionAPI) {
	cachedConnection := messageConnectionConfig{}
	cachedConnection.clientID = clientID
	cachedConnection.connection = connection
	connectionCache.cache = append(connectionCache.cache, cachedConnection)
}

//GetMessagerConnectionCacheInstance returns single instance of messagerConnectionCache
func GetMessagerConnectionCacheInstance() MessagerConnectionCacheAPI {
	//Using once.Do means that this singleton is threadsafe
	once.Do(func() {
		singleInstance = &messagerConnectionCache{}
	})
	return singleInstance
}
