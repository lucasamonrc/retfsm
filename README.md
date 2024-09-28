# 🎰 RegEx To Finite State-Machine - retfsm

During my sophomore year in college I took a class called "Discrete Structures". We spent a great deal of the first quarter of that class talking about regular expression and their FSM counterparts. A lot of my homework involved of taking language grammars (usually represented through regular expressions) and converting them to FSM by hand. This tool is a gift to my past self.

`retfsm` is a CLI application that converts regular expression to is Finite State-Machine (FSM) graph representation. Output to `.dot`, `.svg`, `.png`, and `.jpg` files. Run `retfsm` as a REPL or as command with `retfsm draw`. Input can be passed by stdin redirection, file, or literal strings. Output, if not specified, defaults to a `.dot` file.

### Example

```sh
$ retfsm draw "ab*c" # Creates `a.dot` file in the current directory
$ echo "ab*c" | retfsm draw # Creates `a.dot` file in the current directory
$ retfsm draw "ab*c" out.png # Creates `out.png` file in the current directory
```

Alternatively you can run `retfsm` as a REPL.

```sh
$ refsm
retfsm> ab*c
# The DOT output will be printed to stdout
```

You can also pass text files and redirect `stdin`. Note that any multi-line inputs will produce one output for each line.

## ⚙️ Installation

You must have go installed in your system. This application depends on Go version `1.23.0`. If Go is installed, you can simply run the command below:

```sh
go install github.com/lucasamonrc/retfsm@latest
```

You can check if everything works by running the `help` command:

```sh
retfsm help
```

## 🚀 Technology Stack

This project mostly relies on Go's standard libraries that are included with the version `1.23.0`. Some other important dependencies:

- [graphviz](https://graphviz.org/): Graphviz is open source graph visualization software
- [go-graphviz](https://github.com/goccy/go-graphviz): Provides Go bindings for the `graphviz` package
- [DOT Language](https://graphviz.org/doc/info/lang.html): A simple DSL for defining `graphviz` nodes, edges, and graphs

## 🧑‍💻 Authors

- Lucas Castro ([@lucasamonrc](https://github.com/lucasamonrc)) | [lucasamonrc.dev](https://lucasamonrc.dev)
