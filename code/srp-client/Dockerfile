FROM golang:alpine as builder

WORKDIR /go/src/app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main -mod=vendor

FROM golang:latest
COPY --from=builder /main /main
RUN ls -l /
ENTRYPOINT ["/main"]
