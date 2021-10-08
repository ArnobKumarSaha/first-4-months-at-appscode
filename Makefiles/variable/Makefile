
# if @ is not given before echo .. It prints two lines
# echo "This will run first" 
# This will run first
# But if @ is given before echo , it only prints the second line.

all: some_file just_print var_example multiple_targets

some_file: other_file  # some file is the target name
	@echo "This will run second, because it depends on other_file"
	touch some_file  
# but here, some_file is the filename

other_file:
	@echo "This will run first"
	touch other_file



# very basic example of variable
var1 = Hey
var2 = dude
dot = "-------------------"

just_print:
	@echo ${dot}
	@echo $(var1)
	@echo ${var2}
	@echo ${dot}



# A standard example of variable
files = file1 file2
var_example: $(files)
	echo "Look at this variable: " $(files)
	touch var_example

file1:
	touch file1
file2:
	touch file2

clean:
	rm -f file1 file2 some_file other_file var_example





multiple_targets: f1 f2

f1 f2:
	echo $@

# Equivalent to:
# f1.o
#     echo $@
# f2.o
#     echo $@
