package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/scnewma/pwaudit/pkg/plaintext"
	"github.com/scnewma/pwaudit/pkg/pw"
)

type LineDelimitedInputFlags struct {
	File *string // location of file
}

func (f *LineDelimitedInputFlags) AddFlags(flags *flag.FlagSet) {
	flags.StringVar(f.File, "passwords", "", "Location of line-delimited password dump")
}

func (f *LineDelimitedInputFlags) AllowedSources() []string {
	return []string{"passwords"}
}

func (f *LineDelimitedInputFlags) ToLoader() (pw.Loader, error) {
	if f.File == nil || len(*f.File) == 0 {
		return nil, NoCompatibleLoaderError{AllowedSources: f.AllowedSources()}
	}

	if (*f.File) == "-" {
		return plaintext.NewLoader(os.Stdin), nil
	}

	file, err := ioutil.ReadFile(*f.File)
	if err != nil {
		return nil, fmt.Errorf("error reading --passwords %s, %v\n", *f.File, err)
	}

	return plaintext.NewLoader(bytes.NewReader(file)), nil
}
