package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/sevenreup/duwa/src/duwa"
	"github.com/sevenreup/duwa/src/object"
	"github.com/sevenreup/duwa/src/runtime/native"
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
	console := native.NewConsole()
	duwa := duwa.New(object.New(logger, console))
	value := duwa.RunFile(file)
	if value != nil {
		if object.IsError(value) {
			log.Fatal(value.String())
		}
	}
}
