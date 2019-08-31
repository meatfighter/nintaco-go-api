package nintaco

// A set of constants used as an argument in API.AddAccessPointListener to indicate
// the type of access points that trigger the associated AccessPointListener.
const (
	// Occurs just before a CPU Memory read, providing an opportunity to skip
	// the read and to substitute a value in its place.
	AccessPointTypePreRead = 0

	// Occurs just after a CPU Memory read, providing an opportunity to substitute
	// the read value.
	AccessPointTypePostRead = 1

	// Occurs just before a CPU Memory write, providing an opportunity to substitute
	// the value to be written.
	AccessPointTypePreWrite = 2

	// Occurs just after a CPU Memory write, too late to affect the behavior of
	// the write.
	AccessPointTypePostWrite = 3

	// Occurs just before an instruction is executed.
	AccessPointTypePreExecute = 4

	// Occurs just after an instruction is executed.
	AccessPointTypePostExecute = 5
)
