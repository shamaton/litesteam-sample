FROM golang:1.18.2-buster as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd/
COPY ent ent/
COPY model model/
COPY sqlite sqlite/
COPY main.go main.go

RUN go build -o main .

CMD ["/app/main"]