package nintaco

// ControllersListener is the listener interface for receiving controller events.
type ControllersListener interface {
	ControllersProbed()
}

// ControllersFunc ...
type ControllersFunc func()

// ControllersProbed ...
func (f ControllersFunc) ControllersProbed() {
	f()
}
