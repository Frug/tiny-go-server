FROM golang:1.16.5-alpine3.14 AS builder

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main main.go

FROM scratch
COPY --from=builder /app/main /main
CMD ["/main"]
