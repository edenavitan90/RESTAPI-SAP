FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /main

ENV GO111MODULE=on

COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

RUN go build .
# This container exposes port 8080 to the outside world
EXPOSE 8080

CMD ["go","run","."]
# Run the executable
#ENTRYPOINT ["./main"]