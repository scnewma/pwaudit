package lastpass

import (
	"io"
	"net/url"
	"strconv"

	"github.com/smartystreets/scanners/csv"
)

// CSVScanner provides a clean interface
// to scan LastPass exported CSV files.
type CSVScanner struct {
	*csv.Scanner
}

// NewCSVScanner will construct a CSVScanner
// that wraps the provided io.Reader.
func NewCSVScanner(r io.Reader) *CSVScanner {
	inner := csv.NewScanner(r)
	inner.Scan() // skip header
	return &CSVScanner{Scanner: inner}
}

// Record will return the most recently
// scanned Entry
func (s *CSVScanner) Record() Entry {
	fields := s.Scanner.Record()

	u, _ := url.Parse(fields[0])
	fav, _ := strconv.Atoi(fields[5])
	return Entry{
		URL:      *u,
		Username: fields[1],
		Password: fields[2],
		Extra:    fields[3],
		Name:     fields[4],
		Fav:      fav,
	}
}
