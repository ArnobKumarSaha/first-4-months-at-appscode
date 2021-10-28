
obj_files = a.result b.o lose.o
src_files = a.raw b.go lose.go

CC = gcc
CFLAGS = -g

all: $(obj_files)

# error in this file...
# gcc: error: b.o: No such file or directory

$(filter %.o,$(obj_files)): %.o: %.go
	@echo "target: $@ prereq: $<"
#$(CC) -c $(CFLAGS) $(CPPFLAGS) $< -o $@

$(filter %.result,$(obj_files)): %.result: %.raw
	@echo "target: $@ prereq: $<" 
	touch $@


%.go %.raw:
	touch $@

clean:
	rm -f $(src_files) $(obj_files)