package nintaco

import "fmt"

var eventTypes = [...]int{
	activate,
	deactivate,
	stop,
	access,
	controllers,
	frame,
	scanline,
	scanlineCycle,
	spriteZero,
	status,
}

const (
	eventRequest  = 0xFF
	eventResponse = 0xFE
	heartbeat     = 0xFD
	ready         = 0xFC

	retryMillis = 1000
)

type remoteAPI struct {
	listenerIDs     map[interface{}]int
	listenerObjects map[int]map[int]interface{}
	host            string
	port            int
	stream          *dataStream
	nextID          int
	running         bool
}

func newRemoteAPI(host string, port int) *remoteAPI {
	r := &remoteAPI{
		listenerIDs:     make(map[interface{}]int),
		listenerObjects: make(map[int]map[int]interface{}),
		host:            host,
		port:            port,
	}
	for eventType := range eventTypes {
		r.listenerObjects[eventType] = make(map[int]interface{})
	}
	return r
}

func (r *remoteAPI) Run() {
	if r.running {
		return
	}
	r.running = true
	for {
		// TODO
	}
}

func (r *remoteAPI) fireDeactivated() {
	listeners := r.listenerObjects[deactivate]
	length := len(listeners)
	deactivateListeners := make([]DeactivateListener, 0, length)
	for _, listener := range listeners {
		deactivateListeners = append(deactivateListeners, listener.(DeactivateListener))
	}
	for i := length - 1; i >= 0; i-- {
		deactivateListeners[i].ApiDisabled()
	}
}

func (r *remoteAPI) fireStatusChanged(message string, params ...interface{}) {
	msg := fmt.Sprintf(message, params...)
	listeners := r.listenerObjects[status]
	length := len(listeners)
	statusListeners := make([]StatusListener, 0, length)
	for _, listener := range listeners {
		statusListeners = append(statusListeners, listener.(StatusListener))
	}
	for i := length - 1; i >= 0; i-- {
		statusListeners[i].StatusChanged(msg)
	}
}

func (r *remoteAPI) sendReady() {
	if r.stream != nil {
		r.stream.writeByte(ready)
		r.stream.flush()
	}
}

func (r *remoteAPI) sendListener(listenerID, eventType int, listenerObject interface{}) {
	if r.stream != nil {
		r.stream.writeByte(eventType)
		r.stream.writeInt(listenerID)
		switch eventType {
		case access:
			point := listenerObject.(*accessPoint)
			r.stream.writeInt(point.accessPointType)
			r.stream.writeInt(point.minAddress)
			r.stream.writeInt(point.maxAddress)
			r.stream.writeInt(point.bank)
		case scanline:
			point := listenerObject.(*scanlinePoint)
			r.stream.writeInt(point.scanline)
		case scanlineCycle:
			point := listenerObject.(*scanlineCyclePoint)
			r.stream.writeInt(point.scanline)
			r.stream.writeInt(point.scanlineCycle)
		}
		r.stream.flush()
	}
}

func (r *remoteAPI) addListener(listener interface{}, eventType int) {
	if listener != nil {
		r.sendListener(r.addListenerObject(listener, eventType, listener), eventType, listener)
	}
}

func (r *remoteAPI) removeListener(listener interface{}, eventType, methodValue int) {
	if listener != nil {
		listenerID := r.removeListenerObject(listener, eventType)
		if listenerID >= 0 && r.stream != nil {
			r.stream.writeByte(methodValue)
			r.stream.writeInt(listenerID)
			r.stream.flush()
		}
	}
}

func (r *remoteAPI) addListenerObject(listener interface{}, eventType int, listenerObject interface{}) int {
	listenerID := r.nextID
	r.nextID++
	r.listenerIDs[listener] = listenerID
	r.listenerObjects[eventType][listenerID] = listenerObject
	return listenerID
}

func (r *remoteAPI) removeListenerObject(listener interface{}, eventType int) int {
	if listenerID, contains := r.listenerIDs[listener]; contains {
		delete(r.listenerIDs, listener)
		delete(r.listenerObjects[eventType], listenerID)
		return listenerID
	}
	return -1
}

func (r *remoteAPI) GetPixels(pixels []int) {
	r.stream.writeByte(119)
	r.stream.flush()
	r.stream.readIntArray(pixels)
}
