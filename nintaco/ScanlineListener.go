package nintaco

// ScanlineListener is the listener interface for receiving scanline render events.
type ScanlineListener interface {
	ScanlineRendered(scanline int)
}

// ScanlineFunc ...
type ScanlineFunc func(int)

// NewScanlineFunc ...
func NewScanlineFunc(listener func(int)) *ScanlineFunc {
	f := ScanlineFunc(listener)
	return &f
}

// ScanlineRendered ...
func (f *ScanlineFunc) ScanlineRendered(scanline int) {
	(*f)(scanline)
}
