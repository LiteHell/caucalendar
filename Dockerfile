FROM golang:alpine AS base

WORKDIR /app

# To avoid tls error from swedu.cau.ac.kr
COPY digicert-ca.pem /usr/local/share/ca-certificates/digicert-ca.crt
RUN cat /usr/local/share/ca-certificates/digicert-ca.crt >> /etc/ssl/certs/ca-certificates.crt

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY static ./static
COPY *.go ./

FROM base AS deployment
RUN go build -v -o /app/app
CMD ["/app/app"]

FROM base As test

RUN go test -v ./...
