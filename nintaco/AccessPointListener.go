package nintaco

// AccessPointListener is the listener interface for access point events.
type AccessPointListener interface {
	accessPointHit(accessPointType, address, value int) int
}
