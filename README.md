# Go Data Race Example

This repository demonstrates a simple example of a data race in Go.  Data races occur when multiple goroutines access and modify shared data concurrently without proper synchronization mechanisms in place.  This leads to unpredictable and unreliable program behavior.

## The Bug

The `bug.go` file contains a program that launches multiple goroutines, each incrementing a shared variable `count`. Because the `count` variable is not protected by a mutex or other synchronization primitive, the final value of `count` is not deterministic and can vary between executions.

## The Solution

The `bugSolution.go` file presents a corrected version of the code.  It employs a `sync.Mutex` to protect access to the `count` variable, preventing data races and ensuring consistent results. 

This example highlights the importance of proper synchronization techniques when working with concurrency in Go to avoid subtle and difficult-to-debug issues.