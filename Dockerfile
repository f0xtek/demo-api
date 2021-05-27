FROM golang:alpine AS builder
WORKDIR /
COPY go.mod .
COPY main.go .
RUN GOOS=linux CGO_ENABLED=0 go build -o /demo main.go

FROM scratch
COPY --from=builder /demo /
EXPOSE 8090
CMD ["/demo"]