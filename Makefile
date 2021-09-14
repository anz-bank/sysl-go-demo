SYSLGO_SYSL=specs/petdemo.sysl
SYSLGO_PACKAGES=Petdemo
SYSLGO_APP.Petdemo = Petdemo

-include local.mk
include codegen.mk

clean:
	rm -rf internal/gen
.PHONY: clean

test: gen-all-servers
	go test -v ./...
PHONY: .test
