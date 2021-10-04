SYSLGO_SYSL=specs/petdemo.sysl
SYSLGO_PACKAGES=Petdemo
SYSLGO_APP.Petdemo = Petdemo

-include local.mk
include codegen.mk

.PHONY: clean
clean:
	rm -rf internal/gen

.PHONY: test
test: gen-all-servers
	go test -v ./...

.PHONY: integration-test
integration-test: gen-all-servers
	go test -v ./... -run '_Integration$$'
