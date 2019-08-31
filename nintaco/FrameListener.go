package nintaco

// FrameListener is the listener interface for receiving render events.
type FrameListener interface {
	FrameRendered()
}

// FrameFunc ...
type FrameFunc func()

// NewFrameFunc ...
func NewFrameFunc(listener func()) *FrameFunc {
	f := FrameFunc(listener)
	return &f
}

// FrameRendered ...
func (f *FrameFunc) FrameRendered() {
	(*f)()
}
