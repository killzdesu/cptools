# Configuration Guide

## Overview

`cptools` uses a hybrid configuration system:
- **Embedded defaults**: Configuration is embedded in the binary, so it works out of the box
- **User overrides**: Users can create custom configurations that override the defaults

## User Configuration

### Location

Create your custom configuration at:

**Windows:**
```
%APPDATA%\cptools\languages.yaml
```

**Linux/macOS:**
```
~/.config/cptools/languages.yaml
```

### Format

The configuration file uses YAML format. Here's an example:

```yaml
languages:
  cpp:
    name: "C++"
    extension: "cpp"
    compile: "g++ -std=c++20 -O3 -Wall {{filename}} -o {{output}}"
    run: "./{{output}}"
    template: "template/cpp.cpp"
    compiled: true
```

### Available Variables

The following variables are available for substitution in compile and run commands:

- `{{filename}}` - Full filename with extension (e.g., `solution.cpp`)
- `{{basename}}` - Filename without extension (e.g., `solution`)
- `{{output}}` - Output executable name (platform-specific)
- `{{ext}}` - File extension (e.g., `cpp`)

### Configuration Fields

For each language, you can configure:

- `name`: Display name of the language
- `extension`: File extension (without the dot)
- `compile`: Command to compile the source file (empty string if not compiled)
- `run`: Command to run the program
- `template`: Path to the template file (relative to cptools directory)
- `compiled`: Boolean indicating if the language requires compilation

## Default Languages

`cptools` comes with built-in configurations for:

- **C++**: `g++` with C++17, O2 optimization
- **Go**: `go build`
- **Python**: `python3` interpreter
- **JavaScript**: `node` interpreter
- **Rust**: `rustc` compiler

## Customization Examples

### Change C++ Standard and Flags

```yaml
languages:
  cpp:
    name: "C++"
    extension: "cpp"
    compile: "g++ -std=c++20 -O3 -march=native -Wall -Wextra {{filename}} -o {{output}}"
    run: "./{{output}}"
    template: "template/cpp.cpp"
    compiled: true
```

### Add a New Language (C)

```yaml
languages:
  c:
    name: "C"
    extension: "c"
    compile: "gcc -std=c11 -O2 -Wall {{filename}} -o {{output}}"
    run: "./{{output}}"
    template: "template/c.c"
    compiled: true
```

### Use Different Python Version

```yaml
languages:
  python:
    name: "Python"
    extension: "py"
    compile: ""
    run: "python3.11 {{filename}}"
    template: "template/python.py"
    compiled: false
```

### Use TypeScript with ts-node

```yaml
languages:
  typescript:
    name: "TypeScript"
    extension: "ts"
    compile: ""
    run: "ts-node {{filename}}"
    template: "template/typescript.ts"
    compiled: false
```

## Usage

### Create a new file from template

```bash
cptools new solution cpp
cptools new main go
cptools new script python
```

### Compile a file (for compiled languages)

```bash
cptools compile solution.cpp
cptools compile main.go
```

### Run a program

```bash
cptools run solution.cpp    # Runs the compiled executable
cptools run script.py       # Runs with Python interpreter
```

## Notes

- If you don't create a user config file, `cptools` will use the embedded defaults
- User config files only need to specify the languages you want to override
- When you upgrade `cptools`, embedded defaults will update automatically
- Your user config will always take precedence over embedded defaults
- Template files are still stored in the `template/` directory relative to where you run `cptools`
