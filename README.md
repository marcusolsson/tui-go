# tui: Terminal UI for Go

[![Build Status](https://travis-ci.org/marcusolsson/tui-go.svg?branch=master)](https://travis-ci.org/marcusolsson/tui-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/marcusolsson/tui-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcusolsson/tui-go)](https://goreportcard.com/report/github.com/marcusolsson/tui-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

A UI library for terminal applications.

tui (pronounced _tooey_) provides a higher-level programming model for building rich terminal applications. It lets you build layout-based user interfaces that (should) gracefully handle resizing for you.

_IMPORTANT:_ tui-go is still in an experimental phase so please don't use it for anything other than experiments, yet.

![Screenshot](example/chat/screenshot.png)

## Installation

```
go get github.com/marcusolsson/tui-go
```

## Usage

```go
package main

import "github.com/marcusolsson/tui-go"

func main() {
	box := tui.NewVBox(
		tui.NewLabel("tui-go"),
	)

	ui := tui.New(box)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		panic(err)
	}
}
```

## Getting started

If you want to know what it is like to build terminal applications with tui-go, check out some of the [examples](example).

Documentation is available at [godoc.org](https://godoc.org/github.com/marcusolsson/tui-go).

## Related projects

tui-go is mainly influenced by [Qt](https://www.qt.io/) and offers a similar programming model that has been adapted to Go and the terminal.

For an overview of the alternatives for writing terminal user interfaces, check out [this article](https://appliedgo.net/tui/) by [AppliedGo](https://appliedgo.net/).

## License

tui-go is released under the [MIT License](LICENSE).
