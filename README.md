<p align="center">
  <img width="500" src="./cover.png">
</p>

Faker is a Go library that generates fake data for you. Whether you need to bootstrap your database, create good-looking XML documents, fill-in your persistence to stress test it, or anonymize data taken from a production service, Faker is for you.

Faker is heavily inspired by PHP"s [Faker](https://github.com/fzaninotto/Faker)

Faker requires Go >= 1.11

<a href="https://www.buymeacoffee.com/jaswdr" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 11px !important;width: 104px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

[![PkgGoDev](https://pkg.go.dev/badge/github.com/jaswdr/faker)](https://pkg.go.dev/github.com/jaswdr/faker)
[![Build Status](https://travis-ci.org/jaswdr/faker.svg?branch=master)](https://travis-ci.org/jaswdr/faker)
[![Coverage Status](https://coveralls.io/repos/github/jaswdr/faker/badge.svg?branch=master)](https://coveralls.io/github/jaswdr/faker?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaswdr/faker)](https://goreportcard.com/report/github.com/jaswdr/faker)
[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/jaswdr/faker)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ba14f84a3f824410be0a6f6670de012a)](https://app.codacy.com/gh/jaswdr/faker?utm_source=github.com&utm_medium=referral&utm_content=jaswdr/faker&utm_campaign=Badge_Grade)
[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fjaswdr%2Ffaker&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://hits.seeyoufarm.com)

## Test it in Go Playground

Start at https://play.golang.org/p/JpTagDGBaHK

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

Generate fake data using Structs

```go
type ExampleStruct struct {
	SimpleStringField string
	SimpleNumber int
	SimpleBool bool
	SomeFormatedString string `fake:"??? ###"`
	SomeStringArray [5]string `fake:"????"`
}

example := ExampleStruct{}
f.Struct().Fill(&example)
fmt.Printf("%+v", example)
//{SimpleStringField:87576a01c2a547b2bbf9b7c736d1db40 SimpleNumber:9223372036854775807 SimpleBool:false SomeFormatedString:cxo 321 SomeStringArray:[effr swxp ldnj obcs nvlg]}
```

See more formatters in [docs](https://pkg.go.dev/github.com/jaswdr/faker?tab=doc)

## Get involved

Have a question? Use the [Discussions](https://github.com/jaswdr/faker/discussions) page.

## Development

Create a fork and get the code.

```bash
$ go get github.com/jaswdr/faker
```

Do your changes, add tests, run the tests.

```bash
$ go test
PASS
ok      github.com/jaswdr/faker 0.010s
```

Push to your fork and send a new pull request from your fork to this repository.

## License

Faker is released under the MIT Licence. See the bundled LICENSE file for details.
