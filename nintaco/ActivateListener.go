package nintaco

// ActivateListener is the listener interface for API enabled events.
type ActivateListener interface {
	APIEnabled()
}

// ActivateFunc ...
type ActivateFunc func()

// APIEnabled ...
func (f ActivateFunc) APIEnabled() {
	f()
}
