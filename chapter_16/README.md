# Concurrent Work
- This is the thing I have been wanting to study and understand the most when it comes to coding in Go.
- This section talks about concurrency and how to achieve it. The section will talk about the features that are needed to perform concurrent work.
- The first feature to talk about is a Goroutine; and we will understand how to use this to achieve concurrency.
- Next, we will learn all about WaitGroups and how to use them to synchronize the execution of several Goroutines.
- We will learn how to implement synchronized and thread-safe changes to variables shared across different Goroutines using atomic changes.
- We will learn about mutexes and how to use these to synchronize more complex atomic changes.
- Finally, we will experiment with channel functionalities and use message tracking to track the completion of a task.


## Introduction
- For software meant to be used by several thousands of people at the same time, they have to be programmed to function in certain ways. An example of this is the web server that serves several users at the same time.
- For this kind of service/situation, it calls for a different kind of programming where different tasks can be executed at the same time and independently from each other.
- There is **parallel computation**, where tasks are computed simultaneously and there is **concurrent computation** where the program will execute a small piece of a task at a time until all the tasks lined up in the thread are completed.
- In Go, tasks can be concurrently executed whilst also being executed in parallel. This occurs when there are multiple CPU cores available to be used.

## Goroutines
- A simple way to illustrate Goroutines is this:
  - suppose several people have to hammer some nails into a wall. Each person got a different number of nails but there is only one hammer to use between them.
  - Each person uses the hammer for one nail then passes the hammer to another person, and so on.
  - The person with the fewest nails will finish first.
  - This is how Goroutines work.
- In Goroutines, Go allows multiple tasks to run at the same time; these multiple tasks are called **couroutines**.
  - **coroutines** are routines (read tasks) that can co-run inside the same process and are totally concurrent.
- Goroutines do not share memory and that is what makes them different from threads. We will see how east it is to pass variables across goroutines in code and also how doing this might lead to some unexpected behaviour.
- Writing a goroutine is nothing special, they are just normal functions.
  - a function can easily become a goroutine just by adding the word **go** before calling the function.
- Considering a function called **hello**
```go
func hello() {
  fmt.Println("Hello World!")
}
```
  - in order to call this function as a Goroutine, we do the following `go hello()`.
```go
func main() {
  fmt.Println("Start")
  go hello()
  fmt.Println("End")
}
```
  - *the way it will execute*:
  - the code starts by printing "Start", then it calls the "hello()" function.
  - the code then goes straight to printt "End" without waiting for the "hello()" function to complete.
- It does not matter how long the code will take to run the "hello()" function as both it and the "main()" are called to be run independently.
- **NOTE**: Go is not a parallel language but concurrent, which means that Goroutines do not work in an independent manner but each Goroutine is split into smaller parts and the Goroutines run a subpart at a time.
*see exercise 16.01*


## WaitGroup
- **HEADS_UP**: Even if a program does not explicitly use Goroutines via the **go** call, it still uses one Goroutine ie the main routine.
- When we explicitly create a Goroutine, we will be running two Goroutines, the main() and the one created. 
- In order to synchronize these two Goroutines, Go has provided a functionality called **WaitGroup**.
- A WaitGroup can be defined with the following code: `wg := sync.WaitGroup{}`.
- The WaitGroup needs the **sync** package. A typical skeleton code example of how to use a WaitGroup is:
```go
package main

import "sync"

func main() {
  wg := &sync.WaitGroup{}
  wg.Add(1)
  ...
  wg.Wait()
  ...
  ...
}
```
  - in the above code, first we create a pointer to a new WaitGroup.
  - then, we mention that we are adding an asynchronous operation that adds 1 to the group using **wg.Add(1)**.
    - this is a counter that holds the number of all the concurrent routines that are running.
  - we then add the code that will run the concurrent call.
  - at the end, we tell the WaitGroup to wait for Goroutines to end using **wg.Wait()**.
- How will the WaitGroup know that the routines have finished executing?
  - we tell the WaitGroup about it inside the Goroutine with this: **wg.Done()**.
  - this code must be placed inside the main Goroutine function because it needs a reference to the WaitGroup.
*see exercise 16.02*


## Race Conditions
- An important thing to consider is when we run mutiple functions concurrently, we have no guarantee in what order each instruction in each function will be executed.
- In some architectures, there are functions that are not connected in any way to other functions, and whatever the function does in a routine does not affect the actions performed in other routines.
- Wish life was like that in all scenarios BUT that is not the case.
- Suppose we have a parameter that is shared between some functions. Some of these functions will just read the parameter and others will write and modify this parameter.
- Let us explore this state with some sample code:
```go
package main

import (
  "fmt"
  "sync"
)

var counter int

func increment() {
  counter++
}

func main() {
  wg := sync.WaitGroup{}
  for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
      increment()
      wg.Done()
    }()
  }
  wg.Wait()
  fmt.Println(counter)
}
```
  - in the code above, we create 1000 goroutines that are trying to change the value of the counter at once.
  - here, there is no synchronization as to what function runs before another and the final value of the counter might not be 1000.
  - in this case where none of the routines are aware of what the others are doing, and they all override what one another is doing, it creates a **race condition**.
  - and **race conditions** happen everytime there is a shared resource that is not being worked with in sync.
