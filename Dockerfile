FROM golang:1.17
WORKDIR /go/src/prometheus-dingtalk-webhook
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app cmd/webhook/webhook.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/prometheus-dingtalk-webhook .
ENTRYPOINT ["./app"]