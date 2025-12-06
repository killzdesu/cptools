# cptools
Tool for Competitive Programming -- Creating, Compiling, and Testing

For language such as C++, Go that need compiling and executing then testing with testcases
it takes time to do such thing in situation that is time constrain such as Codeforces contest, Google Codejam

So this tool is used to help automate those stuffs

## Installation
```
git clone https://github.com/killzdesu/cptools.git
cd cptools
go build .
cptools --version
```

## Usage

### Create a new file from template
```bash
cptools new [filename] [language]
# Example:
cptools new solution cpp
cptools new main go
cptools new script python
```

Supported languages: `cpp`, `go`, `python`, `javascript`, `rust`

### Compile a file
```bash
cptools compile [filename]
# Example:
cptools compile solution.cpp
cptools compile main.go
```

### Run a program
```bash
cptools run [filename]
# Example:
cptools run solution.cpp
cptools run script.py
```

## Configuration

cptools uses a hybrid configuration system:
- **Embedded defaults**: Works out of the box with standard compiler settings
- **User overrides**: Customize compile/run commands by creating a config file

See [CONFIG.md](CONFIG.md) for detailed configuration instructions.

### User Config Location
- Windows: `%APPDATA%\cptools\languages.yaml`
- Linux/macOS: `~/.config/cptools/languages.yaml`

## Template Files

Template files are located in the `template/` directory. You can customize them or add new languages through the configuration system.

