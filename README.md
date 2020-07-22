# Entropy

Will calculate the entropy bits from a given string. This is a metric, unintended to be used by itself, to make basic credentials more secure. It is a good idea to use this in conjunction with an updated password hashing algorithm such as [Argon2](https://github.com/P-H-C/phc-winner-argon2). The Go version can be found here: [Argon2 Go Package](golang.org/x/crypto/argon2).

## Installation

```sh
go get -d github.com/adroge/entropy
```

## Usage

```go
password := "monkey"
result, err := entropy.Calculate(password)
if err != nil {
	fmt.Println("oops, saw an error: ", err)
}
fmt.Printf("Strength is %s, with entropy: %f\n", result, result.Bits)
```

It's possible to change the alphabets and entropy strength bounds by utilizing the Alphabet(...) and Bounds(...) functions. If you wanted to include a subset of the Latin-1 Supplement characters, you could do it with Alphabet(...).

Unit tests are a good source of usage.

## Why

This module is intended to be used for backend development where a service needs to quickly validate the entropy of a password.

There are other packages out there that do similar things, but they were either, in my opinion, overkill, or didn't fit the need I had.

## References

These aren't the only places, or even the original places, where the formula used can be found, but for my own reference I used the following:

* [How to Calculate Password Entropy?](https://generatepasswords.org/how-to-calculate-entropy/)
* [Password Security: Complexity vs. Length](http://resources.infosecinstitute.com/password-security-complexity-vs-length/)
* [Entropy as a measure of password strength](https://en.wikipedia.org/wiki/Password_strength#Entropy_as_a_measure_of_password_strength)
