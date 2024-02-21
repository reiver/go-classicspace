# go-classicspace

Package **classicspace** provides a color palette associated with the **classic space** toy series, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-classicspace

[![GoDoc](https://godoc.org/github.com/reiver/go-classicspace?status.svg)](https://godoc.org/github.com/reiver/go-classicspace)

## Example

Here is an example:

```golang
import "github.com/reiver/go-classicspace"

// ...

var bounds image.Rectangle = image.Rectangle{
	Min: image.Point{0,0}
	Max: image.Point{640,480}
}

var img image.Image = image.NewPaletted(bounds, classicspace.Palette)

// ...

```

## Import

To import package **classicspace** use `import` code like the follownig:
```
import "github.com/reiver/go-classicspace"
```

## Installation

To install package **classicspace** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-classicspace
```

## Author

Package **classicspace** was written by [Charles Iliya Krempeaux](http://changelog.ca)
