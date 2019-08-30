package nintaco

// StatusListener is the listener interface for status change messages.
type StatusListener interface {
	statusChanged(message string)
}
