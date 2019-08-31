package nintaco

// ScanlineCycleListener is the listener interface for receiving scanline cycle events.
type ScanlineCycleListener interface {
	CyclePerformed(scanline, scanlineCycle, address int, rendering bool)
}

// ScanlineCycleFunc ...
type ScanlineCycleFunc func(int, int, int, bool)

// CyclePerformed ...
func (f ScanlineCycleFunc) CyclePerformed(scanline, scanlineCycle, address int, rendering bool) {
	f(scanline, scanlineCycle, address, rendering)
}
