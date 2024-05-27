# Complex Types
- Go has complex types that are used to model complex data.
- These complex types help to group data and logic to accurately reflect real-world solutions being built.
- These core type combinations to model data are grouped into `collections` and `structs`.
  - `structs` are used to compose a type made from/of other types and associate logic with them.
  - `collections` allow the user to group data, loop over them and perform operations on them.

## Collection Types
- Suppose you were dealing with 100s of emails. You cannot define a variable to hold each email for each account. 
  - when working with lots of similar data, we put it in a *collection*.
  - the Go collection types are *arrays*, *slices* and *maps*.
  - these are strongly typed and easy to loop over. They each have different use cases that make them more suitable than the other to different use cases.

  ### Arrays
  - Arrays are the most basic collection types.
  - when defining an array, you specify *(1)* the type of data that it contains and *(2)* the size of it / how many elements it may contain.
  - arrays are of the following form: `[<size>]<type>`. eg *[10]int* or *[5]string*.
  - what defines an array is the size limit imposed on it. When there is no size, it becomes another collection type called the *slice*.
  - you can initialize an array with data. Like this: `[<size>]<type>{<value>, <value>, <value>}`. eg. *[5]int{1, 2,3, 4, 5}*.
  - Best of all, you can have Go set the size of the array based on the number of elements you initialize it with. To do this, you use the `...` (triple dots). eg. *[...]]int{1, 2, 3, 4, 5}* will create an array of length 5,and this length is set at compile time and cannot be changed at runtime.

    #### Comparing Arrays
    - the length of an array is a part of its definition.
    - if two arrays that accept the same type but have varying lengths/sizes, they are not compatible and not comparable with each other.
    - arrays of different sizes/lengths and are not of the same type cannot be compared with each other too.

    #### Initializing Arrays with Keys
    - without explicitly setting the keys, Go automatically chooses the keys for us.
    - however we can also choose the key we want to associate with our data when we are setting them.
    - an example of this setting is: `[<size>]<type>{<key1>:<value1>...<keyN>:<valueN>}`.
    - Go is flexible and lets you set the keys with gaps in any order preferrable.
    - the initialization with keys is useful if you have an array where the keys have a specific meaning and you want to set a value for specific key but do not need to set any of the other values.
    - the key acts like an index where data can be placed. *see exercise4.03*.

    #### Reading from an Array
    - It is possible to read data from an array. Even better, we can access a single element from an array using the `<array>[<index>]`.
    - being able to access specific parts of an array can be helpful in many ways:
      - to validate the data in an array by checking the first/last elements.
      - to check the position of data in the array. Positional significance is used when reading comma-separated values (CSV) or other delimited value files.
    
    #### Writing to an Array
    - once an array has been defined, it is easy to make changes to the individual components/elements using their index with `<array>[<index>] = <value>`.
    - in real-world code, you need to modify the data in collections after it has been defined based on input and logic.

    #### Looping an Array
    - the most common way to work with arrays is by using them in loops.
    - this works mostly due to the fact that arrays have indexes which make them easier to loop over.
    - using loops with arrays allow us to create variables to store indexes whcih can then be incremented manually using the `for i` loop.
    -a best practice with loops will be to NOT hardcode the length of an array when using it in a loop. Use the `len` keyword to get the final index of the array. Hardcoding makes code harder to maintain, and introduces hard-to-find bugs and even runtime panics.
    - using loops with arrays allow for the same logic to be used with every element, validating, modifying and outputting data without having to duplicate the same data for multiple variables.

    #### Modifying the Contents of an Array in a Loop
    - the contents of an array can also be changed using a loop.
    - working with the data in the array / each element is like working with variables.
    - modifying the content with loops reduces the code written if each element was a standalone variable.

  ### Slice
  - slices are arrays but without the rigidity or size.
  - arrays are efficient when it comes to managing sorted collections of data. Slices take it up an notch by adding flexibility to this efficiency.
  - a slice is then a thin layer around arrays that let you have a sorted numeric indexed collection without having to worry about the size.
    - in the thin layer, Go manages all th details relating to size.
    - slices can be used just like arrays: holds value of one type, can read and write to each element, and are easy to loop over.
  - a slice can be expanded using the `append` function. This one accepts the values you like to add and returns a new slice with everything merged.
  - in real-world code, slices are the go-to for all sorted collections.

    #### Appending Multiple Items to Slices
    - the `append` function can add more that one value to a slice. Many parameters as needed can be added using `append` since the last parameter is variadic.
    - because it is variadic, the `...` notation can be used to to USE the slice as a variadic parameter allowing for the passage of a dynamic number of parameters to `append`.
    - being able to pass more than one parameter to append comes up in real-world Go code. It keeps code compact by not requiring multiple calls or loops to add multiple values.
    
    #### Creating Slices from Slices and Arrays
    - using the notation used to access a single element in an arrays or slice, new slices can be created from slices and arrays.
    - the most common notation is the `[<low>:<high>]`. This tells Go to create a slice with the same value type as the source (slice/array) and to populate it with new values starting at the low index indicate to (but not including) the high index.
      - the low and high indexes are optional. If low is ommitted, Go defaults to starting from the first value. If high is ommitted, the values up to the last value are used. If you skip both, all the values in the source are populated.
      - Go does NOT copy these values.
      - if the source is an array, then that source array is the hidden array for the new slice. If the source is a slice, then the hidden array for the new slice is the same hidden array that the source slice uses.
    

    - the **make** function can be used to create a slice and set its length and capacity. The syntax looks like this `make(<sliceType>, <length>, <capacity>)`.
      - in creating a slice with **make**, the capacity is optional but the length is required.

  ### Map Fundamentals
  - a map in Go is a hashmap in compsci terms (wth! is a hashmap)
  - a map has a key and a value, much like all the other collection types. The key in a map is data and this data has a real relationship with the value.
  - if you have some data pairing a key and value, you will not have to loop over data when trying to pinpoint a particular key/value.
  - with a map, you can get, set and delete data quickly.
  - you can access the individual elements of a map using **[]**.
  - for a key, the maps have to have a type that is comparable directly. ie keys can be ints, strings ...etc.
  - for a value, in maps, they can be of any type. Including pointers, slices and maps.
  - if you want an ordered list, do not use a map
  - to define a map, use the following notation `map[key_type]<value_type>`. You can also use **make** to create maps but the arguments to use with the make are different from the ones used to create a slice.
  - Go does not help to manage your keys with maps AND you have to specify keys when initializing maps with data.
  - the notation for initializing maps with data are `map[key_type]<value_type>{<key1>:<value1>, ... <keyN>:<valueN>}`.
  - this is how to set a value in a map; `<map>[<key>] = <value>`.
  - to set a value for a map, make sure to have initialized the map first to avoid runtime panic.
  - initialize a map with data, or use **make** to create a map. DO NOT use var to define a map.
  - when to use a map depends on the kinds of data you want to store in it and if the access pattern needs access to individual items rather than a list of items.

    #### Reading from Maps
    - It is not always known if a key exists in a map before needing to use it to get a value.
    - Go returns the zero value for a value type whose key does not exists in the map. Having logic that works with zero values is the right way to program in Go.
    - if the zero value logic cannot be used, maps have an extra return value in the **exists** boolean.
    - the notation for this extra return value is `<value>, <exists_value> := <map>[<key>]`.
      - here the **exists** boolean value is true if a key exists in a map, and false if otherwise.
    - when looping over a map, use the **range** keyword. Go randomizes the order of the elements in a map.
    - to loop over the elements of a map in a specific order, you need to use an array or a slice for that.
  
    #### Deleting Elements from a Map
    - With an array, you cannot remove/delete an element from the array since the length is fixed, the best you can do is to zero out the value.
    - For a slice, you can zero out a value and you can also use a combination of slice ranges to cut and append values (essentially "deleting" a value).
    - With maps, the best thing wil be to use the in-built **delete** function. The notation to use is `delete(<map>, <key>)`. The delete function does not return anything, and if a key does not exist, nothing happens.
    - the **delete** can only be used for maps, NOT slices and arrays.

  ### Simple Custom Types
  - Custom types can be created using simple types as a starting point.
  - the notation is `type <name> <type>`. An example type is `type id string`.
  - the custom type acts as the type it is based upon. Getting the same zero values and the ability to comare with other values of the same type.
  - a custom type is not compatible with its base type but the custom type can be converted back into the type it is based upon to allow for interactions.

  ### Structs
  - Collections are perfect for grouping data of the same type and purppose together.
  - Go has structs, a customm type that is used to specify different field properties and their types.
  - The struct type is used to store data that has different types but have the same purpose and ought to be together.
  - the notation for a struct is as follows:
  ```go
  type <name> struct {
    <fieldName1> <type>
    <fieldName2> <type>
    ...
    <fieldNameN> <type>
  }
  ```
  - Field names must be unique within a struct. Any type can be used for a field; including pointer, collections and other structs.
  - to access a field on a struct, use the following notation: **`<structValue>.<fieldName>`**.
  - to set a value in a struct, use the notation: **`<structValue>.<fieldName> = <value>`**.
  - to read a value from a struct, use the notation: **`value = <structValue>.<fieldName>`**.
  - structs are the closest thing to what other languages call classes. They have been stripped down and do not have any inheritance.
  - structs can be defined at the function scope too. If a struct type is defined in a function, it will only be valid for use in that function.
    - when defined at a package level, the struct is available for user throughout the code.
  - a struct can be defined and initialized at the same time. If you do this, you cannot reuse the type but it is still a useful technique.
  ```go
  type <name> struct {
    <fieldName1> <type>
    <fieldName2> <type>
    ...
    <fieldNameN> <type>
  }{
    <value1>,
    <value2>,
    ...
    <valueN>,
  }
  ```

    #### Comparing Structs to Each Other
    - If all of a structs fields are comparable types, then the struct as a whole is comparable.
    - ie. If the struct fields are made up of ints and strings, then you can compare whole structs to one another. If there is a slice in any field, the struct ceases to be comparable.
    - a bit of flexibiility here is that, if the struct was defined anonymouslyand it has the same structure as a named struct, then Go allows the comparison.

    #### Struct Composition using Embedding
    - an alternative to inheritance (which is not included in Go) is embedding ie. you can embed types in struct types.
    - what it means is that, you can add fields TO a struct FROM other structs.
    - the compostion feature lets you add TO a struct using other structs as components.
    - when you embed, the fields from the embedded struct act as if it has been defined on the target struct.
    - to embed a struct, you like to add it as a field but you do not specify a name. To do this, you add a struct type name to another struct without giving it a field name. Looks like this:
    ```go
    type <name> struct {
      <Type>
    }
    ```
    - although not common, you can embed any other type into structs. For this, there is nothing to promote, so to access the embedded type, you can access it using the type's name like this: `<structValue>.<type>`.
    - there are two(2) valid ways to work with an embedded struct's fields: `<structValue>.<fieldName>` and `<structValue>.<type>.<fieldName>`.
    - the ability to access the type by its name also means thatthey type's names must be unique between the embedded types and the root field names.
    - when embedding pointer types, the type's name is the type without the pointer notation so the name `*<type>` becomes `<type>`. The field remains a pointer, only the name is different.
    - when it comes to promotion, if you were to have any overlap with your struct's field names, Go allows you to embed but the promotion of the overlapping does not happen. You can still access the field by going through the type name path.
    - you cannot use promotion when initializing structs wuth embedded types. To initialize the data, you must use the embedded types name.
    - there is not much embedding in real-life code, Go developers prefer to have other structs as named fields.

    #### Type Conversions
    - there are times when your types will not match up, and with Go's strict type system, if they are not the same, they cannot interact with each other.
    - in this case, Go provides an option. If the two types are compatible, you can do a type conversion ie. you can create a new value by changing one type to another.
      - the notation for the type conversion is: **`<value>.(<type>)`**.
      - when working with strings, we can use this notation to cast a string to a slice of runes or bytes and back again. This works because a string is a special type that stores the string's data as a slice of bytes.
      - this string type conversion is lossless but that is not true of all type conversions.
    - Go does its best to guess at the types that need conversion. Eg `math.MaxInt8` is an int and if you try to assign it to a number other than an int, Go does an implicit type conversion for you.

    #### Type Assertions and interface{}
    - an interface in Go describes what functions a type must have to conform to that interface.
    - interfaces do not describe fields and they do not describe a type's core value (such as a string or number).
    - all types in Go conform to **interface{}** or in this case **any**.
    - the notation for type assertion is **`<value>.(<type>)`**.
    - type assertion results in a valut of the type that was requested and optionally a bool regarding whether it was successful or not.
    - this looks like this **`<value> := <value>.(<type>)`** or **`<value>, ok := <value>.(<type>)`**. If ypue leave the boolean out, the type assertion fails and Go raises a panic.
    - a combinaion of interface{} and type assertions allow you to overcome Go's strict type coctrols, in turn allowing you to create functions that can work with any type of variable. You lose the protection that Go gives you at compile time for type safety

    #### Typw Switch
    - if we wanted to expand the doubler function we write in `exercise21` to include all int types, we would end up with a lot of duplicated logic.
    - Go has a way of dealing with more complex type assertion situations, known as the **type switch**.
    - here is the notation:
    ```go
    switch <value> := <value>.<type> {
      case <type>:
        <statement>
      case <type>, <type>:
        <statement>
      default:
        <statement>
    }
    ```
    - the type switch only runs the logic if it matches the type you are looking for, and it sets the value to that type
    - you can match on more than one type in a case, but Go cannot change the type of the value for you, so you will still need to do type assertion.