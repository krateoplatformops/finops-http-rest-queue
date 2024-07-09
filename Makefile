ARCH?=amd64
REPO?=#your repository here 
VERSION?=0.1

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=$(ARCH) go build -o ./bin/http-rest-queue main.go

container:
	docker build -t $(REPO)finops-http-rest-queue:$(VERSION) .
	docker push $(REPO)finops-http-rest-queue:$(VERSION)
