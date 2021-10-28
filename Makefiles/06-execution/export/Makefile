# .EXPORT_ALL_VARIABLES:  
# Uncomment the above line , if you want to export all the variables


new_contents = "hello:\n\\t@echo \$$(cooly)"

all:
	mkdir -p subdir
	echo $(new_contents) | sed -e 's/^ //' > subdir/makefile
	@echo "---MAKEFILE CONTENTS---"
	@cd subdir && cat makefile
	@echo "---END MAKEFILE CONTENTS---"
	cd subdir && $(MAKE)

# Note that variables and exports. They are set/affected globally.
cooly = "Hi from subdicectory!"
export cooly
# This would nullify the line above: unexport cooly

export chk="for checking purpose"

clean:
	rm -rf subdir