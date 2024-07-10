# Hard Mode

* Create a program that reads the contents of the text file then prints its content to the terminal. 
* The file to open should be provided as an argument to the program when it is executed at the terminal.
* To read in the arguments provided to the program, you can reference the variable 'os.Args', which is a slice of type string
* To open a file checkout the documentation for the Open function in the os package.
* What interface does the File type implement?
* If the File type implements the Reader interface, you might be able to reuse that io.Copy function.
