REGISTRY = mmungdong
PROJECT = chatgpt-web
BASE_VERSION = v4
GIT_VERSION = $(shell git rev-parse --short HEAD)

.PHONY: build-local
build-local:
	go build -v --ldflags="-w -X main.Version=$(GIT_VERSION)" -o dist/server cmd/*.go

.PHONY: build
build:
	mkdir -p dist && docker run --rm -ti -e GOPROXY=https://goproxy.cn,direct -v $(GOPATH):/go -v `pwd`:/app -w /app golang:1.19-alpine \
	go build -v --ldflags="-w -X main.Version=$(GIT_VERSION)" -o dist/server cmd/*.go

apiserver: build-local
	cd ./docker/apiserver
	docker build -f docker/apiserver/Dockerfile -t apiserver:$(GIT_VERSION) .

tokenizer:
	cd ./docker/tokenizer
	docker build -t tokenizer:$(GIT_VERSION) .

# package: build
# 	docker build -t $(REGISTRY)/$(PROJECT):$(GIT_VERSION) .

release: package
	docker tag $(REGISTRY)/$(PROJECT):$(GIT_VERSION) $(REGISTRY)/$(PROJECT):latest
	docker push $(REGISTRY)/$(PROJECT):$(GIT_VERSION)
	docker push $(REGISTRY)/$(PROJECT):latest

base:
	docker build -t $(REGISTRY)/$(PROJECT)-base:$(BASE_VERSION) -f Dockerfile.base .

release-base: base
	docker push $(REGISTRY)/$(PROJECT)-base:$(BASE_VERSION)

clean:
	rm -rf dist
	docker images | grep -E "$(REGISTRY)/$(PROJECT)" | grep -v "base"  | awk '{print $$3}' | uniq | xargs -I {} docker rmi --force {}