package main

type Spec struct {
	Suite []Test `yaml:"suite"`
}

type Test struct {
	Name     string       `yaml:"name"`
	Method   string       `yaml:"method"`
	Path     string       `yaml:"path"`
	Query    string       `yaml:"query"`
	Body     string       `yaml:"body"`
	Response TestResponse `yaml:"response"`
}

type TestResponse struct {
	Status int    `yaml:"status"`
	Body   string `yaml:"body"`
}
