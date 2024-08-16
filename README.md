# Golang Study

### Why Go?

- What was happening at Google...
- Created by Ken Thompson (Unix, B, C), Rob Pike (UTF-8), and Robert Griesemer.
- In 2006, there was no language that was fast to compile, fast to execute, and easy to program. It is a language created to solve performance and complexity issues.
- https://golang.org/doc/faq#creating_a_new_language
- Efficient
 - Standard library is available
 - Multiplataform.
 - Garbage collection (lightning fast!)
 - Cross-compile.
- Easy to use
 - It is a compiled language, with strong and static typing,
 - There are very few reserved words, which we will all learn in the course, in other words, they are very good to learn
 - only goes up in popularity lists
- Killer feature: In 2006, right after the first dual core. Thread: 1mb. Goroutine: 2kb.
- It's awesome!
- When to use Go?
 - Scale
 - Web services, networks, servers (machine learning, image processing, crypto, ...)
 - When you need a fast, simple, easy-to-learn, and easy-to-use language.
- Used in: APIs, CLIs, microservices, libraries/framework, data processing, ... It is the basis of cloud services and container orchestration.
- Who use? https://github.com/golang/go/wiki/GoUsers
- Is Go OOP? https://golang.org/doc/faq#Is_Go_an_object-oriented_language
- One more: https://golang.org/doc/faq#principles

### Hello world!

- := looks like a groundhog (gopher) or the punisher.
- Usage:
 - Automatic typing
 - Can only be repeated if there are new variables
 - != do assignment operator
 - Only works within codeblocks
- Terminology:
 - keywords are reserved terms
 - operators, operands
 - statement (declaration, statement) → a line of code, an instruction that forms an action, made up of expressions
 - expression → anything that "produces a result"
 - scope
 - package-level scope
- Key lesson:
 - := used to create new variables, within code blocks
 - = to assign values ​​to existing variables

 ### Exploring types

