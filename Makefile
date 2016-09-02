NNDB_VERSION=sr28asc
NNDB_ZIP=$(NNDB_VERSION).zip
NNDB_URL=https://www.ars.usda.gov/ARSUserFiles/80400525/Data/SR/SR28/dnload/sr28asc.zip
default: build

prepare: unzip dependencies tools

tools: 	
	go get -u github.com/Masterminds/glide
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

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
	glide nv | xargs -n1 gometalinter --disable=gotype --dupl-threshold=90 --deadline=30s --vendor

fmt:
	gofmt -s -w .

compile: 
	go build -race `glide nv` 

build: fmt vet lint compile

test: fmt vet lint
	go test `glide nv` -cover -race
