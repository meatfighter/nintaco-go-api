package nintaco

// SpriteZeroListener is the listener interface for sprite zero hit events.
type SpriteZeroListener interface {
	SpriteZeroHit(scanline, scanlineCycle int)
}

// SpriteZeroFunc ...
type SpriteZeroFunc func(int, int)

// NewSpriteZeroFunc ...
func NewSpriteZeroFunc(listener func(int, int)) *SpriteZeroFunc {
	f := SpriteZeroFunc(listener)
	return &f
}

// SpriteZeroHit ...
func (f *SpriteZeroFunc) SpriteZeroHit(scanline, scanlineCycle int) {
	(*f)(scanline, scanlineCycle)
}
