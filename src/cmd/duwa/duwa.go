package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/sevenreup/duwa/src/duwa"
	"github.com/sevenreup/duwa/src/object"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "f", "", "Source file")
}

func main() {
	flag.Parse()

	if file == "" {
		log.Fatal("Please provide a file to run")
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	duwa := duwa.New(object.New(logger))
	duwa.RunFile(file)
}
