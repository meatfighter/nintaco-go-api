package nintaco

// AccessPointListener is the listener interface for access point events.
type AccessPointListener interface {
	AccessPointHit(accessPointType, address, value int) int
}

// AccessPointFunc ...
type AccessPointFunc func(int, int, int) int

// NewAccessPointFunc ...
func NewAccessPointFunc(listener func(int, int, int) int) *AccessPointFunc {
	f := AccessPointFunc(listener)
	return &f
}

// AccessPointHit ...
func (f *AccessPointFunc) AccessPointHit(accessPointType, address, value int) int {
	return (*f)(accessPointType, address, value)
}
