package nintaco

// ScanlineCycleListener is the listener interface for receiving scanline cycle events.
type ScanlineCycleListener interface {
	CyclePerformed(scanline, scanlineCycle, address int, rendering bool)
}
