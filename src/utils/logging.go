package utils

import (
	"log/slog"
	"strings"
)

const ERROR_HEDEAR = "Errorr!!"

func PrintParserErrors(logger *slog.Logger, errors []string) {
	var builder strings.Builder
	builder.WriteString(ERROR_HEDEAR)
	builder.WriteString("Parser errors:\n")
	for _, msg := range errors {
		builder.WriteString("\t" + msg + "\n")
	}

	logger.Error(builder.String(), "type", "parser")
}
