package nintaco

// SpriteZeroListener is the listener interface for sprite zero hit events.
type SpriteZeroListener interface {
	SpriteZeroHit(scanline, scanlineCycle int)
}
