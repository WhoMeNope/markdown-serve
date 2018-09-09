# Builder
FROM golang:1.10.3 as builder

WORKDIR /go/src/github.com/WhoMeNope

RUN go get -d -v -u github.com/gomarkdown/markdown \
 && go get -d -v -u github.com/valyala/fasthttp

COPY . .

RUN GOPATH=/go GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

# Deploy
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root
COPY --from=builder /go/src/github.com/WhoMeNope/app .
COPY --from=builder /go/src/github.com/WhoMeNope/core.css .

EXPOSE 3000
ENTRYPOINT ["./app"]
