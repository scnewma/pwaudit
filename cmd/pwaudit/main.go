package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/scnewma/pwaudit/pkg/haveibeenpwned"
	"github.com/scnewma/pwaudit/pkg/pw"
	"github.com/scnewma/pwaudit/pkg/version"
)

const (
	concurrencyLevel = 5
)

func main() {
	args := os.Args[1:]

	// print version
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			fmt.Println(version.Print())
			os.Exit(0)
		}
	}

	if err := run(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flags := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flags.Usage = func() {
		printHelp(os.Stderr)
	}
	inputFlags := NewInputFlags()
	inputFlags.AddFlags(flags)
	outputFlags := NewOutputFlags()
	outputFlags.AddFlags(flags)
	if err := flags.Parse(args); err != nil {
		return err
	}

	loader, err := inputFlags.ToLoader()
	if err != nil {
		return err
	}

	printer, err := outputFlags.ToPrinter()
	if err != nil {
		return err
	}

	checker := &haveibeenpwned.PasswordChecker{}

	passwords := loader.Load()
	results := make(chan pw.CheckedPassword)
	var wg sync.WaitGroup
	for i := 0; i < concurrencyLevel; i++ {
		wg.Add(1)
		go func() {
			worker(checker, passwords, results)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var checked []pw.CheckedPassword
	for p := range results {
		checked = append(checked, p)
	}

	// check for error loading passwords
	if err := loader.Error(); err != nil {
		return err
	}

	printer.Print(checked)

	return nil
}

func worker(checker pw.Checker, passwords <-chan pw.Password, results chan<- pw.CheckedPassword) {
	for pass := range passwords {
		checked, err := checker.Check(pass)
		if err != nil {
			log.Fatal(err)
		}

		results <- checked
	}
}

func printHelp(w io.Writer) {
	helpText := `
Usage: pwaudit [options]

  Check the provided passwords to see if they have been compromised.

Input Options:

  --lastpass-csv=path       Path to a LastPass exported CSV file. You can export
                            this file in the browser extension via
                            More Options > Advanced > Export

  --passwords=path|-        Path to a line-delimited password dump or '-'. If '-' 
                            is provided then the passwords are read from stdin.

Output Options:

  --show-passwords          Print the plaintext passwords that were checked as well
                            as the password description. The description and the
                            password will be the same for some inputs.

  --show-all                Print both compromised and non-compromised passwords
                            to the screen.

Other Options:

  -v,--version              Print the version.

`

	fmt.Fprintf(w, strings.TrimSpace(helpText))
}
