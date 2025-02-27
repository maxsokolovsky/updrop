FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN ./scripts/generate-cert.sh
RUN go mod download
RUN CGO_ENABLED=0 go build -o updrop .

FROM alpine:3
COPY --from=builder /app/updrop /usr/local/bin/
COPY --from=builder /app/cert.pem /etc/cert.pem
COPY --from=builder /app/key.pem /etc/key.pem
ENTRYPOINT ["updrop"]
CMD ["-cert", "/etc/cert.pem", "-key", "/etc/key.pem"]
