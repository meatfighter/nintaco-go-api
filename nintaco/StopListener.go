package nintaco

// StopListener is the listener interface for stop events.
type StopListener interface {
	Dispose()
}

// StopFunc ...
type StopFunc func()

// NewStopFunc ...
func NewStopFunc(listener func()) *StopFunc {
	f := StopFunc(listener)
	return &f
}

// Dispose ...
func (f *StopFunc) Dispose() {
	(*f)()
}
