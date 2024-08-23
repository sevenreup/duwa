# Duwa Programming Language (in progress)
This project is an interpreter for the `duwa` language (Name still getting worked shopped), a language based on the Chewa Bantu language.
This project is written in Golang.

## Setup and Moni ku dziko
Download the prebuilt binaries for your platform from the [release](https://github.com/sevenreup/duwa/releases)

Create a new source file, `main.ny` (💀 for now the extension does not matter).

Paste the following cool code

```c#
lemba("Moni Dziko");
```
Run you new application

```bash
duwa -f ./main.ny
```

## Zowerenga bwa?
Documentation is on its way, stay on the cutting edge and freestyle your code.

You can check out examples in `./examples` folder to quickly see what the language can do for now.

## Main Milestones
- [ ] Create an initial interpreter
    - [x] Lexer
    - [x] Parser (Recursive descent parser)
    - [x] Interpreter (Tree walking interpreter)
- [ ] Language features
    - [x] Basic data types (string, numbers)
    - [x] arithmetic operations 
    - [x] Control Flow (loops, conditional statements)
    - [x] Functions
    - [ ] Type checking (for now it does not strictly enforce types)
    - [ ] Data Structures
        - [x] arrays
        - [x] dictionaries
    - [ ] Input/ Output
    - [ ] Error Handling
    - [ ] Class support

Other Milestones
- [ ] Modularity
- [ ] Standard library
