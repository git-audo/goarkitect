package should

import (
	"bytes"
	"fmt"
	"goarkitect/internal/arch/rule"
	"os"
	"path/filepath"

	"golang.org/x/exp/slices"
)

func HaveContentMatching(want []byte, opts ...Option) *Expression {
	return NewExpression(
		func(rb rule.Builder, filePath string) bool {
			data, err := os.ReadFile(filePath)
			if err != nil {
				rb.AddError(err)
				return true
			}

			match := "SINGLE"
			separator := []byte("\n")
			for _, opt := range opts {
				switch opt.(type) {
				case IgnoreNewLinesAtTheEndOfFile:
					data = bytes.TrimRight(data, "\n")
					want = bytes.TrimRight(want, "\n")
				case IgnoreCase:
					data = bytes.ToLower(data)
					want = bytes.ToLower(want)
				case MatchSingleLines:
					match = "MULTIPLE"
					if sep := opt.(MatchSingleLines).Separator; sep != "" {
						separator = []byte(sep)
					}
				}
			}

			if match == "SINGLE" {
				return slices.Compare(data, want) != 0
			}

			linesData := bytes.Split(data, separator)
			for _, ld := range linesData {
				if slices.Compare(ld, want) != 0 {
					return true
				}
			}

			return false
		},
		func(filePath string, options options) rule.Violation {
			format := "file '%s' does not have content matching '%s'"

			if options.matchSingleLines {
				format = "file '%s' does not have all lines matching '%s'"
			}

			if options.negated {
				format = "file '%s' does have content matching '%s'"
			}

			if options.negated && options.matchSingleLines {
				format = "file '%s' does have all lines matching '%s'"
			}

			return rule.NewViolation(
				fmt.Sprintf(format, filepath.Base(filePath), want),
			)
		},
		opts...,
	)
}