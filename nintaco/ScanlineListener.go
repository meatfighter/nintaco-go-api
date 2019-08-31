package nintaco

// ScanlineListener is the listener interface for receiving scanline render events.
type ScanlineListener interface {
	ScanlineRendered(scanline int)
}

// ScanlineFunc is a ScanlineListener.
type ScanlineFunc func(int)

// NewScanlineFunc casts a function into a ScanlineListener.
func NewScanlineFunc(listener func(int)) *ScanlineFunc {
	f := ScanlineFunc(listener)
	return &f
}

// ScanlineRendered is the method that delegates the call to the listener function.
func (f *ScanlineFunc) ScanlineRendered(scanline int) {
	(*f)(scanline)
}
