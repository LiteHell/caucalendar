FROM golang:1.21-rc-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY *.go ./
RUN go build -v -o /app/app

COPY static ./static

CMD ["/app/app"]