# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/bauenlabs/echo

# Build and install Echo and its dependencies
RUN go install github.com/bauenlabs/echo

# Configure Environmental Variables
ENV ECHO_REDIS_HOST 54.167.127.6
ENV ECHO_REDIS_PORT 6379
ENV ECHO_SERVER_PORT 8000
ENV ECHO_MODE release
ENV ECHO_CACHE true
# Run echo
ENTRYPOINT /go/bin/echo

# Document that the service listens on port 8080.
EXPOSE 8000
