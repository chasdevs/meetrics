# Multi-stage build to get a lean go container: https://docs.docker.com/engine/userguide/eng-image/multistage-build

FROM golang as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY pkg ./pkg
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -a -installsuffix cgo -o meetrics .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/meetrics .
COPY config ./config
CMD ["./meetrics"]
