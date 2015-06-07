[![BuildStatus](https://travis-ci.org/h8liu/e8vm.png?branch=master)](https://travis-ci.org/h8liu/e8vm)

# E8VM

**Paused development because I started working at Google, and Google has policies for open source hobby projects that I have to follow. I promise I will finish this when the time comes.**

Goal: A book written in working code and markdown document that
describes how computer system works, including architecture,
assemblers, compilers, and operating systems.

Read it [here](http://8k.lonnie.io).

Planned Features:

- **Modularized.** File based modules. No circular dependency (not only on packages,
  but also on files). A reader can always read the project a file by
  a file, either from bottom to top, or from top to bottom.
- **Small files.** Each file is shorter than 300 lines of code.
- **Tested and Documented.**
  Each file (will) come with test cases, examples, and markdown description.
- **Real.** The simulation (will) work like a real computer.

## TODOs

- Global `VarDecl`
- Type conversion
- Basic built-in panic
- Pointer
- Array and slice
- String
- Struct
- Fields and methods
- Interface					
- Big number constants
- Unused variable check
- Unreachable code check
- Missing return check
- Break, continue with labels

Small things:

- VarDecl ast printing

And more...

- Improve code reading website
- Complete consts in asm8
- Clean up the symbol linking in asm8 a little bit
- Package building system that tracks timestamps
- Online filesystem and online editing
- Code formatter
