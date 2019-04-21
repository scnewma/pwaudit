package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/scnewma/pwaudit/pkg/lastpass"
	"github.com/scnewma/pwaudit/pkg/pw"
)

type LastPassInputFlags struct {
	CSVFile *string // location of csv file
}

func (f *LastPassInputFlags) AddFlags(flags *flag.FlagSet) {
	flags.StringVar(f.CSVFile, "lastpass-csv", "", "Location of LastPass CSV export")
}

func (f *LastPassInputFlags) AllowedSources() []string {
	return []string{"lastpass-csv"}
}

func (f *LastPassInputFlags) ToLoader() (pw.Loader, error) {
	if f.CSVFile == nil || len(*f.CSVFile) == 0 {
		return nil, NoCompatibleLoaderError{AllowedSources: f.AllowedSources()}
	}

	csvFile, err := ioutil.ReadFile(*f.CSVFile)
	if err != nil {
		return nil, fmt.Errorf("error reading --lastpass-csv %s, %v\n", *f.CSVFile, err)
	}

	scanner := lastpass.NewCSVScanner(bytes.NewReader(csvFile))
	loader := lastpass.NewCSVLoader(scanner)
	return loader, nil
}
