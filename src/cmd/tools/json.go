package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (p *Parser) generateJson() error {
	if err := os.MkdirAll(p.config.OutputDir, 0755); err != nil {
		return err
	}

	// Generate documentation for types
	for _, t := range p.types {
		filename := filepath.Join(p.config.OutputDir, fmt.Sprintf("type_%s.json", strings.ToLower(t.Name)))
		err := writeStructToFile(filename, t)
		if err != nil {
			return err
		}
	}

	// Generate documentation for libraries
	for _, lib := range p.libraries {
		filename := filepath.Join(p.config.OutputDir, fmt.Sprintf("library_%s.json", strings.ToLower(lib.Name)))
		err := writeStructToFile(filename, lib)
		if err != nil {
			return err
		}
	}

	// Generate documentation for builtins
	filename := filepath.Join(p.config.OutputDir, "builtins.json")
	err := writeStructToFile(filename, p.builtins)
	if err != nil {
		return err
	}

	return nil
}

func writeStructToFile(filename string, v any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(v); err != nil {
		return err
	}
	return nil
}
