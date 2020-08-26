FROM golang:alpine as builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /main

FROM golang:latest
COPY --from=builder /main /main
RUN ls -l /
ENTRYPOINT ["/main"]
