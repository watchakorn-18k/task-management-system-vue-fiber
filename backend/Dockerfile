FROM golang:1.23-alpine3.20 as builder

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app
COPY . .

COPY ./src/docs/swagger.yaml /app/src/docs/swagger.yaml

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

FROM scratch
WORKDIR /
COPY --from=builder /main /main
COPY --from=builder /app/src/docs /src/docs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 1818
ENTRYPOINT ["/main"]

