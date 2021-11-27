# go-c2dmc
[![Go Reference](https://pkg.go.dev/badge/github.com/syke99/go-c2dmc.svg)](https://pkg.go.dev/github.com/syke99/go-c2dmc)
[![go reportcard](https://goreportcard.com/badge/github.com/syke99/go-c2dmc)](https://goreportcard.com/report/github.com/syke99/go-c2dmc)

A Go package for converting RGB and other color formats/colorspaces into DMC thread colors (DMC color name and floss number). Implemented as a wrapper for a handful of pieces of functionality from the  [go-colorful](https://github.com/lucasb-eyer/go-colorful/) package with additional functionality to return DMC thread colors. Supports Go 1.13 onwards.

Why was go-c2dmc built?
====
When wanting to build a personal project that invloves DMC thread colors, I could only find one pre-written package for converting RGB values into DMC thread colors (and their corresponding floss numbers) or finding the closest match. However, the package I found was written in R and I wanted to build the project I had in mind in Go. I explored methods of using this newly-found R package and found some interesting solutions. One being a "proxy" of sorts between Go and R written in C. While that solution was cool nonetheless, it felt much too bloated for the usecase I was encountering, as I would only be needing to use the return value from one single function in one single R package and not calling between Go and R several different times. Another option I had was to learn R enough to use the package to do the calculations and print the result to a blank text file, which my Go program would subsequently read from. But again, that required learning R. Which, if I'm being honest, is soemthing I would like to tackle. But my main goal was to get better at Go, specifically. So, seeing as there were no existing Go packages, I decided I would build my own (for the first time!) and share it with other developers who may run into this niche problem.

What problem does go-c2dmc solve?
=====
go-c2dmc is a go package that allows a developer calculate the closest (if not exact) match to DMC color threads. The ColorBank is populated with values matching DMC color threads (and their floss numbers) to RBG and Hexcode values already calculated by [threadcolors](https://threadcolors.com). go-c2dmc can calculate the closest match from this list with any provided RGB, HSV, LAB, or Hexcode values by calculating the distance between the provided color and each color in the ColorBank in LAB colorspace using the [go-colorful](https://github.com/lucasb-eyer/go-colorful/) package.


For more information regarding the accuracy of these calculations specific to each colorspace, please refer to the go-colorful README



How do I use go-c2dmc?
====

### Installing
To install go-c2dmc, simply run

```bash
$ go get github.com/syke99/go-c2dmc
```

Then you can import the package in any go file you'd like

```go
import "github.com/syke99/go-c2dmc"
```

### Basic usage

Calculate the closest DMC thread color based on provided RGB values

```go
// Initialize the ColorBank to have DMC colors to test against
cb := dmc.NewColorBank()

// Call cb.Rgb() and pass in RGB values
color, floss := cb.Rgb(245.0, 173.0, 173.0)

fmt.SprintF("DMC Color Name: %s, Floss number: %s", color, floss)
```

Running the above calculation will return:


```bash
DMC Color Name: Salmon, Floss number: 760
```

And then converting this RGB color into LAB and HSV colorspaces, as well as hexcode:

```go
l, a, b := cb.RgbLab(245.0, 173.0, 173.0)
h, s, v := cb.RgbHsv(245.0, 173.0, 173.0)
hexcode := cb.RgbHex(245.0, 173.0, 173.0)
```

**!!NOTE!!**

Converting from Hexcodes to RGB/LAB/HSV colorspaces is currently unusable, but will be implemented in v2

Who?
====

This library was developed by Quinn Millican (@syke99)


## License

This repo is under the MIT license, see [LICENSE](LICENSE) for details.
