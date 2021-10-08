#target: prerequisites
#<TAB> recipe


# Usage:
# make        # compile all binary
# make clean  # remove ALL binaries and objects

.PHONY = all clean 		# phony targets, all target that are not files

# variables
COMPILER = gcc                
LINKERFLAG = -lm

# ':=' for direct assignment , it is calculated only once
# '=' for recursive assignment, it is calculated everytime it is mentioned.
# '?=' for conditional assignment, only calculated if not initialized yet.

SRCS := $(wildcard *.c)  #  all files with the .c extension
BINS := $(SRCS:%.c=%)   # substitution reference.  foo.c converts to foo

#each values of ${BINS} as individual targets
all: ${BINS}  



# rules
# $@ matches the target
# $< matches pre-requisites

%: %.o
	@echo "Checking.."
	${COMPILER} ${LINKERFLAG} $< -o $@

%.o: %.c
	@echo "Creating object.."
	${COMPILER} -c $<

clean:
	@echo "Cleaning up..."
	rm -rvf *.o ${BINS}
