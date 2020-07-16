# sysl-template

Demo REST Application generated using Sysl-Go

## Prerequisites

- [Sysl v0.140.0 or later ](https://sysl.io/docs/install/)
- [Arr.ai v0.87.0 or later](https://github.com/arr-ai/arrai)
- Go 1.14

## Building & Running

- run `make gen`. 
  - This generates the Go code.
- run `make build`. 
  - This builds the generated code using the installed Go runtime.
- run `make run`.
  - This runs the petdemo app and the downstream apps i.e. flickr and petstore
- `curl http://localhost:8080/random-pet-pic` or open in browser.
  - Output: `{"name":"white golden retriever","uri":"http://www.example.com/dog/wgr"}`

Congrats! You've just built and run the sysl-go demo application!

## Development

- Refer to the [Makefile](Makefile) to generate client and server code for all of your applications.

## File structure

[internal/gen/petdemo]: Contains the all the generated code for the petdemo app.
[internal/gen/flickr]: Contains the generated code for the flickr downstream service.
[internal/gen/petstore]: Contains the generated code for the petstore downstream service.

[internal/petdemo]: Hand crafted config and server for petdemo app.
[internal/petdemo/handlers]: Hand crafted handlers for the petdemo app.

[internal/flickr]: Hand crafted config and server for flickr downstream service.
[internal/flickr/handlers]: Hand crafted handlers for the flickr downstream service.

[internal/petstore]: Hand crafted config and server for petstore downstream service.
[internal/petstore/handlers]: Hand crafted handlers for the petstore downstream service.

## Configuration

The configuration is driven through the YAML files in the repository root.
- [config-flickr.yml](config-flickr.yml)
- [config-petdemo.yaml](config-petdemo.yaml)
- [config-petstore.yaml](config-petstore.yaml)
The arrai and sysl versions can be locked in the arraiw.properties and syslw.properties files in repository root.

## Application specs

The application specs are stored in the form of yaml files in specs directory. 
- specs/backend/flickr/flickr.yaml
- specs/backend/petstore/petstore.yaml
- specs/frontend/petdemo/petdemo.yaml
