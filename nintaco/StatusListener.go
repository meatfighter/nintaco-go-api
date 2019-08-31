package nintaco

// StatusListener is the listener interface for status change messages.
type StatusListener interface {
	StatusChanged(message string)
}

// StatusFunc is a StatusListener.
type StatusFunc func(string)

// NewStatusFunc casts a function into a StatusListener.
func NewStatusFunc(listener func(string)) *StatusFunc {
	f := StatusFunc(listener)
	return &f
}

// StatusChanged is the method that delegates the call to the listener function.
func (f *StatusFunc) StatusChanged(message string) {
	(*f)(message)
}
