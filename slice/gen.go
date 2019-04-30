// +build ignore

// This program generate slice.go and slice_test.go
package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var types = []string{
	"bool",
	// As is said in: https://golang.org/pkg/builtin/#byte, we can skip this type
	// because is an alias
	//"byte",
	"complex128",
	"complex64",
	"error",
	"float32",
	"float64",
	"int",
	"int16",
	"int32",
	"int64",
	"int8",
	// As is said in: https://golang.org/pkg/builtin/#rune, we can skip this type
	// because is an alias
	//"rune",
	"string",
	"uint",
	"uint16",
	"uint32",
	"uint64",
	"uint8",
	"uintptr",
	//"interface{}",
}

var (
	// Output file + .go; Input file without .tpl
	relations = map[string]string{
		"slices":      "code",
		"slices_test": "code_test",
	}

	tplFuncs = template.FuncMap{
		"Title": strings.Title,
	}
)

func main() {
	template := template.Must(template.New("").Funcs(tplFuncs).ParseGlob("*.tpl"))

	for output, input := range relations {
		generateCode(output, input, template)
	}
}

func generateCode(output, input string, tpl *template.Template) {
	f, err := os.Create(output + ".go")
	die(err)
	defer f.Close()

	err = tpl.ExecuteTemplate(f, input+".tpl", struct {
		Timestamp time.Time
		Types     []string
	}{
		Timestamp: time.Now(),
		Types:     types,
	})
	die(err)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
