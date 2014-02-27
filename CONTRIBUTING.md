## Contribute To Wishful

This describes how developers may contribute to Wishful.

## How To Contribute

* Report bugs (via GitHub)
* Give feedback on new feature discussions (via GitHub)
* Propose your own ideas (via GitHub)

### Gofmt Your Code

Set your editor to run "go fmt" every time you save so that whitespace / style
comments are kept to a minimum.

Howtos:
* [Emacs](http://blog.golang.org/2013/01/go-fmt-your-code.html)
* [vim](http://blog.golang.org/go-fmt-your-code)

### Run The Tests

Typically running the main set of unit tests will be sufficient:

```
go test ./useful
```

#### Code coverage

If you're implementing a new feature you should always make sure that code
coverage is never less than when you start a feature (onwards and upwards).

```
go test -coverprofile=coverage.out ./useful; go tool cover -html=coverage.out
```