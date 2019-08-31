package nintaco

// FrameListener is the listener interface for receiving render events.
type FrameListener interface {
	FrameRendered()
}

// FrameFunc is a FrameListener.
type FrameFunc func()

// NewFrameFunc casts a function into a FrameListener.
func NewFrameFunc(listener func()) *FrameFunc {
	f := FrameFunc(listener)
	return &f
}

// FrameRendered is the method that delegates the call to the listener function.
func (f *FrameFunc) FrameRendered() {
	(*f)()
}
