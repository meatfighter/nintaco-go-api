package nintaco

// FrameListener is the listener interface for receiving render events.
type FrameListener interface {
	FrameRendered()
}

// FrameFunc ...
type FrameFunc func()

// FrameRendered ...
func (f FrameFunc) FrameRendered() {
	f()
}
