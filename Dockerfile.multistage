FROM golang:1.12-alpine as builder
WORKDIR /welcome
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags="-w -s" -o app
FROM scratch
COPY --from=builder /welcome/app /welcome
CMD ["./welcome","run"]