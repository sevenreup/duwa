package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
)

type ArgumentData struct {
	Name string
	Type string
}

type FunctionInfo struct {
	Name      string
	Arguments []ArgumentData
	RetunType string
	Doc       string
}

type StructInfo struct {
	Name         string
	Alternatives []string
	Doc          string
	Methods      []FunctionInfo
}

// parseStructLine parses the first line of struct documentation
// Example: "type=Mawu alternative=String"
func parseStructLine(line string) (string, []string, error) {
	// Extract type name
	typeRegex := regexp.MustCompile(`type=(\w+)`)
	typeMatch := typeRegex.FindStringSubmatch(line)
	if len(typeMatch) < 2 {
		return "", nil, fmt.Errorf("type name not found")
	}
	typeName := typeMatch[1]

	// Extract alternative names
	altRegex := regexp.MustCompile(`alternative=(\w+)`)
	altMatch := altRegex.FindStringSubmatch(line)
	var alternatives []string
	if len(altMatch) >= 2 {
		alternatives = strings.Split(altMatch[1], ",")
	}

	return typeName, alternatives, nil
}

// parseStructDocumentation splits the type information line from the rest of the documentation
func parseStructDocumentation(docText string) (string, string) {
	lines := strings.Split(strings.TrimSpace(docText), "\n")
	if len(lines) == 0 {
		return "", ""
	}

	typeLine := lines[0]
	docLines := lines[1:]
	return typeLine, strings.TrimSpace(strings.Join(docLines, "\n"))
}

// parseMethodLine parses the first line of the documentation that contains method information
func parseMethodLine(line string) (string, []ArgumentData, string, error) {
	// Extract method name
	methodRegex := regexp.MustCompile(`method=(\w+)`)
	methodMatch := methodRegex.FindStringSubmatch(line)
	if len(methodMatch) < 2 {
		return "", nil, "", fmt.Errorf("method name not found")
	}
	methodName := methodMatch[1]

	// Extract arguments
	argsRegex := regexp.MustCompile(`args=\[(.*?)\]`)
	argsMatch := argsRegex.FindStringSubmatch(line)
	var args []ArgumentData
	if len(argsMatch) >= 2 && argsMatch[1] != "" {
		argParts := strings.Split(argsMatch[1], ",")
		for _, arg := range argParts {
			// Parse each argument in format "type{name}"
			argRegex := regexp.MustCompile(`(\w+)\{(\w+)\}`)
			argMatch := argRegex.FindStringSubmatch(strings.TrimSpace(arg))
			if len(argMatch) >= 3 {
				args = append(args, ArgumentData{
					Type: argMatch[1],
					Name: argMatch[2],
				})
			}
		}
	}

	// Extract return type
	returnRegex := regexp.MustCompile(`return=\{(.*?)\}`)
	returnMatch := returnRegex.FindStringSubmatch(line)
	if len(returnMatch) < 2 {
		return "", nil, "", fmt.Errorf("return type not found")
	}
	returnType := returnMatch[1]

	return methodName, args, returnType, nil
}

// parseMethodDocumentation splits the method information line from the rest of the documentation
func parseMethodDocumentation(docText string) (string, string) {
	lines := strings.Split(strings.TrimSpace(docText), "\n")
	if len(lines) == 0 {
		return "", ""
	}

	methodLine := lines[0]
	docLines := lines[1:]
	return methodLine, strings.TrimSpace(strings.Join(docLines, "\n"))
}

// AnalyzeFile finds all structs and their methods in the file
func AnalyzeFile(filename string) ([]StructInfo, error) {
	// Create the file set
	fset := token.NewFileSet()

	// Parse the Go source file
	astFile, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %v", err)
	}

	// Create a new package to analyze
	pkg := &ast.Package{
		Name:  astFile.Name.Name,
		Files: map[string]*ast.File{filename: astFile},
	}

	// Use the doc package to extract documentation
	docPkg := doc.New(pkg, ".", doc.AllDecls)

	var structs []StructInfo

	// Process each type in the package
	for _, t := range docPkg.Types {
		// Parse struct documentation
		if t.Doc != "" {
			typeLine, documentation := parseStructDocumentation(t.Doc)
			typeName, alternatives, err := parseStructLine(typeLine)
			if err != nil {
				continue
			}

			structInfo := StructInfo{
				Name:         typeName,
				Alternatives: alternatives,
				Doc:          documentation,
				Methods:      []FunctionInfo{},
			}

			// Process methods
			for _, method := range t.Methods {
				if method.Doc != "" {
					methodLine, documentation := parseMethodDocumentation(method.Doc)
					name, args, returnType, err := parseMethodLine(methodLine)
					if err == nil {
						functionInfo := FunctionInfo{
							Name:      name,
							Arguments: args,
							RetunType: returnType,
							Doc:       documentation,
						}
						structInfo.Methods = append(structInfo.Methods, functionInfo)
					}
				}
			}

			structs = append(structs, structInfo)
		}
	}

	return structs, nil
}

func main() {
	filename := "./src/object/string.go"

	structs, err := AnalyzeFile(filename)
	if err != nil {
		fmt.Printf("Error analyzing file: %v\n", err)
		return
	}

	// Print the results
	for _, structInfo := range structs {
		fmt.Printf("\nStruct: %s\n", structInfo.Name)
		fmt.Printf("Alternatives: %v\n", structInfo.Alternatives)
		fmt.Printf("Documentation:\n%s\n", structInfo.Doc)

		fmt.Printf("\nMethods:\n")
		for _, method := range structInfo.Methods {
			fmt.Printf("\n  Method: %s\n", method.Name)
			fmt.Printf("  Arguments:\n")
			for _, arg := range method.Arguments {
				fmt.Printf("    - %s: %s\n", arg.Name, arg.Type)
			}
			fmt.Printf("  Return Type: %s\n", method.RetunType)
			fmt.Printf("  Documentation:\n%s\n", method.Doc)
		}
	}
}
