package nintaco

import (
	"strings"
	"sync"
)

var apiMutex = &sync.Mutex{}
var api API

func initRemoteAPI(host string, port int) {
	apiMutex.Lock()
	defer apiMutex.Unlock()
	if api == nil && len(strings.TrimSpace(host)) > 0 {
		api = newRemoteAPI(host, port)
	}
}

func getAPI() API {
	return api
}
