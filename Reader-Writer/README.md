# Reader-Writer Problem in Go

## Introduction
This repository delves into the Reader-Writer problem, a classic synchronization challenge in concurrent programming. It illustrates the coordination between multiple readers and writers attempting to access a shared resource, with an emphasis on ensuring that no reader is starved and writers have fair access.

## Implementation Rationale

### Synchronization with Mutexes
- **sync.Mutex & sync.RWMutex:** We use these synchronization primitives to manage access to the shared byte array. The `sync.Mutex` ensures exclusive access for writers, while `sync.RWMutex` allows multiple readers to access the resource concurrently without being blocked by other readers, thus preventing reader starvation.

### Concurrency Management
- **Context and CancelFunc:** The `context` package provides control over the lifecycle of goroutines. It is used here to manage timeouts and cancel operations, ensuring that the application does not hang indefinitely. The `CancelFunc` allows for a controlled shutdown of goroutines when the context's deadline is exceeded or cancellation is triggered.

### Coordination of Goroutines
- **sync.WaitGroup:** This is used to wait for a collection of goroutines to finish executing. In our code, `WaitGroup` ensures that the main function only exits after all reader and writer goroutines have completed their operations, preventing any premature termination of the program.

## Synchronization Details
- **Reader Count Management:** We utilize a separate mutex to safely increment and decrement the count of active readers. This allows us to log the number of concurrent reads without interfering with the reading and writing operations on the shared resource.
- **Writing with Random Delays:** To simulate a more realistic scenario with variable writing speeds, we introduce random delays in the writer's loop. This can help visualize the writer's behavior in a concurrent environment.

## Graceful Termination
- **Use of Context and CancelFunc:** We incorporate context with a timeout to gracefully handle the termination of goroutines. This ensures that after a specified duration, all goroutines cease their operations, allowing the program to exit cleanly without leaving any operations hanging.

## Learning Resources
- [Reader-Writer Problem - Wikipedia](https://en.wikipedia.org/wiki/Readers%E2%80%93writers_problem)
- [Go by Example - Mutexes](https://gobyexample.com/mutexes)
- [Go by Example - Context](https://gobyexample.com/context)
