# tui: Terminal UI for Go

[![Build Status](https://travis-ci.org/marcusolsson/tui-go.svg?branch=master)](https://travis-ci.org/marcusolsson/tui-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/marcusolsson/tui-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcusolsson/tui-go)](https://goreportcard.com/report/github.com/marcusolsson/tui-go)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)
![stability-experimental](https://img.shields.io/badge/stability-experimental-red.svg)

A UI library for terminal applications.

tui (pronounced _tooey_) provides a higher-level programming model for building rich terminal applications. It lets you build layout-based user interfaces that (should) gracefully handle resizing for you.

![Screenshot](example/chat/screenshot.png)

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

If you want to know what it is like to build terminal applications with tui-go, check out some of the [examples](example).

## Documentation

The documentation is rather bare at the moment due to me changing the API pretty frequently. You can however explore the API in its current form at [godoc.org](https://godoc.org/github.com/marcusolsson/tui-go).

For now, the best way to learn tui-go is to study and learn from the [examples](example).

## Contributing

Feel free to submit pull requests, but consider letting me know by posting an issue first to make sure that your contributions will outlive any major refactoring in the near future.

Please post any feature requests you might have. Smaller requests might end up being implemented rather quickly and larger ones will be considered for the road map.

### Contributors

- Marcus Olsson ([@marcusolsson](https://github.com/marcusolsson))
- Doug Reese ([@dougreese](https://github.com/dougreese))
- Pontus Leitzler ([@leitzler](https://github.com/leitzler))
- Johan Sageryd ([@jsageryd](https://github.com/jsageryd))
- Yann Malet ([@yml](https://github.com/yml))

## Related projects

tui-go is mainly influenced by [Qt](https://www.qt.io/) and offers a similar programming model that has been adapted to Go and the terminal.

For an overview of the alternatives for writing terminal user interfaces, check out [this article](https://appliedgo.net/tui/) by [AppliedGo](https://appliedgo.net/).

## License

tui-go is released under the [MIT License](LICENSE).
