FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

COPY . $GOPATH/src/simpleapi/
WORKDIR $GOPATH/src/simpleapi/

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /go/src/simpleapi/main /app/
WORKDIR /app
ENTRYPOINT ["/app/main"]
EXPOSE 8081:8081
