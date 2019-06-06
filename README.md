# EPITECHFETCHER

## Description
EpitechFetcher is a golang library that fetch the user informations with his autologin
## Installation

Use the package manager [pip](https://pip.pypa.io/en/stable/) to install foobar.

```bash
go get -v github.com/lbrulet/EpitechFetcher
```

## Usage

```golang
package main

import (
	"fmt"

	"github.com/lbrulet/EpitechFetcher"
)

func main() {
	if user, notes, err := EpitechFetcher.Link("auth-c6ec93ec7229061b2564fb41c575d37c915cc55e", "luc.brulet@epitech.eu"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user, notes)
	}
}
```
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License
[MIT](https://choosealicense.com/licenses/mit/)