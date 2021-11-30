define USAGE
Usage:

    make examples # Build all example binaries

    make partey  # Build partey example

    bin/partey ~/.ssh/config ~/.ssh/config.d "#" # Append the default ssh config file with the partials from config.d
endef

coreLibs = $(wildcard src/*.go)

default:
	$(info $(USAGE))

examples: partey

partey: bin/party

bin/party: $(coreLibs)
	go build -o bin/partey ./examples/main.go $(coreLibs)

ssh:
	bin/partey ~/.ssh/config ~/.ssh/config.d "#"

.PHONY: tests
tests:
	go build -o bin/tests tests/*.go