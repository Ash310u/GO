# GO Language

## FUN FACTS You Should Care About

1. Statically typed language: variable types are known at compile time.
```
```go
var x int = 5 // x is always int
```
2. Strongly typed language: you can't mix incompatible types without explicit conversion.
```
```go
var x int = 10
var y string = "hello"
// x = y // this will cause a compile-time error
```
3. Go is compiled: the code is turned into a binary before running, making it much faster!
```
```go
for i := 0; i < 1000000; i++ {} // Looping from 0 to 1 million in Go takes ~50ms, but in Python ~6s!
```
4. Fast compile time: Go compiles code very quickly, even for large projects.

5. Built-in concurrency: you can easily run code in parallel using goroutines.
```
```go
go func() { fmt.Println("This runs concurrently!") }()
```
6. Simplicity (incl. garbage collection): Go handles memory cleanup automatically, keeping your code simple.
```
```go
var s = make([]int, 1000) // You do not worry about memory deallocation, Go's GC handles it.
```


