package utils

import "io"

const ERROR_HEDEAR = "Errorr!!"

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ERROR_HEDEAR)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
