[![BuildStatus](https://travis-ci.org/h8liu/e8vm.png?branch=master)](https://travis-ci.org/h8liu/e8vm)

```
go get lonnie.io/e8vm/...
```

# E8VM

Project goal: a self-contained simulated computer system, including:

- `arch8`: A simulated dead simple instruction set that is barely enough for writing an OS (done).
- `asm8`: An assembler for `arch8` (done). [Try it live!](http://lonnie.io/asmplay/)
- `g8`: A programming language that looks like Go but actually works like C (working on it).
- `os8`: A dead simple operating system that is written in `g8` (not started).
- Since `arch8` is dead simple, it can be easily ported to Javascript so that everything can run (slowly) in a browser.
- For self hosting, I could either rewrite everything in `g8`, or port golang to `os8`. Not sure which one is more practical.

Note: I paused the development because I started working at Google, and Google has some policies for hobby projects. 
I promise I will finish this when the time comes. If you would like to contribute, please contact me via email for copyright/license related details.

## Readability

I hope the project can be readable like a novel. This is how I plan to achieve it:

- **Use a simple language.** Written in golang.
- **Write in small files.** Each file has no more than 300 lines, and each line contains no more than 80 chars.
- **Keep no circular dependency.** With no circular dependency among files, the project can be plotted as a [DAG](http://8k.lonnie.io). 

The DAG visualization gives the project an auto-generated "Table of Contents", where a reader can read the entire project 
from left to right in the graph. While the code might be still hard to read, I hope that now a reader can provide 
detailed feedback without the need to dive super deep first. 
For example, to read and provide feedback to the left-most block in 
a package, you now do not need to understand the details of any other blocks in the package, 
because the left-most block depends on nothing.

Try read the code [here](http://8k.lonnie.io).

## For Contributers

To use the `makefile`, you also need to install some tools:

```
go get lonnie.io/e8tools/...
go get github.com/golang/lint/golint
go get github.com/jstemmer/gotags
```

If you would like to contribute, please contact me via email, or just open an issue/pull request.

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
