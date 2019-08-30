package nintaco

// SpriteZeroListener is the listener interface for sprite zero hit events.
type SpriteZeroListener interface {
	spriteZeroHit(scanline, scanlineCycle int)
}
