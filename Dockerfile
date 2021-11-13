FROM golang:latest as builder
WORKDIR /go/src/tlswebsite
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/tlswebsite .

FROM alpine
WORKDIR /go/src/tlswebsite
COPY --from=builder /go/src/tlswebsite/build/tlswebsite .
COPY --from=builder /go/src/tlswebsite/ .
ENTRYPOINT ["./tlswebsite"]

