package lastpass

import (
	"fmt"

	"github.com/scnewma/pwaudit/pkg/pw"
)

// CSVLoader loads passwords from an exported LastPass CSV file.
type CSVLoader struct {
	scanner *CSVScanner

	passwords chan pw.Password
	err       error
}

// NewCSVLoader returns a CSVLoader bootstrapped
// with the provided CSVScanner
func NewCSVLoader(scanner *CSVScanner) *CSVLoader {
	l := &CSVLoader{
		scanner:   scanner,
		passwords: make(chan pw.Password),
	}
	go l.run()
	return l
}

// Load fulfills the pw.Loader interface
func (l *CSVLoader) Load() <-chan pw.Password {
	return l.passwords
}

// Load fulfills the pw.Loader interface
func (l *CSVLoader) Error() error {
	return l.err
}

func (l *CSVLoader) run() {
	for l.scanner.Scan() {
		record := l.scanner.Record()
		l.passwords <- pw.Password{
			Plaintext:   record.Password,
			Description: fmt.Sprintf("%s - %s", record.URL.String(), record.Username),
		}
	}
	close(l.passwords)
	if err := l.scanner.Error(); err != nil {
		l.err = err
	}
}
