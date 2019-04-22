package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/scnewma/pwaudit/pkg/pw"
)

type NoCompatibleLoaderError struct {
	AllowedSources []string
}

func (e NoCompatibleLoaderError) Error() string {
	return fmt.Sprintf("unable to load passwords, allowed sources are: %s", strings.Join(e.AllowedSources, ","))
}

func IsNoCompatibleLoaderError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(NoCompatibleLoaderError)
	return ok
}

type InputFlags struct {
	LastPassInputFlags      *LastPassInputFlags
	LineDelimitedInputFlags *LineDelimitedInputFlags
}

func NewInputFlags() *InputFlags {
	return &InputFlags{
		LastPassInputFlags: &LastPassInputFlags{
			CSVFile: new(string),
		},
		LineDelimitedInputFlags: &LineDelimitedInputFlags{
			File: new(string),
		},
	}
}

func (f *InputFlags) AddFlags(flags *flag.FlagSet) {
	f.LastPassInputFlags.AddFlags(flags)
	f.LineDelimitedInputFlags.AddFlags(flags)
}

func (f *InputFlags) ToLoader() (pw.Loader, error) {
	if f == nil {
		return nil, NoCompatibleLoaderError{}
	}

	if l, err := f.LineDelimitedInputFlags.ToLoader(); !IsNoCompatibleLoaderError(err) {
		return l, err
	}

	if l, err := f.LastPassInputFlags.ToLoader(); !IsNoCompatibleLoaderError(err) {
		return l, err
	}

	return nil, NoCompatibleLoaderError{AllowedSources: f.AllowedSources()}
}

func (f *InputFlags) AllowedSources() []string {
	sources := f.LastPassInputFlags.AllowedSources()
	return append(sources, f.LineDelimitedInputFlags.AllowedSources()...)
}
