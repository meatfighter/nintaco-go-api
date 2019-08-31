package nintaco

// SpriteZeroListener is the listener interface for sprite zero hit events.
type SpriteZeroListener interface {
	SpriteZeroHit(scanline, scanlineCycle int)
}

// SpriteZeroFunc ...
type SpriteZeroFunc func(int, int)

// SpriteZeroHit ...
func (f SpriteZeroFunc) SpriteZeroHit(scanline, scanlineCycle int) {
	f(scanline, scanlineCycle)
}
