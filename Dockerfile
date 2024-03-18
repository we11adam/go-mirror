FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

FROM scratch
COPY --from=builder /app/server /server
EXPOSE 9080

ENTRYPOINT ["/server"]
