# About Time
In this session, we will be able to create our own time format, compare and manage time, calculate the duration of time series and format time according to user requirements.

## Making Time
- Making time means declaring a variable to hold the time formatted in a specific way.
- For the introduction, we will be using the default time formatting provided by Go.
- We will have a example that wil show the how to measure the duration of the execution of a script.
  - we can do this by capturing the current time in a variable; at the beginning and at the end, and then we find the difference and know how long it took the specific action to complete.
  ```go
  import (
    "fmt"
    "time"
  )

  func main() {
    start := time.Now()
    fmt.Println("The script has started at: ", start)
    fmt.Println("Saving the World...")
    time.Sleep(2 * time.Second)
    end := time.Now()
    fmt.Println("The script has completed at: ", end)
  }
  ```
  - the output of this is not sexy, but we will learn to make it more readable.
- Consider a scenario where an application version is release every Monday at 12:00pm CEST, with a downtime window of 2 hours; 30 for deployment and 1.5hr for testing. You have to perform a simple hit-and-ruun test on any other day but on release day, there should be a full functional and integration test.
  - the code for this can be:
  ```go
  import (
    "time"
    "fmt"
  )

  func main() {
    Day := time.Now().Weekday()
    Hour := time.Now().Hour()
    
    fmt.Printf("Day: %v\nHour: %v", Day, Hour)

    if Day.String() == "Monday" {
      if Hour >= 1 {
        fmt.Println("Performing full blown testing")
      }
    } else {
      fmt.Println("Performing hit-and-run testing")
    }
  }
  ```
- Another scenario is to create the log filenames for scripts in Go. The goal is to have the log output and a timestamp concatenated to the name of the log file. Looks like this *Application_Action_Year_Month_Day*.
  - the ELEGANT and SIMPLE way to do this is:
  ```go
  import (
    "strconv"
    "time"
  )

  var (
    AppName = "HTTPCHECKER"
    Action = "BASIC"
    Date = time.Now()
    LogFileName = AppName + "_" + Action + "_" + strconv.Itoa(Date.Year()) + "_" + Date.Month().String() + "_" +
      strconv.Itoa(Date.Day()) + ".log"
  )

  func main() {
    fmt.Println("The name of the logfile is: ", LogFileName)
  }
  ```
- **NOTE**: Any operating system you work with will provide two types of clocks to measure the time; one is called the "monotonic" clock and the other called the "wall" clock. The wall clock is subject to change and is mostly what is seen on the taskbar of a Windows machine; it is subject to change and is usually synchronized with a public or corporate NTP (Network Time Protocol) server based on your location. The NTP is used to tell clients the time based off on an atomic clock or from a satellite reference.

## Comparing Time
- When working with Go on smaller scripts, it is important, for your statistics, to know when a script should run or between what hours and minutes a script should be completed.
  - by statistics, it is the time difference between automatically running a script and manually running one.
  - this will allow the measurement of improvement of the script over time when we develop the functionality further.
- Take a look at this script, which has been intended to NOT run before or after a specified time; the time can arrive via another automation or after a trigger file is placed in there manually. The script needs to run everyday, at different times, specifically after the specified time as soon as possible.
- Here is the code:
```go
import (
  "time"
  "fmt"
)

func main() {
  now := time.Now()
  only_after, _ := time.Parse(time.RFC3339, "2019-09-27T22:08:41+00:00")
  
  fmt.Println(now, only_after)
  fmt.Println(now.After(only_after))

  if now.After(only_after) {
    fmt.Println("Executing actions!")
  } else {
    fmt.Println("Now is not the time yet!")
  }
}
```
  - We have the **now** variable, which is crucial for the execution.
  - We have the **time** string parsed based on RFC3339. RFC3339 specifies the format that should be used for the **time** and **date** strings.
  - the function returns two values, one value is the output if the conversion succeeds, and the other is an error if there is one.
  - the **time** package has two useful functions; the **After()** and the **Before()**. There is also a third cone called **Equal()**. All these help us to simply compare the time variables. In the code above, we are able to use the **Before()** and **After()** to compare the times we have set in the variables.
- Another simulation:
```go
import (
  "time"
)

func main() {
  now := time.Now()
  now_too := now
  // simulate a sleep time
  time.Sleep(2 * time.Second)
  later := time.Now()

  if now.Equal(now_too) {
    fmt.Println("the times are the same!")
  } else {
    fmt.Ptintln("the times are different!")
  }
}
```

## Duration Calculation
- It is forever handy to calculate the duration it takes for a program to execute in the world of programming.
- The duration calculation can be used to monitor discrepancies and performance bottlenecks from the infrastructure used. Eg. if a script takes a certain time to execute and there is a huge bump in time as shown by a monitoring software, we are sure to investigate that.
  - another aspect where you might want to use this is in web applications, to measure the duration of the request-response of the applications to see how well they handle high loads and even gain insight on how to provision capacity at certain times.
- To add such a functionality, duration calculation, there is little coding involved. Here is an example:
```go
func main() {
  start := time.Now()
  fmt.Println("The script started at:", start)

  sum := 0
  for i := 0; i < 10000000000; i++ {
    sum += i
  }

  end := time.Now()
  duration := end.Sub(start)
  
  fmt.Println("The script completed at:", end)
  fmt.Println("The task took", duration.Hours(), "hour(s) to complete.")
  fmt.Println("The task took", duration.Minutes(), "minute(s) to complete.")
  fmt.Println("The task took", duration.Seconds(), "second(s) to complete.")
  fmt.Println("The task took", duration.Nanoseconds(), "nanosecond(s) to complete")
}
```
- For measuring duration calculation, there are six(6) different time resolutions to choose from:
  - hours
  - minutes
  - seconds
  - milliseconds
  - microseconds
  - nanoseconds
