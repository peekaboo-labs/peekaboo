package text

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/mickep76/color"
)

type Table struct {
	Headers []string
	Rows    [][]string
}

type Tables []*Table

func (ts Tables) PrintTable(output io.Writer) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, strings.Join(ts[0].Headers, "\t"))
	for _, t := range ts {
		for _, r := range t.Rows {
			fmt.Fprintln(w, strings.Join(r, "\t"))
		}
	}

	w.Flush()
}

func (ts Tables) PrintVertTable(output io.Writer) {
	if len(ts) < 1 {
		return
	}

	w := tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
	for _, t := range ts {
		for _, r := range t.Rows {
			for i, c := range r {
				if i == 0 {
					fmt.Fprintf(w, "%s%s\t: %s%s%s\n", color.LightCyan, t.Headers[i], color.LightYellow, c, color.Reset)
				} else {
					fmt.Fprintf(w, "%s%s\t: %s%s%s\n", color.Cyan, t.Headers[i], color.Yellow, c, color.Reset)
				}
			}
			fmt.Fprintln(w)
		}
	}

	w.Flush()
}

func (ts Tables) PrintCSV(output io.Writer) {
	if len(ts) < 1 {
		return
	}

	w := csv.NewWriter(output)
	w.Write(ts[0].Headers)
	for _, t := range ts {
		w.WriteAll(t.Rows)
	}
}