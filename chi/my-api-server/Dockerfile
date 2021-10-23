FROM golang:latest

WORKDIR /mydir

#COPY go.mod .
#COPY go.sum .
COPY . ./

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "."]