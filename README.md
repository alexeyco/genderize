# Genderize

![CircleCI](https://img.shields.io/circleci/build/github/alexeyco/genderize/master)
[![Coverage Status](https://coveralls.io/repos/github/alexeyco/genderize/badge.svg?branch=master)](https://coveralls.io/github/alexeyco/genderize?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexeyco/genderize)](https://goreportcard.com/report/github.com/alexeyco/genderize)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/alexeyco/genderize)](https://pkg.go.dev/github.com/alexeyco/genderize)

API-client for [genderize.io](https://genderize.io) – determine the gender of a name.

## Usage
### Check single name
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alexeyco/genderize"
)

func main() {
    client := genderize.NewClient()
    req := genderize.NewRequest(context.TODO()).
        Name("Alex")
	
    gender := client.ExecuteX(req).FirstX()
    // Or:
    // gender := client.ExecuteX(req).FindX("Alex")

    log.Println(fmt.Sprintf("%s is %s", gender.Name, gender.Gender))
}
```

### Iterate over results collection
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/alexeyco/genderize"
)

func main() {
	client := genderize.NewClient()
	req := genderize.NewRequest(context.TODO()).
		Name("Alex").
		Name("John").
		Name("Alice")

	collection := client.ExecuteX(req)

	collection.EachX(func(g *genderize.Gender) {
		log.Println(fmt.Sprintf("%s is %s", g.Name, g.Gender))
	})
}
```

## License
```
MIT License

Copyright (c) 2020 Alexey Popov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
