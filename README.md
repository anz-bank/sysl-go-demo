# sysltemplate

Simple REST endpoint generation in go

## Prerequisites

- [Sysl v0.11.0 or later ](https://sysl.io/docs/install/)

## How to use this template;

- Click the `Use this template` button up the top
- ctrl-f and replace all instances of `github.com/anz-bank/sysltemplate` with your import path, eg `github.com/foobar/myrepo`
- run `make`
  - note: `make` must be run with a working internet connection, as it fetches transforms and grammars over the network.
- run   `go run main.go`
- congrats! you've just made your first sysl!


## Development
- Refer to the make file to generate client and server code for all of your applications

## File structure 

model/: contains the sysl application that's built

gen/: contains all the generated code for the service

implementation/: The manual code that's written; Server config and such

gencallback.go: has the server config `Callback` and `Config` structs which are used in the `LoadServices` func that sets up the server. 

methods.go: Contains all these Functions, (yes, functions), that are then composed into the `ServiceInterface` struct in `LoadServices`

When new endpoints are added, they need to be added to the `simple.ServiceInterface` variable in `LoadServices`

main.go: runs the actual server


run `make` to regenerate application code
First you need to edit the start of the Makefile:

```
input = your input
app = < the app you want to develop>
down = <downstreams in a list separated by spaces>
basepath = <Your current basepath to the generated documentation>
```

so: `make input=model/simple.sysl app=Simple` for this example

run `go run main.go` to start the server
