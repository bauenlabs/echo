# Echo
Echo is a highly optimized caching mechanism.

## Installation
* Install Homebrew.
* Install Go: `brew install go`
* Install Glide: `brew install glide`
* Run `go get github.com/golang/lint/golint`
* Run `go get github.com/pilu/fresh`
* Clone this directory.
* In this directory's root, run: `glide install`.

## Running
* Ensure the following environment variables exist and have valid values:
```shell
# Path at which Redis can be contacted.
export ECHO_REDIS_HOST="localhost"

# Port at which Redis can be contacted.
export ECHO_REDIS_PORT="6379"

# Port on which the Echo server should run.
export ECHO_SERVER_PORT="8000"

# Logging mode for server.
export ECHO_SERVER_MODE="release|debug|test"
```
