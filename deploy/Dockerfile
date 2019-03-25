FROM golang:latest

WORKDIR /go/src/github.com/srgupta5328/Employee
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["employee"]