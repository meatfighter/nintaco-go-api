package nintaco

// StopListener is the listener interface for stop events.
type StopListener interface {
	Dispose()
}

// StopFunc is a StopListener.
type StopFunc func()

// NewStopFunc casts a function into a StopListener.
func NewStopFunc(listener func()) *StopFunc {
	f := StopFunc(listener)
	return &f
}

// Dispose is the method that delegates the call to the listener function.
func (f *StopFunc) Dispose() {
	(*f)()
}
