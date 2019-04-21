package printers

import (
	"io"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/scnewma/pwaudit/pkg/pw"
)

// PrintOptions defines the configurable options that
// each printer must accomodate.
type PrintOptions struct {
	ShowPasswords bool
	ShowAll       bool
}

type Table struct {
	w    *tablewriter.Table
	opts PrintOptions
}

func NewTablePrinter(w io.Writer, opts PrintOptions) *Table {
	tw := tablewriter.NewWriter(w)
	tw.SetBorder(false)
	tw.SetHeaderLine(false)
	tw.SetColumnSeparator("")

	return &Table{w: tw, opts: opts}
}

func (p *Table) Print(passwords []pw.CheckedPassword) {
	p.w.SetHeader(p.header())
	p.w.AppendBulk(p.rows(passwords))
	p.w.Render()
}

func (p *Table) header() []string {
	header := []string{"Description"}
	if p.opts.ShowPasswords {
		header = append(header, "Password")
	}
	// do not show column if only showing compromised passwords
	if p.opts.ShowAll {
		header = append(header, "Compromised")
	}
	return header
}

func (p *Table) rows(passwords []pw.CheckedPassword) [][]string {
	var rows [][]string
	for _, pass := range passwords {
		if !p.opts.ShowAll && !pass.Compromised {
			continue
		}

		cols := []string{pass.Description}
		if p.opts.ShowPasswords {
			cols = append(cols, pass.Plaintext)
		}
		if p.opts.ShowAll {
			cols = append(cols, strconv.FormatBool(pass.Compromised))
		}

		rows = append(rows, cols)
	}
	return rows
}
