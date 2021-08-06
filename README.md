[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/karantin2020/go-fb2/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/karantin2020/go-fb2?status.svg)](https://godoc.org/github.com/karantin2020/go-fb2)

---

### Features
- [Documented API](https://godoc.org/github.com/karantin2020/go-fb2)
- Creates valid FB 2.1 files
- Includes support for adding CSS, images

Python package for working with FictionBook2

## Usage example

```go
package main

import (
    fb2 "github.com/karantin2020/go-fb2"
)

func main() {
    book := fb2.NewFB2("Example book")
    // "Example book" is a book title
    err := book.SetCover("./testdata/AirPlane_400x600.jpg")
    if err != nil {
        panic(err)
    }
    book.SetAuthor(fb2.AuthorType{
        FirstName: "TestFirstName",
        LastName:  "TestLastName",
    })
    book.SetDescription(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut alios omittam, hunc appello, quem ille unum secutus est.`)
	d.AddSection(`<p>Chapter text.</p>
<p><strong>Strong text.</strong></p>
<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.`, "Chapter 1")
    err = d.WriteToFile("ExampleBook.fb2")
    if err != nil {
        panic(err)
    }
}
```

## Installation

- use [Go modules](https://golang.org/ref/mod)

### Contributions

Contributions are welcome.

### Development

Clone this repository using Git. Run tests as documented below.

Dependencies are managed using [Go modules](https://golang.org/ref/mod)

#### Run tests

```
go test
```