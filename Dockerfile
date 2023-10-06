#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v .
ENTRYPOINT [ "./app" ]

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=routesms Version=0.0.1
EXPOSE 6000
