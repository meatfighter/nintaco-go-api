package nintaco

// ScanlineListener is the listener interface for receiving scanline render events.
type ScanlineListener interface {
	ScanlineRendered(scanline int)
}
