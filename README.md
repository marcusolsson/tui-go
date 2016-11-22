# tui: Terminal UI for Go

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/marcusolsson/tui-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcusolsson/tui-go)](https://goreportcard.com/report/github.com/marcusolsson/tui-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)
![stability-experimental](https://img.shields.io/badge/stability-experimental-red.svg)

A UI library for terminal applications.

tui (pronounced _tooey_) provides a higher-level programming model for building rich terminal
applications. It lets you build layout-based user interfaces that (should)
gracefully handle resizing for you.

![Example](docs/example.png)

## Status

This project is highly experimental and will change a lot. __Use at your own risk__.

## Installation

```
go get github.com/marcusolsson/tui-go
```

## Usage

```
import "github.com/marcusolsson/tui-go"
```

If you want to know how it's like to build terminal applications with tui, check out some of the [examples](example).

## Documentation

The documentation is rather bare at the moment due to me changing the API
pretty frequently. You can however explore the API in its current form at
[godoc.org](https://godoc.org/github.com/marcusolsson/tui-go).

## Contributing

I'm currently not accepting contributions or bug reports due to the project
being pretty volatile at the moment. That being said, feel free to create an
GitHub issue to suggest features you'd like to see.

## Related projects

tui-go is mainly influenced by [Qt](https://www.qt.io/) and offers a similar programming model that has been adapted to Go and the terminal.

Following Go projects are related to tui but offers different approaches for creating terminal applications.

- [termbox-go](https://github.com/nsf/termbox-go), is used by tui-go for drawing to the terminal.
- [gocui](https://github.com/jroimartin/gocui), is a more minimalistic library for creating console user interfaces.
- [termui](https://github.com/gizak/termui), focuses on building terminal dashboards.

## License

tui-go is released under the [MIT License](LICENSE).
