BIN=ft_linear_regression_predict

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFLAGS=-ldflags '-w -s'

DOCKER=docker
DOCKERBUILD=$(DOCKER) build
DOCKERRUN=$(DOCKER) run
DOCKERPUSH=$(DOCKER) push

DEP=dep ensure

all: go-build

go-build:
	@echo "Golang build executable..."
	$(DEP)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) $(GOBUILD) -o $(BIN) $(GOFLAGS) main.go
	mv $(BIN) ../ft_linear_regression/bin