package nintaco

// ControllersListener is the listener interface for receiving controller events.
type ControllersListener interface {
	ControllersProbed()
}

// ControllersFunc is a ControllersListener.
type ControllersFunc func()

// NewControllersFunc casts a function into a ControllersListener.
func NewControllersFunc(listener func()) *ControllersFunc {
	f := ControllersFunc(listener)
	return &f
}

// ControllersProbed is the method that delegates the call to the listener function.
func (f *ControllersFunc) ControllersProbed() {
	(*f)()
}
