FROM golang:latest

WORKDIR /mydir

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go ./

EXPOSE 8080

CMD ["go", "run", "."]