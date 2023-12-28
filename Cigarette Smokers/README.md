# Cigarette Smokers Problem in Go

## Introduction
This repository presents a Go implementation of the Cigarette Smokers problem, a classic example in the study of concurrency to illustrate synchronization and deadlock scenarios. The problem simulates a situation where three smokers need different ingredients to roll a cigarette and smoke.

## Implementation Rationale

### Use of Conditional Variables (`sync.Cond`)
- **Conditional Variables:** The `sync.Cond` in Go is used to coordinate the smoking process among smokers. It allows goroutines (representing smokers) to wait for certain conditions (availability of ingredients) and be notified when those conditions are met.
- **Conditional Wait and Signal:** The `Wait` method is used for smokers to wait for their required ingredients. The `Signal` method is called by the agency to notify a waiting smoker when the ingredients are available. This approach avoids busy waiting and ensures efficient use of resources.

### Synchronization with WaitGroup
- **WaitGroup:** The `sync.WaitGroup` is essential for synchronizing multiple goroutines in the program. It ensures that the main function waits until all smoker and agency goroutines complete their execution, providing a controlled and orderly termination of the program.

### Randomness in Ingredient Distribution
- **Randomness:** The randomness in ingredient distribution by the agency is crucial for simulating a realistic scenario and avoiding biases in the availability of ingredients, ensuring that no single smoker is favored repeatedly.

## Conclusion
The implementation of the Cigarette Smokers problem in Go is a demonstration of how synchronization primitives like conditional variables, and WaitGroup can be effectively employed to manage complex concurrency scenarios. This solution highlights the importance of efficient inter-goroutine communication and synchronization in concurrent programming.

## Key Go Features Demonstrated
- **Conditional Variables for Synchronization:** Use of `sync.Cond` for managing dependencies between concurrent processes and efficient signaling.
- **WaitGroup for Coordinating Goroutines:** Using WaitGroup to ensure that all components of the concurrent system (smokers and agency) complete their tasks.

## Learning Resources
- [Cigarette Smokers Problem - Wikipedia](https://en.wikipedia.org/wiki/Cigarette_smokers_problem)
- [Go by Example](https://gobyexample.com/)