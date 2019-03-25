FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Rohan Gupta <rohanclouddev@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/srgupta5328/Employee

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8934

# Run the executable
CMD ["employee"]