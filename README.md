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