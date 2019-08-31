package nintaco

// ControllersListener is the listener interface for receiving controller events.
type ControllersListener interface {
	ControllersProbed()
}

// ControllersFunc ...
type ControllersFunc func()

// NewControllersFunc ...
func NewControllersFunc(listener func()) *ControllersFunc {
	f := ControllersFunc(listener)
	return &f
}

// ControllersProbed ...
func (f *ControllersFunc) ControllersProbed() {
	(*f)()
}
