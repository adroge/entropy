# Entropy

- [Entropy](#entropy)
	- [Description](#description)
	- [Usage](#usage)
	- [References](#references)

## Description

This package will calculate the entropy from a given string. The alphabet used
to calculate the entropy can be specified.

## Usage

```go
package main

import (
	"fmt"

	"github.com/adroge/entropy"
)

func main() {
	password := "monkey"
	result, err := entropy.Calculate(password)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("Strength of \"%s\" is %s, with entropy: %f\n", password, result, result.Bits)
}
```

```sh
$ go run main.go
Strength of "monkey" is very weak, with entropy: 28.202638
```

It's possible to change the alphabets and entropy strength bounds by utilizing the Alphabet(...) and Bounds(...) functions. If you wanted to include a subset of the Latin-1 Supplement characters, you could do it with Alphabet(...).

Unit tests are a good source of usage.

## References

These aren't the only places, or even the original places, where the formula used can be found, but for my own reference I used the following:

- [How to Calculate Password Entropy?](https://generatepasswords.org/how-to-calculate-entropy/)
- [Password Security: Complexity vs. Length](http://resources.infosecinstitute.com/password-security-complexity-vs-length/)
- [Entropy as a measure of password strength](https://en.wikipedia.org/wiki/Password_strength#Entropy_as_a_measure_of_password_strength)
