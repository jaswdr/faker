# Faker

Faker is a Go library that generates fake data for you. Whether you need to bootstrap your database, create good-looking XML documents, fill-in your persistence to stress test it, or anonymize data taken from a production service, Faker is for you.

Faker is heavily inspired by PHP"s [Faker](https://github.com/fzaninotto/Faker)

Faker requires Go >= 1.9.4.

[![GoDoc](https://godoc.org/github.com/jaschweder/faker?status.svg)](https://godoc.org/github.com/jaschweder/faker)
[![Linux and macOS Build Status](https://api.travis-ci.org/jaschweder/faker.svg?branch=master&label=Linux+and+macOS+build "Linux and macOS Build Status")](https://travis-ci.org/jaschweder/faker)
[![Windows Build Status](https://ci.appveyor.com/api/projects/status/cgjqw3h5b59p7at9?svg=true&label=Windows+build "Windows Build Status")](https://ci.appveyor.com/project/jaschweder/faker/branch/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/jaschweder/faker)](https://goreportcard.com/report/github.com/jaschweder/faker)

# Table of Contents

- [Installation](#installation)
- [Basic Usage](#basic-usage)
- [Formatters](#formatters)
	- :ballot_box_with_check: [Base](#base-faker)
	- :ballot_box_with_check: [Lorem Ipsum Text](#lorem-fakerlorem)
	- :ballot_box_with_check: [Person](#person-fakerperson)
	- :ballot_box_with_check: [Address](#address-fakeraddress)
	- :ballot_box_with_check: [Phone](#phone-fakerphone)
	- :ballot_box_with_check: [Company](#fakerprovideren_uscompany)
	- :black_square_button: [Date and Time](#fakerprovidertime.Time)
	- :black_square_button: [Internet](#fakerproviderinternet)
	- :black_square_button: [User Agent](#fakerprovideruseragent)
	- :black_square_button: [Payment](#fakerproviderpayment)
	- :black_square_button: [Color](#fakerprovidercolor)
	- :black_square_button: [File](#fakerproviderfile)
	- :black_square_button: [Image](#fakerproviderimage)
	- :black_square_button: [Uuid](#fakerprovideruuid)
	- :black_square_button: [Barcode](#fakerproviderbarcode)
	- :black_square_button: [Miscellaneous](#fakerprovidermiscellaneous)
	- :black_square_button: [Biased](#fakerproviderbiased)
	- :black_square_button: [Html Lorem](#fakerproviderhtmllorem)
- [License](#license)


## Installation

Add this to your Go file

```go
import "jaschweder/faker"
```

And run `go get` or `dep ensure` to get the package.

## Basic Usage

Use `faker.NewFaker()` to create and initialize a faker generator, which can generate data by accessing properties named after the type of data you want.

```go
import "jaschweder/faker"

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

## Formatters

Each of the generator properties (like `name`, `address`, and `lorem`) are called "formatters". A faker generator has many of them, packaged in "providers". Here is a list of the bundled formatters in the default locale.

### Base (`faker.*`)

    RandomDigit()                              // 7
    RandomDigitNot(1,2,3)                      // 5
    RandomDigitNotNull()                       // 5
    RandomNumber(8)                            // 79907610
    RandomFloat(4, 0, 5)                       // 48.8932
    IntBetween(1000, 9000)                  // 8567
    RandomLetter()                             // "b"
    RandomStringElement([]string("a","b","c")) // "b"
    ShuffleString("hello, world")              // "rlo,h eoldlw"
    Numerify("Hello ###")                      // "Hello 609"
    Lexify("Hello ???")                        // "Hello wgt"
    Bothify("Hello ##??")                      // "Hello 42jz"
    Asciify("Hello ***")                       // "Hello R6+"
    Regexify("[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}"); // sm0@y8k96a.ej

### Lorem (`faker.Lorem().*`)
    Word()               // "aut"
    Words(3)             // array("porro", "sed", "magni")
    Sentence(6)          // "Sit vitae voluptas sint non voluptates."
    Sentences(3)         // []string{"Optio quos qui illo error.", "Laborum vero a officia id corporis.", "Saepe provident esse hic eligendi."}
    Paragraph(3)         // "Ut ab voluptas sed a nam. Sint autem inventore aut officia aut aut blanditiis. Ducimus eos odit amet et est ut eum."
    Paragraphs(3)        // []string{Quidem ut sunt et quidem est accusamus aut. Fuga est placeat rerum ut. Enim ex eveniet facere sunt.", "Aut nam et eum architecto fugit repellendus illo. Qui ex esse veritatis.", "Possimus omnis aut incidunt sunt. Asperiores incidunt iure sequi cum culpa rem. Rerum exercitationem est rem."}
    Text(200)            // "Fuga totam reiciendis qui architecto fugiat nemo. Consequatur recusandae qui cupiditate eos quod."

### Person (`faker.Person().*`)

    Title()           // "Ms."
    TitleMale()       // "Mr."
    TitleFemale()     // "Ms."
    Suffix()          // "Jr."
    Name()            // "Dr. Zane Stroman"
    FirstName()       // "Maynard"
    FirstNameMale()   // "Maynard"
    FirstNameFemale() // "Rachel"
    LastName()        // "Zulauf"

### Address (`faker.Address().*`)

    CityPrefix()                     // "Lake"
    SecondaryAddress()               // "Suite 961"
    State()                          // "NewMexico"
    StateAbbr()                      // "OH"
    CitySuffix()                     // "borough"
    StreetSuffix()                   // "Keys"
    BuildingNumber()                 // "484"
    City()                           // "West Judge"
    StreetName()                     // "Keegan Trail"
    StreetAddress()                  // "439 Karley Loaf Suite 897"
    Postcode()                       // "17916"
    Address()                        // "8888 Cummings Vista Apt. 101, Susanbury, NY 95473"
    Country()                        // "Falkland Islands (Malvinas)"
    Latitude(min = -90, max = 90)    // 77.147489
    Longitude(min = -180, max = 180) // 86.211205

### Phone (`faker.Phone().*`)

    Number()              // "201-886-0269 x3767"
    TollFreePhoneNumber() // "(888) 937-7238"
    E164PhoneNumber()     // "+27113456789"

### `faker.Company().*`

    CatchPhrase()   // "Monitored regional contingency"
    Bs()            // "e-enable robust architectures"
    Company()       // "Bogan-Treutel"
    CompanySuffix() // "and Sons"
    JobTitle()      // "Cashier"

### `faker.Time().*`

    unixTime(max = "now")                    // 58781813
    time.Time(max = "now", timezone = null)   // time.Time("2008-04-25 08:37:17", "UTC")
    time.TimeAD(max = "now", timezone = null) // time.Time("1800-04-29 20:38:49", "Europe/Paris")
    iso8601(max = "now")                     // "1978-12-09T10:10:29+0000"
    date(format = "Y-m-d", max = "now")      // "1979-06-09"
    time(format = "H:i:s", max = "now")      // "20:49:42"
    time.TimeBetween(startDate = "-30 years", endDate = "now", timezone = null) // time.Time("2003-03-15 02:00:49", "Africa/Lagos")
    time.TimeInInterval(startDate = "-30 years", interval = "+ 5 days", timezone = null) // time.Time("2003-03-15 02:00:49", "Antartica/Vostok")
    time.TimeThisCentury(max = "now", timezone = null)     // time.Time("1915-05-30 19:28:21", "UTC")
    time.TimeThisDecade(max = "now", timezone = null)      // time.Time("2007-05-29 22:30:48", "Europe/Paris")
    time.TimeThisYear(max = "now", timezone = null)        // time.Time("2011-02-27 20:52:14", "Africa/Lagos")
    time.TimeThisMonth(max = "now", timezone = null)       // time.Time("2011-10-23 13:46:23", "Antarctica/Vostok")
    amPm(max = "now")                    // "pm"
    dayOfMonth(max = "now")              // "04"
    dayOfWeek(max = "now")               // "Friday"
    month(max = "now")                   // "06"
    monthName(max = "now")               // "January"
    year(max = "now")                    // "1993"
    century()                            // "VI"
    timezone()                           // "Europe/Paris"

Methods accepting a `timezone` argument default to `date_default_timezone_get()`. You can pass a custom timezone string to each method, or define a custom timezone for all time methods at once using `faker::setDefaultTimezone(timezone)`.

### `faker.Internet().*`

    email()                   // "tkshlerin@collins.com"
    safeEmail()               // "king.alford@example.org"
    freeEmail()               // "bradley72@gmail.com"
    companyEmail()            // "russel.durward@mcdermott.org"
    freeEmailDomain()         // "yahoo.com"
    safeEmailDomain()         // "example.org"
    userName()                // "wade55"
    password()                // "k&|X+a45*2["
    domainName()              // "wolffdeckow.net"
    domainWord()              // "feeney"
    tld()                     // "biz"
    url()                     // "http://www.skilesdonnelly.biz/aut-accusantium-ut-architecto-sit-et.html"
    slug()                    // "aut-repellat-commodi-vel-itaque-nihil-id-saepe-nostrum"
    ipv4()                    // "109.133.32.252"
    localIpv4()               // "10.242.58.8"
    ipv6()                    // "8e65:933d:22ee:a232:f1c1:2741:1f10:117c"
    macAddress()              // "43:85:B7:08:10:CA"

### `faker.UserAgent().*`

    userAgent()              // "Mozilla/5.0 (Windows CE) AppleWebKit/5350 (KHTML, like Gecko) Chrome/13.0.888.0 Safari/5350"
    chrome()                 // "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) AppleWebKit/5312 (KHTML, like Gecko) Chrome/14.0.894.0 Safari/5312"
    firefox()                // "Mozilla/5.0 (X11; Linuxi686; rv:7.0) Gecko/20101231 Firefox/3.6"
    safari()                 // "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_1 rv:3.0; en-US) AppleWebKit/534.11.3 (KHTML, like Gecko) Version/4.0 Safari/534.11.3"
    opera()                  // "Opera/8.25 (Windows NT 5.1; en-US) Presto/2.9.188 Version/10.00"
    internetExplorer()       // "Mozilla/5.0 (compatible; MSIE 7.0; Windows 98; Win 9x 4.90; Trident/3.0)"

### `faker.Payment.*`

    creditCardType()           // "MasterCard"
    creditCardNumber()         // "4485480221084675"
    creditCardExpirationDate() // 04/13
    creditCardExpirationDateString() // "04/13"
    creditCardDetails()       // array("MasterCard", "4485480221084675", "Aleksander Nowak", "04/13")
    // Generates a random IBAN. Set countryCode to null for a random country
    iban(countryCode)         // "IT31A8497112740YZ575DJ28BP4"
    swiftBicNumber()          // "RZTIAT22263"

### `faker.Color().*`

    hexcolor()               // "#fa3cc2"
    rgbcolor()               // "0,255,122"
    rgbColorAsArray()        // array(0,255,122)
    rgbCssColor()            // "rgb(0,255,122)"
    safeColorName()          // "fuchsia"
    colorName()              // "Gainsbor"

### `faker.File.*`

    fileExtension()          // "avi"
    mimeType()               // "video/x-msvideo"
    // Copy a random file from the source to the target directory and returns the fullpath or filename
    file(sourceDir = "/tmp", targetDir = "/tmp") // "/path/to/targetDir/13b73edae8443990be1aa8f1a483bc27.jpg"
    file(sourceDir, targetDir, false) // "13b73edae8443990be1aa8f1a483bc27.jpg"

### `faker.Image().*`

    // Image generation provided by LoremPixel (http://lorempixel.com/)
    ImageUrl(width = 640, height = 480)                    // "http://lorempixel.com/640/480/"
    ImageUrl(width, height, "cats")                        // "http://lorempixel.com/800/600/cats/"
    ImageUrl(width, height, "cats", true, "Faker")         // "http://lorempixel.com/800/400/cats/Faker"
    ImageUrl(width, height, "cats", true, "Faker", true)   // "http://lorempixel.com/grey/800/400/cats/Faker/" Monochrome image
    Image(dir = "/tmp", width = 640, height = 480)         // "/tmp/13b73edae8443990be1aa8f1a483bc27.jpg"
    Image(dir, width, height, "cats")                      // "tmp/13b73edae8443990be1aa8f1a483bc27.jpg" it"s a cat!
    Image(dir, width, height, "cats", false)               // "13b73edae8443990be1aa8f1a483bc27.jpg" it"s a filename without path
    Image(dir, width, height, "cats", true, false)         // it"s a no randomize images (default: `true`)
    Image(dir, width, height, "cats", true, true, "Faker") // "tmp/13b73edae8443990be1aa8f1a483bc27.jpg" it"s a cat with "Faker" text. Default, `null`.

### `faker.UUID().*`

    UUID()                   // "7e57d004-2b97-0e7a-b45f-5387367791cd"

### `faker.Barcode().*`

    EAN13()          // "4006381333931"
    EAN8()           // "73513537"
    ISBN13()         // "9790404436093"
    ISBN10()         // "4881416324"

### `faker.Miscellaneous().*`

    Boolean() // false
    Boolean(chanceOfGettingTrue = 50) // true
    MD5()           // "de99a620c50f2990e87144735cd357e7"
    SHA1()          // "f08e7f04ca1a413807ebc47551a40a20a0b4de5c"
    SHA256()        // "0061e4c60dac5c1d82db0135a42e00c89ae3a333e7c26485321f24348c7e98a5"
    Locale()        // en_UK
    CountryCode()   // UK
    LanguageCode()  // en
    CurrencyCode()  // EUR
    Emoji()         // 😁

### `faker.Biased().*`

    // get a random number between 10 and 20,
    // with more chances to be close to 20
    BiasedIntBetween(min = 10, max = 20, function = "sqrt")

### `faker.HTML().*`

    //Generate HTML document which is no more than 2 levels deep, and no more than 3 elements wide at any level.
    RandomHtml(2,3)   // <html><head><title>Aut illo dolorem et accusantium eum.</title></head><body><form action="example.com" method="POST"><label for="username">sequi</label><input type="text" id="username"><label for="password">et</label><input type="password" id="password"></form><b>Id aut saepe non mollitia voluptas voluptas.</b><table><thead><tr><tr>Non consequatur.</tr><tr>Incidunt est.</tr><tr>Aut voluptatem.</tr><tr>Officia voluptas rerum quo.</tr><tr>Asperiores similique.</tr></tr></thead><tbody><tr><td>Sapiente dolorum dolorem sint laboriosam commodi qui.</td><td>Commodi nihil nesciunt eveniet quo repudiandae.</td><td>Voluptates explicabo numquam distinctio necessitatibus repellat.</td><td>Provident ut doloremque nam eum modi aspernatur.</td><td>Iusto inventore.</td></tr><tr><td>Animi nihil ratione id mollitia libero ipsa quia tempore.</td><td>Velit est officia et aut tenetur dolorem sed mollitia expedita.</td><td>Modi modi repudiandae pariatur voluptas rerum ea incidunt non molestiae eligendi eos deleniti.</td><td>Exercitationem voluptatibus dolor est iste quod molestiae.</td><td>Quia reiciendis.</td></tr><tr><td>Inventore impedit exercitationem voluptatibus rerum cupiditate.</td><td>Qui.</td><td>Aliquam.</td><td>Autem nihil aut et.</td><td>Dolor ut quia error.</td></tr><tr><td>Enim facilis iusto earum et minus rerum assumenda quis quia.</td><td>Reprehenderit ut sapiente occaecati voluptatum dolor voluptatem vitae qui velit.</td><td>Quod fugiat non.</td><td>Sunt nobis totam mollitia sed nesciunt est deleniti cumque.</td><td>Repudiandae quo.</td></tr><tr><td>Modi dicta libero quisquam doloremque qui autem.</td><td>Voluptatem aliquid saepe laudantium facere eos sunt dolor.</td><td>Est eos quis laboriosam officia expedita repellendus quia natus.</td><td>Et neque delectus quod fugit enim repudiandae qui.</td><td>Fugit soluta sit facilis facere repellat culpa magni voluptatem maiores tempora.</td></tr><tr><td>Enim dolores doloremque.</td><td>Assumenda voluptatem eum perferendis exercitationem.</td><td>Quasi in fugit deserunt ea perferendis sunt nemo consequatur dolorum soluta.</td><td>Maxime repellat qui numquam voluptatem est modi.</td><td>Alias rerum rerum hic hic eveniet.</td></tr><tr><td>Tempore voluptatem.</td><td>Eaque.</td><td>Et sit quas fugit iusto.</td><td>Nemo nihil rerum dignissimos et esse.</td><td>Repudiandae ipsum numquam.</td></tr><tr><td>Nemo sunt quia.</td><td>Sint tempore est neque ducimus harum sed.</td><td>Dicta placeat atque libero nihil.</td><td>Et qui aperiam temporibus facilis eum.</td><td>Ut dolores qui enim et maiores nesciunt.</td></tr><tr><td>Dolorum totam sint debitis saepe laborum.</td><td>Quidem corrupti ea.</td><td>Cum voluptas quod.</td><td>Possimus consequatur quasi dolorem ut et.</td><td>Et velit non hic labore repudiandae quis.</td></tr></tbody></table></body></html>

## License

Faker is released under the MIT Licence. See the bundled LICENSE file for details.
