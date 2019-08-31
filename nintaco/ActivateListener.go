package nintaco

// ActivateListener is the listener interface for API enabled events.
type ActivateListener interface {
	APIEnabled()
}

// ActivateFunc ...
type ActivateFunc func()

// NewActivateFunc ...
func NewActivateFunc(listener func()) *ActivateFunc {
	f := ActivateFunc(listener)
	return &f
}

// APIEnabled ...
func (f *ActivateFunc) APIEnabled() {
	(*f)()
}
