package main

import (
	"bytes"

	"io/ioutil"
	"log"

	"text/template"

	ts "github.com/james-maloney/termstyle"
)

func main() {
	t := template.Must(template.New("init").Parse(tmpl))

	buf := &bytes.Buffer{}
	data := struct {
		FG []string
		BG []string
		B  string
		U  string
		C  string
	}{ts.FG, ts.BG, ts.B, ts.U, ts.C}

	if err := t.Execute(buf, data); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("../generated.go", buf.Bytes(), 0644); err != nil {
		log.Fatal(err)
	}
}

var tmpl = `// Generated code, DO NOT EDIT.
package termstyle

var (
	B = "{{.B}}"
	U = "{{.U}}"
	C = "{{.C}}"
	FG = []string{
	{{ range .FG }}
		"{{.}}",
	{{ end }}
	}
	BG = []string{
	{{ range .BG }}
		"{{.}}",
	{{ end }}
	}
)
`
