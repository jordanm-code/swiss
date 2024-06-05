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
- [func Chunk\[T \~\[\]E, E any\]\(s T, size int\) \(chunks \[\]T\)](<#Chunk>)
- [func Deref\[T any\]\(v \*T\) \(r T\)](<#Deref>)
- [func ExtractURLs\(s string\) \[\]string](<#ExtractURLs>)
- [func FileExists\(filePath string\) bool](<#FileExists>)
- [func IsAlpha\(s string\) bool](<#IsAlpha>)
- [func IsAlphaNumeric\(s string\) bool](<#IsAlphaNumeric>)
- [func IsEmail\(s string\) bool](<#IsEmail>)
- [func IsEmpty\(x interface\{\}\) bool](<#IsEmpty>)
- [func IsHexChar\(c byte\) bool](<#IsHexChar>)
- [func IsLower\(s string\) bool](<#IsLower>)
- [func IsNumeric\(s string\) bool](<#IsNumeric>)
- [func IsSnakeCase\(s string\) bool](<#IsSnakeCase>)
- [func IsURL\(s string\) bool](<#IsURL>)
- [func IsUUID\(s string\) bool](<#IsUUID>)
- [func IsUpper\(s string\) bool](<#IsUpper>)
- [func Map\[S \~\[\]E, E comparable\]\(s S\) map\[E\]bool](<#Map>)
- [func PascalCase\(s string\) string](<#PascalCase>)
- [func PtrTo\[T any\]\(v T\) \*T](<#PtrTo>)
- [func RandomSeed\(\)](<#RandomSeed>)
- [func RandomString\(length int\) string](<#RandomString>)
- [func Reverse\(s string\) string](<#Reverse>)
- [func Slugify\(s string\) string](<#Slugify>)
- [func SnakeCase\(s string\) string](<#SnakeCase>)
- [func SwapCase\(str string\) string](<#SwapCase>)
- [func TitleCase\(s string\) string](<#TitleCase>)


## Variables

<a name="Language"></a>

```go
var Language = language.English
```

<a name="CamelCase"></a>
## func CamelCase

```go
func CamelCase(s string) string
```

CamelCase converts a string to camel case.

<a name="Chunk"></a>
## func Chunk

```go
func Chunk[T ~[]E, E any](s T, size int) (chunks []T)
```

Chunk divides a slice into chunks of the specified size.

<a name="Deref"></a>
## func Deref

```go
func Deref[T any](v *T) (r T)
```

Deref returns the value that the pointer points to.

<a name="ExtractURLs"></a>
## func ExtractURLs

```go
func ExtractURLs(s string) []string
```

ExtractURLs will take a string and return all URLs found within it in the order they are encountered.

<a name="FileExists"></a>
## func FileExists

```go
func FileExists(filePath string) bool
```

FileExists returns true if the file exists

<details><summary>Example</summary>
<p>



```go
fmt.Printf("file_test.go exists: %v\n", FileExists("file_test.go"))
fmt.Printf("file_test.go.not exists: %v\n", FileExists("file_test.go.not"))

// Output:
// file_test.go exists: true
// file_test.go.not exists: false
```

#### Output

```
file_test.go exists: true
file_test.go.not exists: false
```

</p>
</details>

<a name="IsAlpha"></a>
## func IsAlpha

```go
func IsAlpha(s string) bool
```

IsAlpha checks if a string is all alphabetic characters.

<a name="IsAlphaNumeric"></a>
## func IsAlphaNumeric

```go
func IsAlphaNumeric(s string) bool
```

IsAlphaNumeric checks if a string is all alphanumeric characters.

<a name="IsEmail"></a>
## func IsEmail

```go
func IsEmail(s string) bool
```

IsEmail checks if a string is a valid standard email address. RFC 5322 is too permissive for general use such as quoted strings and local hosts.

<a name="IsEmpty"></a>
## func IsEmpty

```go
func IsEmpty(x interface{}) bool
```

IsEmpty checks if the given value is empty/zero value.

<a name="IsHexChar"></a>
## func IsHexChar

```go
func IsHexChar(c byte) bool
```

IsHexChar checks if a character is a valid hexadecimal character.

<a name="IsLower"></a>
## func IsLower

```go
func IsLower(s string) bool
```

IsLower checks if a string is all lowercase.

<a name="IsNumeric"></a>
## func IsNumeric

```go
func IsNumeric(s string) bool
```

IsNumeric checks if a string is all numeric characters.

<a name="IsSnakeCase"></a>
## func IsSnakeCase

```go
func IsSnakeCase(s string) bool
```

IsSnakeCase checks to see if supplied string is in snake case.

<a name="IsURL"></a>
## func IsURL

```go
func IsURL(s string) bool
```

IsURL checks if a string is a valid URL.

<a name="IsUUID"></a>
## func IsUUID

```go
func IsUUID(s string) bool
```

IsUUID will take a string and determine if it is a valid UUID, it will return true if it is and false if it is not.

<a name="IsUpper"></a>
## func IsUpper

```go
func IsUpper(s string) bool
```

IsUpper checks if a string is all uppercase.

<a name="Map"></a>
## func Map

```go
func Map[S ~[]E, E comparable](s S) map[E]bool
```

Map creates a map from a slice of keys. The value of each key is a boolean indicating whether the key is present in the slice.

<a name="PascalCase"></a>
## func PascalCase

```go
func PascalCase(s string) string
```

PascalCase converts a string to pascal case, it can take in a string that is both title case and camel case and convert it to camel case.

<a name="PtrTo"></a>
## func PtrTo

```go
func PtrTo[T any](v T) *T
```

PtrTo returns a pointer to the value passed in.

<a name="RandomSeed"></a>
## func RandomSeed

```go
func RandomSeed()
```

RandomSeed will generate a seed based on the current UnixNano.

<a name="RandomString"></a>
## func RandomString

```go
func RandomString(length int) string
```

RandomString creates an alphanumeric string of a given length.

<a name="Reverse"></a>
## func Reverse

```go
func Reverse(s string) string
```

Reverse returns the string in reverse order. This function is an alias for \[bidi.ReverseString\].

<a name="Slugify"></a>
## func Slugify

```go
func Slugify(s string) string
```

Slugify will take a string and create a dash separated string for use in URLs.

<a name="SnakeCase"></a>
## func SnakeCase

```go
func SnakeCase(s string) string
```

SnakeCase converts a string to snake case. Inputs can be space separated, camel case, or pascal case.

<a name="SwapCase"></a>
## func SwapCase

```go
func SwapCase(str string) string
```

SwapCase swaps the case of a string.

<a name="TitleCase"></a>
## func TitleCase

```go
func TitleCase(s string) string
```

TitleCase converts a string to title case.

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)


<!-- gomarkdoc:embed:end -->
