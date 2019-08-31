package nintaco

// DeactivateListener is the listener interface for API disabled events.
type DeactivateListener interface {
	APIDisabled()
}

// DeactivateFunc ...
type DeactivateFunc func()

// APIDisabled ...
func (f DeactivateFunc) APIDisabled() {
	f()
}
