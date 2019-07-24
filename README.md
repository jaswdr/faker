# Faker

Faker is a Go library that generates fake data for you. Whether you need to bootstrap your database, create good-looking XML documents, fill-in your persistence to stress test it, or anonymize data taken from a production service, Faker is for you.

Faker is heavily inspired by PHP"s [Faker](https://github.com/fzaninotto/Faker)

Faker requires Go >= 1.9.4.

<a href="https://www.buymeacoffee.com/jaswdr" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 11px !important;width: 104px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

[![GoDoc](https://godoc.org/github.com/jaswdr/faker?status.svg)](https://godoc.org/github.com/jaswdr/faker)
[![Linux and macOS Build Status](https://api.travis-ci.org/jaswdr/faker.svg?branch=master&label=Linux+and+macOS+build "Linux and macOS Build Status")](https://travis-ci.org/jaswdr/faker)
[![Windows Build status](https://ci.appveyor.com/api/projects/status/6x5okgq8xe8h73ov?svg=true)](https://ci.appveyor.com/project/jaschweder/faker)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaswdr/faker)](https://goreportcard.com/report/github.com/jaswdr/faker)

## Installation

Add this to your Go file

```go
import "github.com/jaswdr/faker"
```

And run `go get` or `dep ensure` to get the package.

## Basic Usage

Use `faker.New()` to create and initialize a faker generator, which can generate data by accessing properties named after the type of data you want.

```go
import "github.com/jaswdr/faker"

func main() {
    faker := faker.New()

    faker.Person().Name()
    // Lucy Cechtelar

    faker.Address().Address()
    // 426 Jordy Lodge

    faker.Lorem().Text()
    // Dolores sit sint laboriosam dolorem culpa et autem. Beatae nam sunt fugit
    // et sit et mollitia sed.
    // Fuga deserunt tempora facere magni omnis. Omnis quia temporibus laudantium
    // sit minima sint.
}
```

Even if this example shows a method access, each call to `faker.Name()` yields a different (random) result.

```go
p := faker.Person()

for i:=0; i < 10; i++ {
  fmt.Println(p.Name())
}
  // Adaline Reichel
  // Dr. Santa Prosacco DVM
  // Noemy Vandervort V
  // Lexi O"Conner
  // Gracie Weber
  // Roscoe Johns
  // Emmett Lebsack
  // Keegan Thiel
  // Wellington Koelpin II
  // Ms. Karley Kiehn V
```

See more formatters in [godoc](https://godoc.org/github.com/jaswdr/faker)

## Support

We support Go 1.11 or later

## License

Faker is released under the MIT Licence. See the bundled LICENSE file for details.
