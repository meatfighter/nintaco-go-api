package nintaco

// StatusListener is the listener interface for status change messages.
type StatusListener interface {
	StatusChanged(message string)
}

// StatusFunc ...
type StatusFunc func(string)

// NewStatusFunc ...
func NewStatusFunc(listener func(string)) *StatusFunc {
	f := StatusFunc(listener)
	return &f
}

// StatusChanged ...
func (f *StatusFunc) StatusChanged(message string) {
	(*f)(message)
}
