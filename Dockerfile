FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -v -o /fizzbuzz-server cmd/fizzbuzz-server/main.go

FROM alpine
RUN apk add ca-certificates

COPY --from=build /fizzbuzz-server /fizzbuzz-server

EXPOSE 3000

CMD ["/fizzbuzz-server"]
