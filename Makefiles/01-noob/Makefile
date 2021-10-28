#target: prerequisites
#<TAB> recipe


# .DEFAULT_GOAL can call one target at a time.
.DEFAULT_GOAL := all


# as /bin/sh is the default shell for makefiles , override the SHELL variable to use bash
SHELL := /bin/bash


# .PHONY actually says which targets are phony, to prevent 'make' from confusion
.PHONY: all say_hello generate clean


#The first target is the DEFAULT_GOAL if nothing is specified.
all: say_hello generate

say_hello:
	@echo "Hello world"

generate:
	@echo "Creating empty text files..."
	touch file-{1..10}.txt

clean:
	@echo "Cleaning up..."
	rm *.txt