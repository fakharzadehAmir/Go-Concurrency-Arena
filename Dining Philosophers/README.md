# Dining Philosophers Problem in Go

## Introduction
This repository contains a Go implementation of the classic concurrency problem known as the Dining Philosophers. The problem illustrates synchronization issues and techniques to resolve them in concurrent algorithm design.

## Implementation Rationale

### Use of Channels
- **Channels:** Channels in Go provide a way for goroutines to communicate safely without explicit locks or condition variables. In this implementation, channels are used as semaphores to represent the chopsticks. Each chopstick channel can hold a single token, ensuring that only one philosopher can hold a chopstick at a time.
- **Channel Operations:** The operations `<-chopsticks[i]` (receiving from a channel) and `chopsticks[i] <- struct{}{}` (sending to a channel) are used to simulate picking up and putting down chopsticks. These operations are atomic and block when the chopstick is not available, ensuring mutual exclusion.

### Context for Timeout Control
- **Context Package:** The context package provides control over the lifetime of the goroutines. In this implementation, a context with a timeout is used to run the dining philosophers for a fixed duration, after which the program gracefully terminates. This helps in avoiding indefinite blocking of goroutines and simulates a real-world scenario where activities are time-bound.
- **Cancel Function:** The `cancel` function provided by the context package is used to signal all goroutines to stop their execution once the timeout is reached.

### Synchronization with WaitGroup
- **WaitGroup:** The `sync.WaitGroup` is used to wait for all philosopher goroutines to finish their execution. It ensures that the main function does not exit prematurely, allowing each philosopher to complete their cycle of thinking, eating, and putting down the chopsticks.

### Problem Solution Strategy
- **Odd-Even Strategy:** To avoid deadlocks, philosophers are assigned an order (odd or even based on their position). Even-numbered philosophers pick up the left chopstick first, followed by the right one, whereas odd-numbered philosophers pick up the right chopstick first. This prevents all philosophers from picking up one chopstick and waiting indefinitely for the other, thus avoiding deadlock.

## Conclusion
This implementation of the Dining Philosophers problem in Go demonstrates how concurrency primitives like channels and context can be used to handle synchronization and communication between goroutines effectively. The use of these primitives provides a clear and concise solution to a classic problem in concurrent programming.

## Key Go Features Demonstrated
- **Channels for Semaphore-like Synchronization:** Utilization of Go channels to enforce mutual exclusion and prevent concurrent access to shared resources (chopsticks).
- **Context for Lifecycle Management:** Use of context to manage the execution time and orderly shutdown of concurrent processes.
- **WaitGroup for Coordination:** Coordination of multiple goroutines using WaitGroup to ensure all philosophers complete their actions before the program terminates.

## Learning Resources
- [Dining Philosophers Problem - Wikipedia](https://en.wikipedia.org/wiki/Dining_philosophers_problem)
- [Go by Example - Channels](https://gobyexample.com/channels)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)