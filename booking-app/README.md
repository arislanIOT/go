Go needs to know the entry point of the language. 


> Go programs are organized as packages.

How do you know which function are in which package. 


Constants - variable value that cannot be changed/ alterterd

When we create vaiable or const and assign the value to it immedeitaly on the same line Go can undetstand the type of the variable. 

so for more robust 
define the variable type explecitly. 


%T is the place holder for the Type of the variable

%v is the place holder for the value of the variable. 

# Array
# Slice abstraction of the array
var var_name [size]var_type

# Loop in Go : How long the application should work. 

// create a map for user 
	// map is a collection of key value pair. Hence 
	// syntax :  var userDate = map[key]value
# Map
All key have the same data type
All value have the same data type

To create an empty map; you can use the build in command called make

var userData = make(map[string]string)
# Struct 
collect different type of data types
eg.: date of birth, subscription boolean, etc..
user entity will have different data types, 

# Sprint
fmt.Sprintf()
this can be used to print a string to a variable as formatted output. 

# Thread (goroutine)
# Sleep function under the package called time

# Handling the blocking code with goroutines
This is why we need concurrency in the application. To make our application cocurrent is easy while doing with the go, 

# Synchronising the goroutine

when main thread exits; the goroutine thread also exits. so if the main thread exit before the goroutine thread completes it will not show  the output. 

# sync.WaitGroup {}
functions of waitgroup

- Add() add the counter
- Wait() // wait till the counter becomes 0
- Done() - removes the thread once // Dereases the counter