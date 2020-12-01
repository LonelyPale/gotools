ifndef GOOS
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S), Darwin)
		GOOS := darwin
	else ifeq ($(UNAME_S), Linux)
		GOOS := linux
	else
		$(error "$$GOOS is not defined. If you are using Windows, try to re-make using 'GOOS=windows make ...' ")
	endif
endif


build: sysinfo

sysinfo: env
	@echo "Building $@ to bin/"
	@go build -o bin/$@ cmd/$@/main.go

install: env
	@echo "Installing sysinfo to $(GOPATH)/bin"
	@go install ./cmd/sysinfo

image: Dockerfile
	@echo "Building gotools to docker image"
	@docker build -t gotools:latest .

clean:
	@echo "Cleaning gotools binaries ..."
	@rm -rf $(GOPATH)/bin/sysinfo
	@rm -rf bin/sysinfo
	@echo "Done."

env:
	@go env -w GOPROXY="https://goproxy.cn,direct" && \
    go env -w GO111MODULE=on && \
    go env GOPROXY


.PHONY: build clean env
