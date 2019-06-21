### Contents
- [Motivation](#Motivation)
- [Learnt](#Learnt)
- [Resources](#Resources)
- [Benchmarking Results](#Benchmarking-Results)
  - [Sequential Sort](#Sequential-Sort)
  - [Concurrent Sort](#Concurrent-Sort)

## Motivation

I wrote merge-sort-go to learn about Golang's concurrency building blocks - Goroutines, Channels and `Select`.

_The concurrent merge sort performs poorly compared to the serial implementation. I'm looking at the profiling to zero down on what is causing the performance hit. My current hunch is that in the concurrent implementation, a lot of goroutines (and corresponding channels) are spawned resulting in an excess overhead due to blocking and communication between the goroutines as opposed to the sorting computation_  

## Learnt

1. Working with Slices
   
   1. Popping an element out of a Slice - Use the slice syntax `[1:]` or `[:len(arr)-1]`
   
   2. Copying/ Cloning a slice - Using the `copy` function while taking note that the destination slice should have a minimum length as the source array or elements will be left behind

2. Generating Random numbers - Using `rand.Intn()` from `math/rand` package

3. Working with Goroutines, Channels and `Select`

4. Testing Go code

5. Benchmarking performance of Go code

6. Profiling Go code

## Resources

1. For Testing and Benchmarking - [Chapter 9, Go in Action](https://www.manning.com/books/go-in-action)

2. For Concurrency in Golang - [Concurrency in Go: Tools and Techniques for Developers](https://www.amazon.com/Concurrency-Go-Tools-Techniques-Developers/dp/1491941197)

3. For Profiling
   1. [Make Your Go Programs Lightning Fast With Profiling](https://code.tutsplus.com/tutorials/make-your-go-programs-lightning-fast-with-profiling--cms-29809)
   2. [Go Blog - Profiling Go Programs](https://blog.golang.org/profiling-go-programs)
   3. [Getting started with Go CPU and memory profiling](https://flaviocopes.com/golang-profiling/)

## Benchmarking Results

**Context**

Go OS - Linux, 
Go Arch - amd64

### Sequential Sort

`go test -v -run="none" -bench="BenchmarkSeqSort" -benchtime 5s`


| Array Size | Nanosec per op | No. of runs (Benchmarking) |
| :--: | :--: | :--: |
| 1000 | 385,186 | 20,000 |
| 10,000 | 5,095,454 | 2000 |
| 100,000 | 61,157,041 | 100 |
| 1,000,000 | 629,770,486 | 10 |

### Concurrent Sort

`go test -v -run="none" -bench="BenchmarkConcSort" -benchtime 5s`

| Array Size | Nanosec per op | No. of runs (Benchmarking) |
| :--: | :--: | :--: |
| 1000 | 1,607,284 | 5,000 |
| 10,000 | 27,618,040 | 300 |
| 100,000 | 314,162,988 | 20 |
| 1,000,000 | 3,194,230,361 | 2 |
