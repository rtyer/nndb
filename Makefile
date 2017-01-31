NNDB_VERSION=sr28asc
NNDB_ZIP=$(NNDB_VERSION).zip
NNDB_URL=https://www.ars.usda.gov/ARSUserFiles/80400525/Data/SR/SR28/dnload/sr28asc.zip
default: build

prepare: unzip dependencies tools

tools: 	
	go get -u github.com/Masterminds/glide
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install -u

dependencies:
	glide up 

clean: 
	rm -rf $(NNDB_VERSION) $(NNDB_ZIP) bin debug
	find . -name debug.test | xargs rm

$(NNDB_ZIP): 
	wget $(NNDB_URL) 

unzip: $(NNDB_ZIP)
	unzip -o $(NNDB_ZIP) -d $(NNDB_VERSION)  

vet: 
	go vet `glide nv`

lint: 
	gometalinter --disable=gotype --dupl-threshold=120 --deadline=30s --vendor ./...

fmt:
	gofmt -s -w .

build: fmt vet lint 
	go install -race `glide nv`

quick: 
	go install -race `glide nv`

test: fmt vet lint
	go test `glide nv` -cover -race