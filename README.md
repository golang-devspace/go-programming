## Go lang
Go is a static type language 

## Go package types 
Go has two types of packages 
1. Executable -> Generates a file that we an run 
2. Reusable     -> Code used as helpers. Good place to put reusable logic

## Go CLI 
 1. go build    -> compiles a bunch of go source files
 2. go run      -> Compiles and executes on or more files 
 3. go fmt      -> Formats all the code in each file in the current directory
 4. go install  -> Compiles and install go package
 5. go get      -> Downloads the raw source code of someone else package 
 6. go test     -> Runs ant test associated with the current project. 
 7. go mod init -> 
 7. go list -m all -> 

## Managing dependency in Go 

## Install Go
https://go.dev/doc/install


## go.mod and go.sum files in Golang project
Reference : https://go.dev/ref/mod
 * go.mod is used for dependency management in golang projects.
 * All the modules which are needed or to be used in the project are maintained in go.mod file.
 * For all the packages we are going to import/use in our project, it will create an entry of those modules in go.mod. 
 * Having a go mod file saves the efforts of running the go get command for each dependent module to run  the project successfully.

### go.sum 
After running any package building command like go build, go test for
the first time, it will install all the packages with specific versions i.e which are the latest at that moment.
It will also create a go.sum file which maintains the checksum so when you run the project again it will not install all packages again. But use the cache which is stored inside $GOPATH/pkg/mod directory (module cache directory)
```
$ go mod verify
# Allows you to verify the checksum of dependencies
```

# Data structure is Go
## Struct
It is data structure. A collection of properties that are related together.
Struct in different language
Javascript -> plain object
ruby       -> hash
python     -> Dictionary

 ## Difference between Map and Struct
 ### Map
 * All keys must be of same type
 * Use to represent a collection of related properties.
 * All value must be of same type.
 * Don't need to know all the kays at compile time.
 * Keys are indexed - we can iterate over them.
 * Reference type
### Struct
* Values can be of different types
* You need to know all the different fields at compile time.
* keys don't support  indexing
* Use to represent a thing with a lot of different properties
* Value type 

## Interface in Go
* Interface are not generic types in go like Java
* Interface are implicit
* Interfaces are a contract to help us manage types
* Interfaces are tough. Step #1 is understanding how to read them.