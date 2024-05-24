# Swiss

<p align="center">
<img src="https://raw.githubusercontent.com/jordanm-code/swiss/main/images/swiss.png" width="400">
</p>

Swiss is a collection of functions to help reduce boilerplate code in golang.

## Godoc

<!-- gomarkdoc:embed:start -->

<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# swiss

```go
import "go/swiss"
```

## Index

- [Variables](<#variables>)
- [func CamelCase\(s string\) string](<#CamelCase>)
- [func Deref\[T any\]\(v \*T\) \(r T\)](<#Deref>)
- [func IsAlpha\(s string\) bool](<#IsAlpha>)
- [func IsAlphaNumeric\(s string\) bool](<#IsAlphaNumeric>)
- [func IsEmail\(s string\) bool](<#IsEmail>)
- [func IsEmpty\(x interface\{\}\) bool](<#IsEmpty>)
- [func IsHexChar\(c byte\) bool](<#IsHexChar>)
- [func IsLower\(s string\) bool](<#IsLower>)
- [func IsNumeric\(s string\) bool](<#IsNumeric>)
- [func IsSnakeCase\(s string\) bool](<#IsSnakeCase>)
- [func IsUUID\(s string\) bool](<#IsUUID>)
- [func IsUpper\(s string\) bool](<#IsUpper>)
- [func Map\[S \~\[\]E, E comparable\]\(s S\) map\[E\]bool](<#Map>)
- [func PascalCase\(s string\) string](<#PascalCase>)
- [func PtrTo\[T any\]\(v T\) \*T](<#PtrTo>)
- [func Reverse\(s string\) string](<#Reverse>)
- [func SnakeCase\(s string\) string](<#SnakeCase>)
- [func SwapCase\(str string\) string](<#SwapCase>)
- [func TitleCase\(s string\) string](<#TitleCase>)


## Variables

<a name="Language"></a>

```go
var Language = language.English
```

<a name="CamelCase"></a>
## func [CamelCase](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L134>)

```go
func CamelCase(s string) string
```

CamelCase converts a string to camel case.

<a name="Deref"></a>
## func [Deref](<https://github.com/jordanm-code/swiss/blob/main/generic.go#L9>)

```go
func Deref[T any](v *T) (r T)
```

Deref returns the value that the pointer points to.

<a name="IsAlpha"></a>
## func [IsAlpha](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L33>)

```go
func IsAlpha(s string) bool
```

IsAlpha checks if a string is all alphabetic characters.

<a name="IsAlphaNumeric"></a>
## func [IsAlphaNumeric](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L53>)

```go
func IsAlphaNumeric(s string) bool
```

IsAlphaNumeric checks if a string is all alphanumeric characters.

<a name="IsEmail"></a>
## func [IsEmail](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L18>)

```go
func IsEmail(s string) bool
```

IsEmail checks if a string is a valid standard email address. RFC 5322 is too permissive for general use such as quoted strings and local hosts.

<a name="IsEmpty"></a>
## func [IsEmpty](<https://github.com/jordanm-code/swiss/blob/main/reflection.go#L6>)

```go
func IsEmpty(x interface{}) bool
```

IsEmpty checks if the given value is empty/zero value.

<a name="IsHexChar"></a>
## func [IsHexChar](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L63>)

```go
func IsHexChar(c byte) bool
```

IsHexChar checks if a character is a valid hexadecimal character.

<a name="IsLower"></a>
## func [IsLower](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L28>)

```go
func IsLower(s string) bool
```

IsLower checks if a string is all lowercase.

<a name="IsNumeric"></a>
## func [IsNumeric](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L43>)

```go
func IsNumeric(s string) bool
```

IsNumeric checks if a string is all numeric characters.

<a name="IsSnakeCase"></a>
## func [IsSnakeCase](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L68>)

```go
func IsSnakeCase(s string) bool
```

IsSnakeCase checks to see if supplied string is in snake case.

<a name="IsUUID"></a>
## func [IsUUID](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L74>)

```go
func IsUUID(s string) bool
```

IsUUID will take a string and determine if it is a valid UUID, it will return true if it is and false if it is not.

<a name="IsUpper"></a>
## func [IsUpper](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L23>)

```go
func IsUpper(s string) bool
```

IsUpper checks if a string is all uppercase.

<a name="Map"></a>
## func [Map](<https://github.com/jordanm-code/swiss/blob/main/generic.go#L19>)

```go
func Map[S ~[]E, E comparable](s S) map[E]bool
```

Map creates a map from a slice of keys. The value of each key is a boolean indicating whether the key is present in the slice.

<a name="PascalCase"></a>
## func [PascalCase](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L152>)

```go
func PascalCase(s string) string
```

PascalCase converts a string to pascal case, it can take in a string that is both title case and camel case and convert it to camel case.

<a name="PtrTo"></a>
## func [PtrTo](<https://github.com/jordanm-code/swiss/blob/main/generic.go#L4>)

```go
func PtrTo[T any](v T) *T
```

PtrTo returns a pointer to the value passed in.

<a name="Reverse"></a>
## func [Reverse](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L96>)

```go
func Reverse(s string) string
```

Reverse returns the string in reverse order. This function is an alias for \[bidi.ReverseString\].

<a name="SnakeCase"></a>
## func [SnakeCase](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L102>)

```go
func SnakeCase(s string) string
```

SnakeCase converts a string to snake case. Inputs can be space separated, camel case, or pascal case.

<a name="SwapCase"></a>
## func [SwapCase](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L158>)

```go
func SwapCase(str string) string
```

SwapCase swaps the case of a string.

<a name="TitleCase"></a>
## func [TitleCase](<https://github.com/jordanm-code/swiss/blob/main/strings.go#L129>)

```go
func TitleCase(s string) string
```

TitleCase converts a string to title case.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->
