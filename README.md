# WIP: tests-generator [![CircleCI](https://circleci.com/gh/sganon/tests-generator/tree/master.svg?style=svg)](https://circleci.com/gh/sganon/tests-generator/tree/master)
Generate tests for your routes based on a declaration file (yaml)

## Usage
This projects aims to generate tests for your API's routes.  It tries to follow [table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) philosophy. The goal is to declare your test in a conf file written in yaml and generates it as it quite a repetetive process.

The specs file consist of a list of test defined as follow:
```yaml
suite:
  - name: GetIndex # The name of the test functions => TestGetIndex(t *testing.T)
    method: GET # The route method
    path: / # The route path
    query: '?page=0' # Optional query parameters to pass
    body: '{}' # Optional stringified body to send
    response:
      status: 200 # The expected response status code
      body: `[{...}]` # The expected stringified response
```

To be able to generate test your package you need to have a function returning the `http.Handler` to test.

Assuming this function is called `NewRouter`, is in the `api` package, and your specs file is `api/specs.yaml`, create the file `api/api_test.go` with content:
```golang
//go:generate tests-generator --specs-file ./specs.yaml --pkg api --handler-func NewRouter
package api
```

Then simply go in `api` and run `go generate`.


## Todos
- [x] Better handling of target package (Use of go generate kinda handle that)
- [ ] Pass parameters to handler func
- [ ] Handles other types of body than `application/json`
- [ ] Use openapi parser to define response/request bodies
- [ ] Give option to compare response to openapi spec

