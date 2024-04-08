# Logic and Loops
Logic is a rapper!

- Lol, logic in code is the way you want your code to carry put a specific task.
- Loops are a way to continually use a particular logic multiple times.


### if statements
- the most basic form of logic in Go. 
- it has the notation, `if <boolean expression> { (code block) }`.
- the `if` statement will run or not, depending on the boolean expression used.
  - the boolean expression can be simple code tha results in a boolean value.
  - the code block can be any logic that can also be placed in a function.
- `if` statements can also be used in the function scope.


### if-else statements
- an `if-else` statement adds another logic to the `if` statement.
- it adds another block of code that runs only if the first `if` blovk does not run. Both blocks cannot run at the same time.
- the notation is this:
```go
if <boolean expression> {
  <code block>
} else {
  <another code block>
}
```

### else-if statements
- an `else-if` statement is used when the logic to be run has many choices to make before setting on a 'true' value.
- there is a interesting case in the way this block works with and without an `else` statement.
  - the first kind without an else block is like this:
  ```go
  if <boolean expression> {
    <code block>
  } else if {
    <code block>
  }
  ```
  - and this, with an else block:
  ```go
  if <boolean expression> {
    <code block>
  } else if {
    <code block>
  } else {
    <code block>
  }
  ```
- with the code without an else block, when Go iterates through the code and finds no true statement, it just exits the block and moves on. No code is ran. Conversely for the code with an else block.


### the initial if statement
- this procedure is used when the coder needs to call a function but does not care too much about the returned value.
- it means, the coder wants to just check if the code/function has executed successfully and then discard the returned value.
  - this method is used mainly for sending emails, writing to a file, inserting into a database; if these operations execute successfully, we do not need to worry about the values that they return.
- these variables do not go away but are still lingering in scope and we need to get rid of them.
- the best way to check for errors from these code executions is to use an initial `if` statement.
- the notation is like this:
```go
if <initial statement>; <boolean expression> {
  <code block>
}
// remember the ; that separates them.
```
- in the initial `if` statement, Go ony allows simple statements, these include:
  - assignment and short variable assignment ; `i := 0`.
  - expressions such as math or logic expressions ; `i = (j * 10) == 40`.
  - sending statements for working with channels
  - increment and decrement expressions ; `i++ or i--`
- A common MISTAKE is to use the `var` keyword here, this is NOT ALLOWED. Only the assignment method is used.


### expression switch statements
- this is the case where you have a lot of `if` and `else if` statements that make it hard to read at some point.
- Go has the `switch` statement that can be used in place of the `if` statements. The `switch` is a more compact alternative.
- the notation for the `switch` statement is as follows:
```go
switch <initial statement>; <expression> {
  case <expression>:
    <statements>
  case <expression2>:
    <statements>
  case <statement3>:
    <statements>
  default:
    <statements>
}
```
- both the `<initial statement>` and the `<expression>` are optional. They can be omitted, either `<expression>` is omitted or both are.
- when the value of the `<expression>` is omitted, it is as if the value `true` is placed in there.
- There are two main ways of using case expressions;
  - can be used like the initial if statements to control the logic of which statements get executed,
  - or, a value is put in the case expression and this value is compared to the switch expression. The statements only run when both expressions match.
- a `default` case can be added at the end of the switch statement. This functions like how the else statement runs.
- this form of switch statement is called an "expression switch" statement. There is another one also called "type switch".


### Loops
- loops are the codeo used when the same logic is going to be run repeatedly.
- Go has one looping statement, that is the `for` statement.
- there are two distinct forms of using this for loops:
  - the first is for ordered collections like arrays, slices etc
  - the second is for unordered data collections like maps etc
- For the first type for ordered collections, the loop statement looks like this:
```go
for <initial statement>; <condition>; <post statement> {
  <statements>
}
```
  - the initial statement runs like the ones we have in the if statements; runs before anything else and allows the same simple statements.
  - the conditions are checked to see whether the statements should be run or not. conditions also allow simple statements
  - post statements are run at the end of the loop; they also allow simple statements. The post statements is use for thinsgs like incrementing counters.
- the initial, condition and post statements are all optional to write, therefore it is possible to write a for loop like this:
```go
for {
  <statements>
}
```
  - this loop runs forever unless stopped manually using a `break` statement. There is also a `continue` statement that can be used to skip the remainder of a loop but does not stop the whole loop.
- another form of the for loop is used especially when reading from a source of data the returns a boolean when there is more data to read. eg, reading from databases, network sockets, files, CLI inputs etc...
  - it has the form:
  ```go
  for <condition> {
    <statements>
  }
  ```
- the second form is for loopin through unordered data collections such as maps. When looping over these, the `range` statement is used in the loop.
```go
for <key>, <value> := range <map> {
  <statements>
}
```

### the range loop
- the array and slice types always have an index for the elements, and that starts with 0.
- the map type is different. It uses a range instead of th condition in a for loop. The range yields a key and value for each element in the collection.
- this means that with range, you do not need a condition to stop the loop, range takes care of that.
- the format is described in `exercise2.10`.
- if you do not need a key or a value given, use `_` to tell the compiler that you do not want it.

### break and continue
- there will be times when one has to skip some loops or stop the loop from running altogether.
- the `continue` keyword stops the current loop and starts a new one. The post loop logic runs and and the loop condition gets evaluated.
- the `break` keyword stops the execution of the loop. No new loop is started.
- use continue when you want to skip a single item in a collection. eg. one of the items is invalid but the rest must be good to process.
- use break to stop processing when there are errors in the data and there is no value in processing the rest  of the collection.
- see `exercise2.11`.