.DEFAULT_GOAL := hey

pre:
	touch pre

hey: one two pre
	# Outputs "hey", since this is the first target
	@echo $@

	# Outputs all prerequisites newer than the target
	@echo $?

	# Outputs all prerequisites
	@echo $^

	touch hey

one:
	touch one

two:
	touch two

clean:
	rm -f hey pre one two
