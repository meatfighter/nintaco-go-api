package nintaco

// DeactivateListener is the listener interface for API disabled events.
type DeactivateListener interface {
	APIDisabled()
}

// DeactivateFunc is a DeactivateListener.
type DeactivateFunc func()

// NewDeactivateFunc casts a function into a DeactivateListener.
func NewDeactivateFunc(listener func()) *DeactivateFunc {
	f := DeactivateFunc(listener)
	return &f
}

// APIDisabled is the method that delegates the call to the listener function.
func (f *DeactivateFunc) APIDisabled() {
	(*f)()
}
