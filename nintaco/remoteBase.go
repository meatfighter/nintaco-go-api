package nintaco

import (
	"fmt"
	"io"
	"net"
	"runtime/debug"
	"time"
)

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
	for _, eventType := range eventTypes {
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
		r.fireStatusChanged("Connecting to %s:%d...", r.host, r.port)
		conn, e := net.Dial("tcp", fmt.Sprintf("[%s]:%d", r.host, r.port))
		if e == nil {
			r.stream = newDataStream(conn)
			r.fireStatusChanged("Connection established.")
			r.sendListeners()
			r.sendReady()
			e = r.executeProbeEventsLoop()
			if e != io.EOF {
				fmt.Println(e)
			}
			r.fireDeactivated()
			r.fireStatusChanged("Disconnected.")
			conn.Close()
			r.stream = nil
		} else {
			r.fireStatusChanged("Failed to establish connection.")
		}
		time.Sleep(retryMillis * time.Millisecond)
	}
}

func (r *remoteAPI) executeProbeEventsLoop() (e error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println(string(debug.Stack()))
			e = io.EOF
		}
	}()

	for {
		if err := r.probeEvents(); err != nil {
			return err
		}
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
		deactivateListeners[i].APIDisabled()
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

func (r *remoteAPI) sendListeners() {
	for e1Key, e1Value := range r.listenerObjects {
		for e2Key, e2Value := range e1Value {
			r.sendListener(e2Key, e1Key, e2Value)
		}
	}
}

func (r *remoteAPI) probeEvents() error {

	e := r.stream.writeByte(eventRequest)
	if e != nil {
		return e
	}
	e = r.stream.flush()
	if e != nil {
		return e
	}

	eventType, e := r.stream.readByte()
	if e != nil {
		return e
	}

	if eventType == heartbeat {
		e = r.stream.writeByte(eventResponse)
		if e != nil {
			return e
		}
		e = r.stream.flush()
		if e != nil {
			return e
		}
		return nil
	}

	listenerID, e := r.stream.readInt()
	m, contains := r.listenerObjects[eventType]
	if !contains {
		return fmt.Errorf("Unknown listener type: %d", eventType)
	}
	if obj, contains := m[listenerID]; contains {
		if eventType == access {
			accessPointType, e := r.stream.readInt()
			if e != nil {
				return e
			}
			address, e := r.stream.readInt()
			if e != nil {
				return e
			}
			value, e := r.stream.readInt()
			if e != nil {
				return e
			}
			result := obj.(*accessPoint).listener.AccessPointHit(accessPointType, address, value)
			e = r.stream.writeByte(eventResponse)
			if e != nil {
				return e
			}
			e = r.stream.writeInt(result)
			if e != nil {
				return e
			}
		} else {
			switch eventType {
			case activate:
				obj.(ActivateListener).APIEnabled()
			case deactivate:
				obj.(DeactivateListener).APIDisabled()
			case stop:
				obj.(StopListener).Dispose()
			case controllers:
				obj.(ControllersListener).ControllersProbed()
			case frame:
				obj.(FrameListener).FrameRendered()
			case scanline:
				scanline, e := r.stream.readInt()
				if e != nil {
					return e
				}
				obj.(*scanlinePoint).listener.ScanlineRendered(scanline)
			case scanlineCycle:
				scanline, e := r.stream.readInt()
				if e != nil {
					return e
				}
				scanlineCycle, e := r.stream.readInt()
				if e != nil {
					return e
				}
				address, e := r.stream.readInt()
				if e != nil {
					return e
				}
				rendering, e := r.stream.readBoolean()
				if e != nil {
					return e
				}
				obj.(*scanlineCyclePoint).listener.CyclePerformed(scanline, scanlineCycle, address, rendering)
			case spriteZero:
				scanline, e := r.stream.readInt()
				if e != nil {
					return e
				}
				scanlineCycle, e := r.stream.readInt()
				if e != nil {
					return e
				}
				obj.(SpriteZeroListener).SpriteZeroHit(scanline, scanlineCycle)
			case status:
				message, e := r.stream.readString()
				if e != nil {
					return e
				}
				obj.(StatusListener).StatusChanged(message)
			default:
				return fmt.Errorf("Unknown listener type: %d", eventType)
			}
			e = r.stream.writeByte(eventResponse)
			if e != nil {
				return e
			}
		}
	}

	return r.stream.flush()
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

func (r *remoteAPI) AddActivateListener(listener ActivateListener) {
	r.addListener(listener, activate)
}

func (r *remoteAPI) RemoveActivateListener(listener ActivateListener) {
	r.removeListener(listener, activate, 2)
}

func (r *remoteAPI) AddDeactivateListener(listener DeactivateListener) {
	r.addListener(listener, deactivate)
}

func (r *remoteAPI) RemoveDeactivateListener(listener DeactivateListener) {
	r.removeListener(listener, deactivate, 4)
}

func (r *remoteAPI) AddStopListener(listener StopListener) {
	r.addListener(listener, stop)
}

func (r *remoteAPI) RemoveStopListener(listener StopListener) {
	r.removeListener(listener, stop, 6)
}

func (r *remoteAPI) AddAccessPointListener(listener AccessPointListener, accessPointType, address int) {
	r.AddAccessPointListener3(listener, accessPointType, address, -1, -1)
}

func (r *remoteAPI) AddAccessPointListener2(listener AccessPointListener, accessPointType, minAddress, maxAddress int) {
	r.AddAccessPointListener3(listener, accessPointType, minAddress, maxAddress, -1)
}

func (r *remoteAPI) AddAccessPointListener3(listener AccessPointListener, accessPointType, minAddress, maxAddress, bank int) {
	if listener != nil {
		point := newAccessPoint3(listener, accessPointType, minAddress, maxAddress, bank)
		r.sendListener(r.addListenerObject(listener, access, point), access, point)
	}
}

func (r *remoteAPI) RemoveAccessPointListener(listener AccessPointListener) {
	r.removeListener(listener, access, 10)
}

func (r *remoteAPI) AddControllersListener(listener ControllersListener) {
	r.addListener(listener, controllers)
}

func (r *remoteAPI) RemoveControllersListener(listener ControllersListener) {
	r.removeListener(listener, controllers, 12)
}

func (r *remoteAPI) AddFrameListener(listener FrameListener) {
	r.addListener(listener, frame)
}

func (r *remoteAPI) RemoveFrameListener(listener FrameListener) {
	r.removeListener(listener, frame, 14)
}

func (r *remoteAPI) AddScanlineListener(listener ScanlineListener, line int) {
	if listener != nil {
		point := newScanlinePoint(listener, line)
		r.sendListener(r.addListenerObject(listener, scanline, point), scanline, point)
	}
}

func (r *remoteAPI) RemoveScanlineListener(listener ScanlineListener) {
	r.removeListener(listener, scanline, 16)
}

func (r *remoteAPI) AddScanlineCycleListener(listener ScanlineCycleListener, scanline, cycle int) {
	if listener != nil {
		point := newScanlineCyclePoint(listener, scanline, cycle)
		r.sendListener(r.addListenerObject(listener, scanlineCycle, point), scanlineCycle, point)
	}
}

func (r *remoteAPI) RemoveScanlineCycleListener(listener ScanlineCycleListener) {
	r.removeListener(listener, scanlineCycle, 18)
}

func (r *remoteAPI) AddSpriteZeroListener(listener SpriteZeroListener) {
	r.addListener(listener, spriteZero)
}

func (r *remoteAPI) RemoveSpriteZeroListener(listener SpriteZeroListener) {
	r.removeListener(listener, spriteZero, 20)
}

func (r *remoteAPI) AddStatusListener(listener StatusListener) {
	r.addListener(listener, status)
}

func (r *remoteAPI) RemoveStatusListener(listener StatusListener) {
	r.removeListener(listener, status, 22)
}

func (r *remoteAPI) GetPixels(pixels []int) {
	r.stream.writeByte(119)
	r.stream.flush()
	r.stream.readIntArray(pixels)
}
