package plaintext

import (
	"bufio"
	"io"

	"github.com/scnewma/pwaudit/pkg/pw"
)

// Loader loads passwords from a
// line-delimited reader source.
type Loader struct {
	scanner *bufio.Scanner

	passwords chan pw.Password
	err       error
}

// NewLoader returns a new Loader from r.
func NewLoader(r io.Reader) *Loader {
	l := &Loader{
		scanner:   bufio.NewScanner(r),
		passwords: make(chan pw.Password),
	}
	go l.run()
	return l
}

// Load fulfills the pw.Loader interface
func (l *Loader) Load() <-chan pw.Password {
	return l.passwords
}

// Error fulfills the pw.Loader interface
func (l *Loader) Error() error {
	return l.err
}

func (l *Loader) run() {
	for l.scanner.Scan() {
		password := l.scanner.Text()
		l.passwords <- pw.Password{Plaintext: password, Description: password}
	}
	close(l.passwords)
	if err := l.scanner.Err(); err != nil {
		l.err = err
	}
}
