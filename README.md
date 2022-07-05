# Typher

Typher is a typing game for coding that runs on the CLI.

## Installation

```console
$ go install github.com/skyloken/typher@latest
```

## Usage

```console
$ typher <typing-target>
```
- `target`: File path or URL.

## Example

```console
$ typher target.go
$ typher https://raw.githubusercontent.com/golang/go/master/src/io/io.go
```