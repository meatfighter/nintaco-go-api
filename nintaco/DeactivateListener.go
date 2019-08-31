package nintaco

// DeactivateListener is the listener interface for API disabled events.
type DeactivateListener interface {
	APIDisabled()
}

// DeactivateFunc ...
type DeactivateFunc func()

// NewDeactivateFunc ...
func NewDeactivateFunc(listener func()) *DeactivateFunc {
	f := DeactivateFunc(listener)
	return &f
}

// APIDisabled ...
func (f *DeactivateFunc) APIDisabled() {
	(*f)()
}
