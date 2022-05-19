FROM golang:1.18.2-buster as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd cmd/
COPY ent ent/
COPY model model/
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o /app/setup ./cmd/setup/
#
#FROM gcr.io/distroless/static-debian10 as setup
#
#COPY --from=builder /app/setup /setup
#
#CMD ["/setup"]