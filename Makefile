NNDB_VERSION=sr28asc
NNDB_ZIP=$(NNDB_VERSION).zip
NNDB_URL=https://www.ars.usda.gov/SP2UserFiles/Place/12354500/Data/SR/SR28/dnload/sr28asc.zip

default: build

prepare: unzip dependencies tools

tools: 	
	go get -u github.com/Masterminds/glide
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update 

dependencies:
	glide up 

clean: 
	rm -rf $(NNDB_VERSION) $(NNDB_ZIP) bin debug

$(NNDB_ZIP): 
	wget $(NNDB_URL) 

unzip: $(NNDB_ZIP)
	unzip -o $(NNDB_ZIP) -d $(NNDB_VERSION)  

vet: 
	go vet `glide nv`

lint: 
	gometalinter --vendor --disable=gotype --dupl-threshold=90 --deadline=10s ./...

fmt:
	gofmt -s -w .

compile: 
	go build `glide nv` 

build: fmt vet lint compile

test: fmt vet lint
	go test `glide nv` -cover -race
