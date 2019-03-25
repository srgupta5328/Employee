# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/srgupta5328/Employee

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/srgupta5328/Employee

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/Employee

# Document that the service listens on port 8080.
EXPOSE 8934