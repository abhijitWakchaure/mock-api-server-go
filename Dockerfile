FROM golang:alpine as builder
WORKDIR /go/src/github.com/abhijitWakchaure/mock-api-server-go/
COPY . .
RUN go get -u github.com/gobuffalo/packr/v2/packr2
RUN packr2
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o dist/mock-api-server-go-linux-amd64


FROM alpine
WORKDIR /app
COPY --from=builder /go/src/github.com/abhijitWakchaure/mock-api-server-go/dist/mock-api-server-go-linux-amd64 .
EXPOSE 8080
CMD ["./mock-api-server-go-linux-amd64"]