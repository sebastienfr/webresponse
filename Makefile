# -----------------------------------------------------------------
#
#        ENV VARIABLE
#
# -----------------------------------------------------------------

# go env vars
GO=$(firstword $(subst :, ,$(GOPATH)))
# list of pkgs for the project without vendor
PKGS=$(shell go list ./... | grep -v /vendor/)
export GO15VENDOREXPERIMENT=1

# -----------------------------------------------------------------
#        Version
# -----------------------------------------------------------------

# version
VERSION=0.0.1-SNAPSHOT
BUILDDATE=$(shell date -u '+%s')
BUILDHASH=$(shell git rev-parse --short HEAD)
VERSION_FLAG=-ldflags "-X main.Version=$(VERSION) -X main.GitHash=$(BUILDHASH) -X main.BuildStmp=$(BUILDDATE)"

# -----------------------------------------------------------------
#        Main targets
# -----------------------------------------------------------------

help:
	@echo
	@echo "----- BUILD ------------------------------------------------------------------------------"
	@echo "all                    clean and build the project"
	@echo "build                  build all libraries and binaries"
	@echo "----- TESTS && LINT ----------------------------------------------------------------------"
	@echo "test                   test all packages"
	@echo "format                 format all sources"
	@echo "lint                   lint all packages"
	@echo "start                  starts program locally"
	@echo "stop                   stops all program instances on localhost"
	@echo "----- SERVERS AND DEPLOYMENTS ------------------------------------------------------------"
	@echo "dockerBuild            build the program docker image"
	@echo "dockerUp               start program microservice infrastructure on docker"
	@echo "dockerStop             stop program microservice infrastructure on docker"
	@echo "dockerBuildUp          stop, build and start program microservice infrastructure on docker"
	@echo "dockerWatch            starts a watch of docker ps command"
	@echo "dockerLogs             show logs of program microservice infrastructure on docker"
	@echo "----- OTHERS -----------------------------------------------------------------------------"
	@echo "clean                  clean the project"
	@echo "help                   print this message"

all: clean build

clean:
	@go clean
	@rm -Rf .tmp
	@rm -Rf .DS_Store
	@rm -Rf *.log
	@rm -Rf *.out
	@rm -Rf *.lock
	@rm -Rf build

build: format
	@go build -v $(VERSION_FLAG) -o $(GO)/bin/webresponse webresponse.go

format:
	@go fmt $(PKGS)

test:
	@go test -v $(PKGS)

lint:
	@golint ./.
	@go vet $(PKGS)

start:
	@webresponse -port 8020 -path "/"

stop:
	@killall webresponse

# -----------------------------------------------------------------
#        Docker targets
# -----------------------------------------------------------------

dockerBuild:
	docker build -t sebastienfr/webresponse:latest .

dockerWatch:
	@watch -n1 'docker ps | grep webresponse'
