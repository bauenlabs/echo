# Echo
Echo is a highly optimized caching mechanism.

## Development
### Getting started
* Install Homebrew.
* Install Go: `brew install go`
* Install Glide: `brew install glide`
* Install golint: `go get github.com/golang/lint/golint`
* Install fresh: `go get github.com/pilu/fresh`
* Install godocdown: `go get github.com/robertkrimen/godocdown/godocdown
* Clone this directory.
* In this directory's root, run: `glide install`.

### Workflow
This project uses `fresh` to listen for file changes, and re-compile and run automatically. To use fresh, simply run the `fresh` command in a terminal instance, and leave it running.

### Documentation
This project should adhere to the documentation standards outlined by Go's creaters, which you can find [here](https://blog.golang.org/godoc-documenting-go-code). This project uses a tool called [godocdown](https://github.com/robertkrimen/godocdown) to generate `README.md` files for each of it's packages. When creating or updating a package, before submitting changes, be sure to run the following command for your package:
```bash
godocdown github.com/voiceis/echo/lib/yourpackage > system/path/to/your/package/README.md
```

