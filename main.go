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

type TemplateData struct {
	*Spec
	HandlerFunc string
	Pkg         string
}

func main() {
	type config struct {
		SpecsFile string `opts:"help=specify where is the tests spec file"`
	}
	c := config{}
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
		Spec:        &spec,
		Pkg:         "example",
		HandlerFunc: "NewRouter",
	}

	t, err := template.ParseFiles("./tests.template")
	if err != nil {
		log.Fatalf("error parsing template file: %v", err)
	}

	f, err := os.Create(fmt.Sprintf("%s/%s_test.go", td.Pkg, td.Pkg))
	if err != nil {
		log.Fatalf("error creating test file: %v", err)
	}

	err = t.Execute(f, td)
	if err != nil {
		log.Fatalf("error generating test: %v", err)
	}
}
