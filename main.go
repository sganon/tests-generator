package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/jpillora/opts"
	"gopkg.in/yaml.v2"
)

type Config struct {
	SpecsFile   string `opts:"help=specify where is the tests spec file"`
	Pkg         string `opts:"help=specify the pkg name"`
	HandlerFunc string `opts:"help=which function returns the handler to test"`
}

type TemplateData struct {
	*Spec
	*Config
}

func main() {
	c := Config{}
	opts.Parse(&c)

	b, err := ioutil.ReadFile(c.SpecsFile)
	if err != nil {
		log.Fatalf("error reading spec file: %v", err)
	}

	spec := Spec{}
	err = yaml.Unmarshal(b, &spec)
	if err != nil {
		log.Fatalf("error parsing specs: %v", err)
	}

	td := TemplateData{
		Spec:   &spec,
		Config: &c,
	}

	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		log.Fatalf("error parsing template file: %v", err)
	}

	f, err := os.Create(fmt.Sprintf("%s_test.go", td.Pkg))
	if err != nil {
		log.Fatalf("error creating test file: %v", err)
	}

	err = t.Execute(f, td)
	if err != nil {
		log.Fatalf("error generating test: %v", err)
	}
}

var tmpl = `//go:generate tests-generator --specs-file {{.SpecsFile}} --pkg {{.Pkg}} --handler-func {{.HandlerFunc}}
package {{.Pkg}}

// This file was generated via test-generator

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/json"
)

var h http.Handler = {{.HandlerFunc}}()

{{range .Suite}}
func Test{{.Name}}(t *testing.T) {
	var reqBody io.Reader
	if ` + "`{{.Body}}`" + ` != "" {
		reqBody =  bytes.NewBuffer([]byte(` + "`{{.Body}}`" + `))
	}

	respRec := httptest.NewRecorder()
	req, err := http.NewRequest("{{.Method}}", "{{.Path}}{{.Query}}", reqBody)
	if err != nil {
		t.Fatal("Creating '{{.Method}} {{.Path}}' request failed!")
	}

	h.ServeHTTP(respRec, req)

	assert.Equal(t, respRec.Code, {{.Response.Status}}, "{{.Name}}: unexpected response code")

	body := new(bytes.Buffer)
	ref := new(bytes.Buffer)
	m := minify.New()
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	if err := m.Minify("application/json", body, respRec.Body); err != nil {
		panic(err)
	}
	if err := m.Minify("application/json", ref, bytes.NewBuffer([]byte(` + "`{{.Response.Body}}`" + `))); err != nil {
		panic(err)
	}



	assert.Equal(t, ref, body,"{{.Name}}: response body is not matching")
}
{{end}}
`
