package nintaco

// AccessPointListener is the listener interface for access point events.
type AccessPointListener interface {
	AccessPointHit(accessPointType, address, value int) int
}
