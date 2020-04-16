# sysl-template

Simple REST endpoint generation in Go

## Prerequisites

- [Sysl v0.11.0 or later ](https://sysl.io/docs/install/)
- Go 1.13

## How to use this template;

- Click the `Use this template` button up the top
- ctrl-f and replace all instances of `github.com/anz-bank/sysl-template` with your import path, eg `github.com/foobar/myrepo`
- run `make`
  - note: `make` must be run with a working internet connection, as it fetches transforms and grammars over the network.
- run   `go run main.go`
- congrats! you've just made your first sysl application!


## Development
- Refer to the [Makefile](Makefile) to generate client and server code for all of your applications

## File structure 

api/: contains all API specifications for the generated application

gen/: contains all the generated code for the service

[internal/server/server.go](pkg/server/server.go): The hand-written code that's written; Server config and such

pkg/defaultcallback: contains code that sets up the defaults for generated code. (This will no longer be necessary in future Sysl versions)

When new endpoints are added, they need to be added to the `simple.ServiceInterface` variable in [server.go](server/server.go)

[main.go](main.go): runs the actual server


run `make` to regenerate application code
First you need to edit the start of the Makefile:

```
input = your input sysl file
app = < the app you want to develop>
down = <downstreams in a list separated by spaces>
basepath = <Your current go module path>
```

so: `make input=model/simple.sysl app=Simple` for this example

run `go run main.go` to start the server
