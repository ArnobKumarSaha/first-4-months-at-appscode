# '*' and "%" are wildcards in makefile


thing_wrong := *.o # Don't do this!  it equals to "*.o"
thing_right := $(wildcard *.c)  # use with wildcard function



all: one two three four print

# Fails, because $(thing_wrong) is the string "*.o"
one: $(thing_wrong)

# Stays as *.o if there are no files that match this pattern :(
two: *.o 

# Works as you would expect! In this case, it does nothing.
three: $(thing_right)

# Same as rule three
four: $(wildcard *.c)



print:
	@echo ${thing_right}
	@echo ${thing_wrong}

#print: $(wildcard *.c)
#    ls -la  $?