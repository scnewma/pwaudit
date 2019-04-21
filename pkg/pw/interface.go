package pw

// Password wraps a password and a description.
//
// This allows Loader's to provide a human-readable
// description for the loaded passwords.
type Password struct {
	Plaintext   string
	Description string
}

// CheckedPassword represents a Password that has
// been checked to see if it is compromised (which
// can be retrieved via the Compromised field).
type CheckedPassword struct {
	Password

	Compromised bool
}

// A Checker checks if a password is compromised.
//
// CheckPassword should embed the passed-in password
// in a CheckedPassword wrapper.
type Checker interface {
	Check(password Password) (CheckedPassword, error)
}

// A Loader loads a set of passwords.
//
// Load should return a read-only channel
// in which discovered plaintext passwords will be sent.
// The channel will be closed when the passwords are
// exhausted or an error in encountered.
//
// Error can be called to determine if an error was encountered.
type Loader interface {
	Load() <-chan Password
	Error() error
}

// A Printer prints a slice of CheckedPasswords
type Printer interface {
	Print([]CheckedPassword)
}
