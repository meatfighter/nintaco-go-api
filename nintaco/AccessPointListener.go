package nintaco

// AccessPointListener is the listener interface for access point events.
type AccessPointListener interface {
	AccessPointHit(accessPointType, address, value int) int
}

// AccessPointFunc is an AccessPointListener.
type AccessPointFunc func(int, int, int) int

// NewAccessPointFunc casts a function into an AccessPointListener.
func NewAccessPointFunc(listener func(int, int, int) int) *AccessPointFunc {
	f := AccessPointFunc(listener)
	return &f
}

// AccessPointHit is the method that delegates the call to the listener function.
func (f *AccessPointFunc) AccessPointHit(accessPointType, address, value int) int {
	return (*f)(accessPointType, address, value)
}
