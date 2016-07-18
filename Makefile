NNDB_VERSION=sr28asc
NNDB_ZIP=$(NNDB_VERSION).zip
NNDB_URL=https://www.ars.usda.gov/SP2UserFiles/Place/12354500/Data/SR/SR28/dnload/sr28asc.zip

default: build

prepare: unzip dependencies tools

tools: 	
	@go get -u github.com/Masterminds/glide
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install --update > /dev/null 2>&1

dependencies:
	@glide up > /dev/null 2>&1

clean: 
	@rm -rf $(NNDB_VERSION) $(NNDB_ZIP) bin debug

$(NNDB_ZIP): 
	@wget $(NNDB_URL) > /dev/null 2>&1

unzip: $(NNDB_ZIP)
	@unzip -o $(NNDB_ZIP) -d $(NNDB_VERSION)  > /dev/null 2>&1

vet: 
	@go vet `glide novendor`

lint: 
	@gometalinter --vendor --disable=gotype

fmt:
	@gofmt -s -w .

compile: 
	@go build -o bin/nndb

build: fmt vet lint compile

test: fmt vet lint
	@go test -cover -race
