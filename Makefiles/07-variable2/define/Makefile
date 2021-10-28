# Overrides command line arguments
override option_one = did_override
# Does not override command line arguments
option_two = not_override




one = export blah="I was set!"; echo $$blah

define two
export blah=set
echo $$blah
endef

# One and two are different.



all: 
	echo $(option_one)
	echo $(option_two)
	@echo "This prints 'I was set'"
	@$(one)
	@echo "This does not print 'I was set' because each command runs in a separate shell"
	@$(two)