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