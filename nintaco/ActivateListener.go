package nintaco

// ActivateListener is the listener interface for API enabled events.
type ActivateListener interface {
	APIEnabled()
}

// ActivateFunc is an ActivateListener.
type ActivateFunc func()

// NewActivateFunc casts a function into an ActivateListener.
func NewActivateFunc(listener func()) *ActivateFunc {
	f := ActivateFunc(listener)
	return &f
}

// APIEnabled is the method that delegates the call to the listener function.
func (f *ActivateFunc) APIEnabled() {
	(*f)()
}