- Types in Go are extremely important. (We'll see more when we get to methods and interfaces.)
- Types in Go are static.
- When declaring a variable to contain values ​​of a certain type, that variable can only contain values ​​of that type.
- The type can be deduced by the compiler:
 - x := 10
 - var y = "batima's aunt"
- Or it can be specifically stated:
 - var w string = "this is a string"
 - var z int = 15
 - in var z int declaration with package scope, assignment z = 15 in codeblock (only)
- Primitive data types: available in the language natively as basic building blocks
 - int, string, bool
- Composite data types: these are types composed of primitive types, and created by the user
 - slice, array, struct, map
- The act of defining, creating, structuring compound types is called composition. We will see a lot of this in the future.

### Zero value

- Declaration vs. startup vs. value assignment. Variables: mailboxes.
- What is zero value?
- The zeros:
 - ints: 0
 - floats: 0.0
 - booleans: false
 - strings: ""
 - pointers, functions, interfaces, slices, channels, maps: nil
- Use := whenever possible.
- Use var for package-level scope.

### Creating your own type

- Reviewing: types in Go are extremely important. (We'll see more when we get to methods and interfaces.)
-There's a story that Bill Kennedy said that if one day she got a tattoo, she would say "type is life."
- Most of the more advanced aspects of Go depend almost exclusively on types.
- As a foundation for these tools, we will learn how to declare our own types.
- Reviewing: types are fixed. Once a variable is declared to be of a certain type, it is immutable.
- type hotdog int → var b hotdog (main hotdog)
- A variable of type hotdog cannot be assigned the value of a variable of type int, even if this is the underlying type of hotdog.

### Conversion, not coercion

- Type conversion is what it sounds like.
- In Go we don't say casting, we say conversion.
- a = int(b)
- ref/spec#Conversions
- End of session. Congratulations! Tips, motivation and exercises.

### Boolean type

- Now let's explore the types in more detail. golang.org/ref/spec. Starting with bool.
- The bool type is a binary type, which can only contain one of two values: true and false. (True or false, yes or no, zero or one, etc.)
- Whenever you see relational operators (==, <=, >=, !=, <, >), the result of the expression will be a Boolean value.
- Booleans are fundamental in decision making in conditional logic, switch statements, if statements, control flow, etc.

 ### Numeric types

- int vs. float: Integers vs. numbers with fractions.
- golang.org/ref/spec → numeric types
- Integers:
 - Integers
 - int & uint → “implementation-specific sizes”
 - All numeric types are distinct, except:
 - byte = uint8
 - rune = int32 (UTF8)
 (The Go language source code is always in UTF-8).
 - Types are unique
 - Go is a static language
 - int and int32 are not the same thing
 - To "mix" them, conversion is necessary
 - General rule: use only int
- Floating point:
 - Rational or real numbers
 - General rule: only use float64

### String type (character strings)

- Strings are sequences of bytes.
- Immutable.
- A string is a "slice of bytes" (or, in Portuguese, a slice of bytes).
- In practice:
 - %v %T
 - Raw string literals
 - Conversion to slice of bytes: []byte(x)
 - %#U, %#x

### Number systems

- Base-10: decimal, 0–9
- Base-2: binary, 0–1
- Base-16: hexadecimal, 0–f

### Constants

- These are immutable values.
- They can be typed or not:
 - const hi = "Good morning"
 - const hi string = "Good morning"
- Untyped ones will only have a type assigned to them when they are used.
 - Ex. what type of 42? int? uint? float64?
 - In other words, it is convenient flexibility.
- In practice: int, float, string.
 - const x = y
 -const(x=y)

### Iota

- golang.org/ref/spec
- In a constant declaration, the iota identifier represents sequential numbers.
- In practice.
 - iota, iota + 1, a = iota b c, restarts on each const, _

### Bit-Shift

- Bit shifting is when we shift binary digits to the left or right.
- In practice:
 - %d %b
 - x << y
 - iota * 10 << 10 = kb, mb, gb

 ### Understanding control flow

- Computers read programs in a certain way, in the same way that we read books, for example, in a certain way.
- When we Westerners read books, we read from front to back, from left to right, from top to bottom.
- Computers read from top to bottom.
- In other words, its reading is sequential. This is called sequential control flow.
- In addition to sequential control flow, there are two statements that can affect how the computer reads code:
 - One of them is the loop control flow. In this case, the computer will repeat reading the same code in a specific way. Repetitive control flow is also known as iterative control flow.
 - And the other is the conditional control flow, or selection control flow. In this case, the computer finds a condition and, through an if or switch statement, takes one course or another depending on that condition.
- In other words, there are three types of control flow: sequential, repetition and conditional.

 - Sequential
 - Iterative (loop)
 - for: initialization, condition, post
 - for: hierarchically
 - for: condition ("while")
 - for: ...ever?
 -for:break
 - for: continue
 - Conditional
 - switch/case/default statements
 - there is no fall-through by default
 - creating fall-through
 - default
 - multiple cases
 - cases can be expressions
 - if they result in true, run
 - type
 -if
 - bool
 - the "!" operator
 - initialization declaration
 - if, else
 - if, else if, else
 - if, else if, else if, ..., else

### Loops: initialization, condition, post

-For
 - Startup, condition, post
 - Semicolon?
 - gobyexample.com
 - There is no while!

### Loops: nested loop (hierarchical repetition)

-For
 - Hierarchical repetition
 - Examples: clock, calendar

### Loops: the for statement

- For: initialization, condition, post
- For: condition ("while")
- For: ...ever? (httpservers)
- For: break
- golang.org/ref/spec#For_statements, Effective Go
- (Range comes further forward.)

### Loops: break & continue

- Module operation: %
- For: break
- For: continue

### Conditionals: the if statement

- If: bool
- If: the operator does not → "!"
- If: initialization statement

### Conditionals: if, else if, else

- If, else.
- If, else if, else.
- If, else if, else if, ..., else.

### Conditionals: the switch statement

- Switch:
 - can evaluate an expression
 - switch statement == case (value)
 - default switch statement == true (bool)
 - there is no fall-through by default
 - creating fall-through
 - default
 - composite cases

### Conditional logical operators

- &&
- ||
- !
- What is the result of fmt.Println...
 - true && true
 - true && false
 - true || true
 - true || false
 - !true

 ### Array

- Data structures, or data groups, allow us to group different values. These values ​​may or may not be of the same type.
- The structures we will see are: arrays, slices, structs and maps.
- Let's start with arrays. Arrays in Go are a foundation, and not something we use every day.
- Their size must be present in the declaration: var x [n]int
- Values ​​are assigned to their positions with: x[i] = y (0-based)
- To see the size, use: len(x)
- ref/spec: "The length is part of the array's type" → [5]int != [6]int
- Effective Go: Arrays are useful for [some things we will never do] and serve as a foundation for slices. Use slices instead of arrays.

### Slice: Composite literal

- What are composite data types?
- Wikipedia: Composite_data_type
- Effective Go: Composite literals
- ref/spec: Composite literals
- A slice groups values ​​of a single type.
- Creating a slice: composite literal → x := []type{values}

### Maps:

- Uses the key:value format.
- E.g. name and phone
- Excellent performance for lookups.
- map[key]value{ key: value }
- Access: m[key]
- Key without value returns zero. This can cause problems.
- To check: comma ok idiom.
- v, ok := m[key]
- ok is a boolean, true/false
- In practice: if v, ok := m[key]; ok { }
- To add an item: m[v] = value
- Maps *have no order.*

### Struct

- Struct is a composite data type that allows us to store values ​​of *different* types.
- Its name comes from "structure," or structure.
- Declaration: type x struct { y: z }
- Access: x.y

### Embedded structs

- Structs inside structs inside structs.
- Example: a Formula 1 driver is a person (name, surname, age) *and also* a competitor (name, team, points).

### Reading the documentation

- It is important to familiarize yourself with the Go language documentation.
- In this video we will see a little about what the documentation says about structs.
- We will see:
- ref/spec
- We have already seen more than half of the types in Go!
- Struct types.
- x, y int
- anonymous fields
- promoted fields

### Anonymous structs

- These are structs without identifiers.
- x := struct { name type }{ name: value }

### Functions

- What are functions useful for?
- Abstracting functionality
- Code reuse
- func (receiver) identifier(parameters) (returns) { code }
- The difference between parameters and arguments:
- Functions are defined with parameters
- Functions are called with arguments
- Everything in Go is *pass by value.*
- Pass by reference, pass by copy, ... no.
- Parameters can be ... variadic.

### Defer

- Functions are great because they make our code modular. We can change parts of our program without affecting the rest!
- A defer statement calls a function whose execution will occur when the function it is part of finishes.
- This finish can occur due to a return, at the end of the function's code block, or in the case of a panic in a corresponding goroutine.
- "Leave it for the last minute!"
- ref/spec
- We always use it to close a file after opening it.

### Methods

- A method is a function attached to a type.
- When you attach a function to a type, it becomes a method of that type.
- You can attach a function to a type using its receiver.
- Usage: value.method()
- Example: the type "person" can have a method goodmorning()

### Interfaces & polymorphism

- In Go, values ​​can have more than one type.
- An interface allows a value to have more than one type.
- Declaration: keyword identifier type → type x interface
- After declaring the interface, you must define the methods needed to implement that interface.
- If a type has all the required methods (which, in the case of interface{}, can be none) then that type implicitly implements the interface.
- That type will be your type *and also* the type of the interface.

### Pointers

- All values ​​are stored in memory.
- Every location in memory has an address.
- A pointer refers to that address.
- Notations:
- &variable shows the address of a variable
- %T: variable vs. &variable
- *variable de-references, shows the value that is at that address
- ????: *&var works!
- *type is a type that contains the address of a value of type type, in this case * is not an operator

### When to use pointers

- Pointers allow you to share memory addresses. This is useful when:
- We don't want to pass large amounts of data back and forth
- We want to change a value in its original location (everything in Go is pass by value!)

### Generics

- Starting with version 1.18, Go has added support for generics, also known as type parameters.

- As an example of a generic function, MapKeys takes a map of any type and returns a slice of its keys. This function has two type parameters - K and V; K has the comparable constraint, meaning that we can compare values of this type with the == and != operators. This is required for map keys in Go. V has the any constraint, meaning that it’s not restricted in any way (any is an alias for interface{}).
```
func MapKeys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}
```
- As an example of a generic type, List is a singly-linked list with values of any type.
```
type List[T any] struct {
    head, tail *element[T]
}
type element[T any] struct {
    next *element[T]
    val  T
}
```
- We can define methods on generic types just like we do on regular types, but we have to keep the type parameters in place. The type is List[T], not List.

### Packages

- Packages are the basic unit of code organization in Go. A package is a collection of Go files that are organized into a specific directory. All files in the same directory belong to the same package.

- Each Go file begins with a package declaration, such as package main or package fmt. The package declaration defines the name of the package to which the file belongs.

- Packages allow you to group related functions, types, and variables. This makes it easier to reuse code and improves project organization.

- The main package is special because it defines the entry point of a Go program. A Go project that is intended to be run as a program must have a main package with a main() function.

### Modules

- Modules are a way to organize and manage dependencies in Go projects. A module is a collection of packages that are versioned together. A module is defined by a go.mod file in the root of the project directory.

- The go.mod file specifies the name of the module, as well as the external dependencies it uses. This makes version management and dependency resolution easier.

- With modules, you can import packages from other projects (both local and remote) and version control their dependencies.

### Compiling Project

#### 1. Compiling a Single Go File

- If you have a single Go file (e.g. `main.go`), you can compile it directly using the `go build` command.

- Steps:
1. Navigate to the directory where your `main.go` file is located.

2. Run the command:

```
go build
```

- This will generate an executable in the same directory with the name of the Go file (e.g. `main.exe` on Windows or `main` on Linux/Mac).

#### 3. To run the generated binary, use:

```
./main
```

- or on Windows

```
main.exe
```

#### 2. Building a Project with Multiple Files/Packages

- If your project has multiple Go files or is organized into packages, you can use the same `go build` command, but with some additional considerations.

- Steps:
1. Make sure you are in the root directory of your project, where the `go.mod` file is located (if the project is modular).

2. Run the command:

```
go build
```

- Go will build the entire project, including all Go packages and files. This will generate a single binary in the root directory.

3. To run the generated binary, use the same procedure mentioned above.

#### 3. Compiling with Binary Name Specification

- You can specify the name of the generated binary using the -o flag.

Example:

```
go build -o my_program
```
- This will create a binary named my_program instead of using the default name.

#### 4. Cross Compilation
- Go allows you to compile your code for different operating systems and CPU architectures.

 Example:

- To compile for Linux on a Windows machine, you can use:

```
GOOS=windows GOARCH=amd64 go build -o my_program_windows.exe
```

#### 5. Running the Project Without Compiling Beforehand

- If you just want to run the project without generating a persistent binary, use the command:

```
go run main.go
```
or
```
go run .
```

- This command compiles and executes the code in a single step.

