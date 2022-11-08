define USAGE
Usage:

    make examples # Build all example binaries

    make partey  # Build partey example

    bin/partey ~/.ssh/config ~/.ssh/config.d "#" # Append the default ssh config file with the partials from config.d
endef

coreLibs = $(wildcard src/*.go)

default:
	$(info $(USAGE))

quickstart: bin/partey ssh

bin/partey: $(coreLibs)
	go build -o bin/partey ./main.go

ssh:
	bin/partey ~/.ssh/config ~/.ssh/config.d "#"

.PHONY: tests
tests:
	go build -o bin/tests tests/*.go && bin/tests