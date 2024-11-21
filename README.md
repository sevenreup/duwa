# Duwa Programming Language (in progress)

This project is an interpreter for the `duwa` language, a language based on the Chewa Bantu language.
This project is written in Golang.

## Setup and Zowerenga

Follow the setup documentations at [official site](https://www.duwa.cphiri.dev)

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
  - [ ] Working Wasm version of the language

Other Milestones

- [ ] Modularity
- [ ] Standard library

## Building

### Prerequisites

- TinyGo: If you are building the web assembly version you will need `TinyGo`, hers is the [setup](https://tinygo.org/getting-started/install)
- GNU make (optional): Using make for building. You can build the project without `make`.

### Builing/running

You can build the project by running any of the build make tasks

```bash
make build
```

Or you can run go build directly

```bash
go build -C ./src/cmd/duwa -o ../../../bin/duwa
```

Running the project is the same either use make or go directly
