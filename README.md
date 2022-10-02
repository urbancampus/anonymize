# anonymize [![Go Reference](https://pkg.go.dev/badge/github.com/emmanuelay/anonymize.svg)](https://pkg.go.dev/github.com/emmanuelay/anonymize) 

Small and simple package for anonymizing names, e-mails and domains

## Installation

```go
go get -u github.com/emmanuelay/anonymize
```

## Quick start

```go
package main

import (
  "fmt"
  "github.com/emmanuelay/anonymize"
)

func main() {
  name := "John Doe"
  anonymousName := anonymize.Name(name)
  fmt.Println(name, ":", anonymousName)

  email := "john.doe@gmail.com"
  anonymousEmail := anonymize.Email(email)
  fmt.Println(email, ":", anonymousEmail)

  domain := "www.john-doe.com"
  anonymousDomain := anonymize.Domain(domain)
  fmt.Println(domain, ":", anonymousDomain)
}
```

..this will result in the following output:

```sh
John Doe: J*** D**
john.doe@gmail.com: j***.d**@g****.com
www.john-doe.com: www.j***-d**.com
```

## Contribute

Feel free to [create a ticket](https://github.com/emmanuelay/anonymize/issues/new) if you see room for improvements or want to suggest something. 

### Contributors

*   [Emmanuel Ay](https://github.com/emmanuelay)