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
For create new file from a template, run
```
cptools new [Filename] [Language]
```
make sure that you have template in that language in _[templateDir]/[language].txt_
Default _templateDir_ is at `template` directory in current path
To change that option see `cptools new --help`
