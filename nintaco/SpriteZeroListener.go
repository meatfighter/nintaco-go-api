package nintaco

// SpriteZeroListener is the listener interface for sprite zero hit events.
type SpriteZeroListener interface {
	SpriteZeroHit(scanline, scanlineCycle int)
}

// SpriteZeroFunc is a SpriteZeroListener.
type SpriteZeroFunc func(int, int)

// NewSpriteZeroFunc casts a function into a SpriteZeroListener.
func NewSpriteZeroFunc(listener func(int, int)) *SpriteZeroFunc {
	f := SpriteZeroFunc(listener)
	return &f
}

// SpriteZeroHit is the method that delegates the call to the listener function.
func (f *SpriteZeroFunc) SpriteZeroHit(scanline, scanlineCycle int) {
	(*f)(scanline, scanlineCycle)
}
