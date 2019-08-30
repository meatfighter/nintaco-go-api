package nintaco

type scanlinePoint struct {
	listener ScanlineListener
	scanline int
}

func newScanlinePoint(listener ScanlineListener, scanline int) *scanlinePoint {
	return &scanlinePoint{
		listener: listener,
		scanline: scanline,
	}
}
