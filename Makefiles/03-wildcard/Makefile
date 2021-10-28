# '*' and "%" are wildcards in makefile

SHELL := /bin/bash

thing_wrong := *.c # Don't do this!  it equals to "*.c"
thing_right := $(wildcard *.c)  # use with wildcard function



all: one two three four print printw 

# Fails, because $(thing_wrong) is the string "*.o"
one: $(thing_wrong)

var := *.c

# Stays as *.c if there are no files that match this pattern :(
two: 
	@echo var $(var)

# Works as you would expect! In this case, it does nothing.
three: $(thing_right)

# Same as rule three
four: $(wildcard *.c)



print:
	@echo thing_right ${thing_right}
	@echo thing_wrong ${thing_wrong}

printw: $(wildcard *.c)
	ls -la  $?  
#Here $? is equals to 'a.c b.c'

#print: $(wildcard *.c)
#    ls -la  $?