- There are several ways through which we can mitigate race conditions, to make sure that the code changes are made only once.


## Atomic Operations
- Suppose we want to run independedn functions again BUT here, we want to modify the value held by the variable.
- Lets say we want to sum the numbers from 1 to 100 but we want to split the work between two concurrent Goroutines.
  - one goroutine will sum the numbers from 1 to 50 and the other will sum from 51 to 100.
  - these two routines have to update the same result variable 
- With the Goroutines manipulating the same variable, this can lead to *race conditions*.
  - A `race condition` happens when two processes change the same variable and one process overrides the changes made by another process without taking into account the previous change.
- There is a package called `atomic` which allows us tot safely modify variables across Goroutines.
- A look at a function that this package has: `func AddInt32(addr *int32, delta int32) (new int32)`.
  - this code takes a pointer to int32 and modifies it by adding the value it points at, to the value of delta. ie. if addr has value of 2 and the delta is 4, after calling this function, the add value will become 6.
*see exercise 16.03*


## Invisible Concurrency
- We can see the effects of concurrency through race conditions
- Concurrency problems are difficult to visualize as they do not manifest in the same way everytime the program is run.
- We want to test the concurrency patterns and visualize them. One simple way to do this is to print out each concurrent routine and see the order in which these routines are called. The downside of this is that, it is not easy to use these in testing.
- If we want to visualize the concurrency patterns and still be able to test the code, we use a func/struct in the **sync** package called **mutex**.
  - the **mutex** is short for **mutual exclusion** and this is essentially a way to stop all the routines, run the code in one single routine, and then carry on with the concurrent code.
- We create a mutex this way:
  - `mtx := sync.Mutex{}`.
  - most of the time we will be using mutexes, we will be using them across multiple functions. Hence, it is best practice to create a pointer to the mutex.
    - `mtx := &sync.Mutex{}`
    - this ensures that the same mutex is used everywhere. The reason why this has to be the way will be clear after analyzing the methods in the Mutex struct.
- Consider the following code:
```go
mtx.Lock()
s = s + 5
mtx.Unlock()
```
  - In the code above, we will lock the execution of the routines except the one needed to change the value of the variable **s**. After this, we release the lock so that any other routine can **safely** modify the value of **s**.
  - any code after the unlock will be ran concurrently.
  - **NB**: Try not to add more code between the Lock() and Unlock() constructs as this will tend to render your code less concurrent. There are better ways to ensure safety when modifying a variable.
  - In short, mutexes allow us to (1) lock the execution of the program, (2) add the logic required to ensure safety, (3) unlock, and then carry on with the execution of the rest of the program
- **NB**: Note that the order of asynchronously ordered code can change. Goroutines run idependently of each other and it cannot be predicted which one runs first. However, in this case, each routine runs completely before letting another run. Goroutines do not order things correctly and the final specific order that you want must be set by you.
*see activity 16.01*


## Channels
- we have seen how to create concurrent code with Goroutines, how to synchronize with waitGroup, how to perform atomic operations, and how to temporarily stop the concurrency in order to synchronize access to shared variables.
- There is a new concept to introduce, a **channel**.
- A channel is what the name suggests - it is somewhere the messages can be piped, and any routine can send or receive messages through a channel.
- Similar to slice, a channel is created this way:
```go
var ch chan int
ch = make(chan int)
// you can instantiate the channel directly with the following:
ch := make(chan int)
// just like with slices, we can do the following:
ch := make(chan int, 10) // a channel is created with a buffer of 10 items.
```
- A channel can be of any type, such as **integer, boolean, float, struct etc**, that can be defined, and even the slices and pointers, though the last two (floats and structs) are generally used less frequently.
- Channels can be passed as parameters to functions, and that is how different Goroutines can share content.
- Let us see how to send a message in the channel: `ch <- 2`.
- In this case, we can send value, 2, to the **ch** channel, which is a channel of integers. Trying to send something other than an integer will cause an error.
- After sending the message, we need to be able to receive a message from a channel. To do that, we do the following: `<- ch`
  - doing this ensures that the message is received; however, the message is not stored.
  - it might seem useless to lose the message but we will see that it might actually make sense.
- Nevertheless, we might want to keep the value received from the channel, adn we can do so by storing the value into a new variable.: `i := <- ch`.
- A simple program to showcase what we have learnt so far:
```go
package main

import (
  "log
)

func main() {
  ch := make(chan int, 1)
  ch <- 1 
  i := <- ch
  log.Println(i)
}
```
  - this program creates a new channel, pipes the integer 1 in, and then reads it and finally prints out the value of **i** which should be 1.
- If we make the channel unbuffered, ie `ch := make(chan int)`. In this case, we (and the compiler) do not know big the channel is. This makes the Goroutines  wait indefinitely and ends up in what is called a **deadlock**.
  - this does not mean that we cannot handle unbuffered channels. To handle them, they require more than one routine running. With only one routine, after we send the message, we block the execution and there is no other routine to receive the message. We then end up with a deadlock.
- A chanracteristic of a channel is that, it can be closed. Channels need to be closed when the task they have been created for is done.
  - in order to close a channel, do this: `close(ch)`. Alternatively, we can defer the closing of the channel:
  ```go
  defer close(ch)
  for i := 0; i <= 100; i++ {
    ch <- i
  }
  return
  ```