package nintaco

// StopListener is the listener interface for stop events.
type StopListener interface {
	Dispose()
}

// StopFunc ...
type StopFunc func()

// Dispose ...
func (f StopFunc) Dispose() {
	f()
}
