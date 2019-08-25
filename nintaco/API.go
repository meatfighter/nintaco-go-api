package nintaco

// API provides programmatic control of the emulator at a very granular level.
type API interface {
	AddActivateListener(listener ActivateListener)
	AddDeactivateListener(listener DeactivateListener)

	// TODO AUTOGENERATE THIS INTERFACE
}
