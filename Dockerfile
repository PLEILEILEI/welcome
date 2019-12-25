FROM golang:1.12-alpine
RUN apk --no-cache add ca-certificates

WORKDIR /welcome
COPY / .
RUN go build -mod=vendor -o welcome

WORKDIR /welcome/
CMD ["./welcome","run"]