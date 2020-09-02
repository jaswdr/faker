你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Faker

Faker is a Go library that generates fake data for you. Whether you need to bootstrap your database, create good-looking XML documents, fill-in your persistence to stress test it, or anonymize data taken from a production service, Faker is for you.

Faker is heavily inspired by PHP"s [Faker](https://github.com/fzaninotto/Faker)

Faker requires Go >= 1.11

<a href="https://www.buymeacoffee.com/jaswdr" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 11px !important;width: 104px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

[![GoDoc](https://godoc.org/github.com/jaswdr/faker?status.svg)](https://godoc.org/github.com/jaswdr/faker)
[![Build Status](https://travis-ci.org/jaswdr/faker.svg?branch=master)](https://travis-ci.org/jaswdr/faker)
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

See more formatters in [docs](https://pkg.go.dev/github.com/jaswdr/faker?tab=doc)

## License

Faker is released under the MIT Licence. See the bundled LICENSE file for details.
