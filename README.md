# Echo
Echo is a highly optimized caching mechanism.

## Development
### Getting started
* Install Homebrew.
* Install Go: `brew install go`
* Install Glide: `brew install glide`
* Install golint: `go get github.com/golang/lint/golint`
* Install fresh: `go get github.com/pilu/fresh`
* Install godocdown: `go get github.com/robertkrimen/godocdown/godocdown`
* Clone this directory.
* In this directory's root, run: `glide install`.

### Environment Variables
* Ensure the following environment variables exist and have valid values:
```shell
# Path at which Redis can be contacted. Defaults to localhost.
export ECHO_REDIS_HOST="localhost"

# Port at which Redis can be contacted. Defaults to 6379.
export ECHO_REDIS_PORT="6379"

# Port on which the Echo server should run. Defaults to 8000.
export ECHO_SERVER_PORT="8000"

# Logging mode for server. Defaults to release.
export ECHO_MODE="release|debug|test"

# Toggle whether or not requests are cached. Defaults to true.
export ECHO_CACHE="true"
```

### Workflow
This project uses `fresh` to listen for file changes, and re-compile and run automatically. To use fresh, simply run the `fresh` command in a terminal instance, and leave it running.

### Documentation
This project should adhere to the documentation standards outlined by Go's creaters, which you can find [here](https://blog.golang.org/godoc-documenting-go-code). This project uses a tool called [godocdown](https://github.com/robertkrimen/godocdown) to generate `README.md` files for each of it's packages. When creating or updating a package, before submitting changes, be sure to run the following command for your package:
```bash
godocdown github.com/bauenlabs/echo/lib/yourpackage > system/path/to/your/package/README.md
```

## Testing
All code in this project should be tested. To get an idea how to write tests, check out some of the existing tests. There should be clear examples on how assertions, mocks, etc are created.

### Running Tests
To run the tests in this project, simply run:
```bash
go test -cover
```
