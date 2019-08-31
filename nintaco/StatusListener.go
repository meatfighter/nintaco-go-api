package nintaco

// StatusListener is the listener interface for status change messages.
type StatusListener interface {
	StatusChanged(message string)
}

// StatusFunc ...
type StatusFunc func(string)

// StatusChanged ...
func (f StatusFunc) StatusChanged(message string) {
	f(message)
}
