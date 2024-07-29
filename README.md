# package-headers-go

Generate consistent _Go_ package headers every time.

## Build

You need Go installed on your machine. See the installation guide
[here](https://go.dev/doc/install).

Install the CLI globally like this:

```sh
git clone github.com/Jesserc/package-headers-go
cd package-headers-go/cmd/package-headers-go
go install .
```

## Usage

Generate a Go package header and copy it to your clipboard:

```sh
package-headers-go -package reverser -description "Utility function for reversing strings." -example "import \"github.com/example/reverser\"\n\nstr := \"hello\"\nreversedStr := reverser.ReverseString(str)\nfmt.Println(reversedStr) // Output: \"olleh\""
```

This will generate a header like:

```go
/*
Package reverser provides Utility function for reversing strings.

import "github.com/example/reverser"

str := "hello"
reversedStr := reverser.ReverseString(str)
fmt.Println(reversedStr) // Output: "olleh"
*/
package reverser
```

And it will automatically copy the header to your clipboard.

## Credits

Inspired by t11's [`headers`](https://github.com/transmissions11/headers).
