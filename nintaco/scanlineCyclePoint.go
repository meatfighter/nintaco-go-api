package nintaco

type scanlineCyclePoint struct {
	listener      ScanlineCycleListener
	scanline      int
	scanlineCycle int
}

func newScanlineCyclePoint(listener ScanlineCycleListener, scanline, scanlineCycle int) *scanlineCyclePoint {
	return &scanlineCyclePoint{
		listener:      listener,
		scanline:      scanline,
		scanlineCycle: scanlineCycle,
	}
}
