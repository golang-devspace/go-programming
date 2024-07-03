# Struct Data Structure
Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

https://gobyexample.com/structs

## &Variable
Gives the memory address of variable

## *pointer
Gives the value at the memory address.

## Difference between Array and Slices in Go
Array
* Primitive data structure
* Cant'n be resized
* Rarely used directly 

Slices
* can grow and shrink
* Used 99% of the time for lists of elements

## Value Types & Reference Types

|Value Types    |Reference Types    |
|---------------|-------------------|
|int            |slice              |
|float          |maps               |
|string         |channels           |
|bool           |pointers           |
|structs        |functions          | 