- Let us suppose we have an SRE job and a Service Level Agreement that needed to be met and let us assume it was stated that there were applications that needed to process a request in 1000ms or 5s depending on the criticality of the product. Let us look at some code to illustrate this.
```go
func main() {
  // set a deadline of 6s for request processing
  deadline_seconds := time.Duration(600*10) * time.Millisecond
  start := time.Now()
  fmt.Println("The deadline for the transaction is:", deadline_seconds)
  fmt.Println("The transaction started at:", start)

  sum := 0
  for i := 0; i < 25000000000; i++ {
    sum += i
  }

  end := time.Now()
  // duration := time.Duration(end.Sub(start).Seconds()) * time.Second)
  duration := end.Sub(start)

  transactionTime := time.Duration(Duration.Nanoseconds()) * time.Nanosecond
  fmt.Println("The transaction has completed at:", end, duration)
  
  if transactionTime <= deadline_seconds {
    fmt.Println("Performance is OK! Transaction completed in:", transactionTime)
  } else {
    fmt.Println("Performance problem! Transaction completed in:", transactionTime, "second(s)!")
  }
}
```

## Managing Time
- The Go programming language provides two functions that are used to manage time, **Add()** and **Sub()**.
- They are mostly used to calculate elapsed time; **Sub()** for the difference in elapsed time, **Add()** for the adding more time to elapse.
```go
  TimeToManipulate := time.Now()
  ToBeAdded := time.Duration(10 * time.Second)
  fmt.Println("The original time:",TimeToManipulate)
  fmt.Println(ToBeAdded," duration later:",TimeToManipulate.Add(ToBeAdded))
```
- To add more time, we can simply use the **Add()** function and have the new time show up in the console. For subtracting time,. it is a little more cumbersome as the time will not be removed/subtracted from the time we have. The best way is to craft the duration to remove as a negative value. Something like this: `toBeRemoved := time.Duration(-10 * time.Minute)`. This will help output a time that was 10 minutes before the current time.

## Formatting Time
- When formatting time in Go, the dates mostly are looking horrible and we want to *format* them.
- In formatting time, there are two main concepts we focus on:
  - the first, for instances when we would like the date to be ouputted in a desired looking string when we use it in print.
  - the second option is for when we would like to take a string and parse it to a specific format
- First, let us look at the **Parse()** function. This function takes in two arguments. The first argument is the standard to parse against and the second argument is the string that needs to be parsed.
  - the end of this parse will result in a time variable that can utilize built-in Go functions. Go uses a POSIX-based date format
  - **Parse()** is useful when you have an application that is working with time values from different time zones and you would like to convert them, for example, to the same time zone for better understanding and easier comparison.
  - There are three(3) main standards against which we can parse the time:
    - RFC3339
    - UnixDate : this expects an epoch, An epoch is a unique integer that represents the number of seconds that have passed since January 1, 1970. This date is the beginning of time on UNIX systems.
    - ANSIC
  - So when you provide the right format of the date, the **Parse()** function will provide the correct output.
  ```go
  func main() {
    t1, _ := time.Parse(time.RFC3339, "2019-09-27T22:18:11+00:00")
    t2, _ := time.Parse(time.UnixDate, "2019-09-27T22:18:11+00:00")
    t3, _ := time.Parse(time.ANSIC, "2019-09-27T22:18:11+00:00")
    fmt.Println("RFC3339: ", t1)
    fmt.Println("UnixDate:", t2)
    fmt.Println("ANSIC:", t3)
  }
  ```
- Go allows the user to craft their own time variables.
```go
date := time.Date(2024, 7, 15, 11, 31, 34, 00000, time.UTC)
fmt.Println(date)
```
  - the preceeding code shows how to craft time for oneself. The values are explained by the skeleton syntax that follows:
  ```go
  func Date(year int, month Month, day, hour, min, sec, nanosec int, loc *Location) Time
  ```
  - Time zones are important now that enterprises have fleets of interconnected devices and it makes it important to be able to differentiate between time zones.
- A neat trick to add a **Year, Month and Day** to the current time is to use the **AddDate()** function. This function provides a way to dynamically add to the dates present. For instance, let us add 1 year, 2 months and 3 days to the time we have above.
```go
date := time.Date(2024, 7, 15, 11, 31, 34, 00000, time.UTC)
next_date := date.AddDate(1, 2, 3)
fmt.Println(next_date)
```
  - this reveals that the **AddDate()** function receives three(3) arguments; Year, Month and Day.
- Another important aspect of time formatting is to understanc how to utilize the **LoadLocation()** function of the time package to your local time TO the local time of another TIME ZONE.
  - here, we can use the **Format()** function to tell Go how we will like to see the ouput formatted.
  - the **In()** function in the time package is very useful here, it is used to represent the specific time zone we want our formatting to outputted in.
  ```go
  current := time.Now()
  Berlin, _ := time.LoadLocation("Europe/Berlin")
  fmt.Println("The local current time is:", current.Format(time.ANSIC))
  fmt.Println("The time in Berlin is:", current.In(Berlin).Format(time.ANSIC))
  ```