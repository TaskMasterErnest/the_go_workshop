## Value vs Pointer
Values and Pointers are ways through which we store 'values' to be used in code. Values like *int*, *bool* and *string* are passed into functions to be used.

When these values are passed into functions, what Go does is to make a copy of the value and the copy is what is used by/in the function.
With this, any change that happens to the value withing the function does not affect the original value from which it was copied.

Apparently, passing values by copying them tends to make code have fewer bugs (I do not know which!). 
Using the copy way, Go can use its simple memory management system called the stack (read more on this...).
The downside to this is that, when the values are being passed between a lot of functions, the copying is done multiple times and this causes the system to use more memory to handle the copying.

The alternative to copying that uses less memory is to use something called a *pointer* and pass it to functions.
- a pointer in itself is not a value, it specifies the location of the value.
- in using a pointer, Go does not make a copy of the value at the pointer, it just directs the function to go use that value that is in that location.
- Go manages the pointer on the *heap*. The heap allows the value to exist until no part of the application/software has a pointer to that value anymore.
- Go reclaims these pointer in its garbage-collection process
- A value can be put on the heap, when a pointer is made to it...but ultimately it it up to the dev to work out whether a value should be put on the heap or not. This is called *escape analysis*.
- A pointer has a zero value of 'nil' which can cause a runtime error. The best approach to using pointers will now be to compare the pointer to 'nil' before trying to get its value.

### Getting a Pointer
There are a few options to declare a pointer variable.
1. declare using the **var** statement and add an * (asterisk) at the front of most types. Like this: `var <name> *<type>`.
2. using the built-in function **new**. This is used to get some memory for a type and return a pointer to that address. Like this: `<name> := new(<type>)`. This new function can be used with the var too.
3. you can also get a pointer from an existing variable using **&**. Like this: `<var1> := &<var2>`.

### Getting a Value from a Pointer
To get the value a pointer is associated with, you have to dereference the value using * (asterisk) in front of the variable name. Like this: `fmt.Println(*<val>)`.
- it is best practice to check that a pointer is not 'nil' before dereferencing.
- you do not need to dereference a function/property that is on a struct.

### Using Constants
Constants are values that do not need to change during the runtime of an application/code.
- they are defined by the `const` keyword.
- use constants to set values and refer to them to avoid having to track and modify varibales throughout the code when debugging.
- constants that do not have a explicit type stated are inferred from the value given.
- constants can be declared multiple times in one statement. eg `const <name> <type> = <value>`.
```go
const (
  <name> <type> = <value>
  <nameN> <typeN> = <valueN>
)
```

### Scopes
Scopes basically can de defined as contexts, a way in which something is viewed.
- variables in Go live in a scope. The top-level/highest-level scope is the package level scope. A scpe can have child scopes in it.
- you can see this in code that starts and ends with matching `{}`. Everythin in a pair of these is a scope. The parent scope houses the child scope and so on.
- the parent-child relationship is defined when the code compiles and NOT when it runs.
- when accessing a variable, Go looks in the scope it ws defined it. If it cannot find it there, it looks in the parent scope, then the grand-father scope and so on till it reaches the package scope.
- it stops looking once it finds a variable with a matching name or raises an error if none is found. It raises an error is the variable is found but the type does not match.