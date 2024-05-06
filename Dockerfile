FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o updrop .

FROM alpine:3
COPY --from=builder /app/updrop /usr/local/bin/
ENTRYPOINT ["updrop"]
CMD ["-addr", "8090"]

