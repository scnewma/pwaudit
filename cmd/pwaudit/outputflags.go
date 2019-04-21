package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scnewma/pwaudit/pkg/printers"
	"github.com/scnewma/pwaudit/pkg/pw"
)

// OutputFlags represents the flags that control
// printing the results.
type OutputFlags struct {
	// print the plaintext passwords
	ShowPasswords *bool

	// print all passwords; not just compromised passwords
	ShowAll *bool
}

func NewOutputFlags() *OutputFlags {
	return &OutputFlags{
		ShowPasswords: new(bool),
		ShowAll:       new(bool),
	}
}

func (f *OutputFlags) AddFlags(flags *flag.FlagSet) {
	flags.BoolVar(f.ShowPasswords, "show-passwords", false, "Print passwords")
	flags.BoolVar(f.ShowAll, "show-all", false, "Print non-compromised passwords")
}

func (f *OutputFlags) ToPrinter() (pw.Printer, error) {
	if f == nil {
		return nil, fmt.Errorf("no output flags found")
	}

	showPasswords := false
	if f.ShowPasswords != nil {
		showPasswords = *f.ShowPasswords
	}

	showAll := false
	if f.ShowAll != nil {
		showAll = *f.ShowAll
	}

	return printers.NewTablePrinter(os.Stdout, printers.PrintOptions{
		ShowPasswords: showPasswords,
		ShowAll:       showAll,
	}), nil
}
