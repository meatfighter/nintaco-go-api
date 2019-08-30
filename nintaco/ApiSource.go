package nintaco

import (
	"strings"
	"sync"
)

var apiMutex = &sync.Mutex{}
var api API

// InitRemoteAPI initializes the Remote API.
func InitRemoteAPI(host string, port int) {
	apiMutex.Lock()
	defer apiMutex.Unlock()
	if api == nil && len(strings.TrimSpace(host)) > 0 {
		api = newRemoteAPI(host, port)
	}
}

// GetAPI provides the handle to the API.
func GetAPI() API {
	return api
}
