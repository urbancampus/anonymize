# anonymize [![Go Reference](https://pkg.go.dev/badge/github.com/emmanuelay/anonymize.svg)](https://pkg.go.dev/github.com/emmanuelay/anonymize) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/9d3422b218ee49e49ffe359b88ea2a5f)](https://www.codacy.com/gh/emmanuelay/anonymize/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=emmanuelay/anonymize&amp;utm_campaign=Badge_Grade)

Small and simple package for anonymizing names, e-mails and domains

## Installation

```sh
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