SYSL  =  ./syslw
ARRAI  =  ./arraiw
CODEGEN_ROOT = internal/gen
SPECS_ROOT = specs/frontend/petdemo

# -- Generate --------------------------------------------------------------------

gen: backend-gen frontend-gen

# frontend code
.PHONY: frontend-gen
.PHONY: frontend
frontend-gen: petdemo-gen

.PHONY: Petdemo
PETDEMO = Petdemo

frontend = rest-app
.PHONY: petdemo-gen
petdemo-gen: $(PETDEMO) frontend
	$(run-codegen)

# backend code
.PHONY: backend-gen
.PHONY: backend
backend-gen: petstore-gen flickr-gen

.PHONY: Flickr
FLICKR = Flickr

.PHONY: flickr-gen
flickr-gen: $(FLICKR) backend
	$(run-codegen)

.PHONY: Petstore
PETSTORE = Petstore

.PHONY: petstore-gen
petstore-gen: $(PETSTORE) backend
	$(run-codegen)

SYSL_TEMPLATE_ROOT = github.com/anz-bank/sysl-template
backend = rest-app

define run-codegen
	rm -rf $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
	rm -rf $(SPECS_ROOT)/sysl.json
	mkdir -p $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
	$(SYSL) pb --mode json $(SPECS_ROOT)/petdemo.sysl >> $(SPECS_ROOT)/sysl.json
	$(ARRAI) run service_stub.arrai \
		$(SYSL_TEMPLATE_ROOT)/$(CODEGEN_ROOT) $(SPECS_ROOT)/sysl.json \
		$< $($(word 2,$^)) | tar xf - -C $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
	goimports -w $(CODEGEN_ROOT)/$(shell echo $< | tr A-Z a-z)
endef

# -- Build ---------------------------------------------------------------------

build: build-flickr build-petstore build-petdemo

.PHONY: build-flickr
build-flickr:
	go mod tidy
	go build -v -o flickrserver ./cmd/flickrserver

.PHONY: build-petstore
build-petstore:
	go mod tidy
	go build -v -o petstoreserver ./cmd/petstoreserver

.PHONY: build-petdemo
build-petdemo:
	go mod tidy
	go build -v -o petdemoserver ./cmd/petdemoserver

install:  ## Install the server binary
	go install ./...

.PHONY: install

# -- Test ---------------------------------------------------------------------

test: build-flickr build-petstore build-petdemo  ## Execute application unit tests
	go test -race -coverprofile=$(COVERFILE) ./...

.PHONY: test

# -- Run ---------------------------------------------------------------------

run:  ## Run the server
	go run cmd/flickrserver/main.go flickr &
	go run cmd/petstoreserver/main.go petstore &
	go run cmd/petdemoserver/main.go petdemo &

.PHONY: run

# -- Docs ---------------------------------------------------------------------

.PHONY: docs
docs:  ## Generate service documentation (requires sysl-catalog installation)
	rm -r docs/* || true
	sysl-catalog specs/frontend/petdemo/petdemo.sysl -o docs
