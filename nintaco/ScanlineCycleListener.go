package nintaco

// ScanlineCycleListener is the listener interface for receiving scanline cycle events.
type ScanlineCycleListener interface {
	CyclePerformed(scanline, scanlineCycle, address int, rendering bool)
}

// ScanlineCycleFunc is a ScanlineCycleListener.
type ScanlineCycleFunc func(int, int, int, bool)

// NewScanlineCycleFunc casts a function into a ScanlineCycleListener.
func NewScanlineCycleFunc(listener func(int, int, int, bool)) *ScanlineCycleFunc {
	f := ScanlineCycleFunc(listener)
	return &f
}

// CyclePerformed is the method that delegates the call to the listener function.
func (f *ScanlineCycleFunc) CyclePerformed(scanline, scanlineCycle, address int, rendering bool) {
	(*f)(scanline, scanlineCycle, address, rendering)
}
