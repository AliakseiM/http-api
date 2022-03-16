FROM golang:1.17-alpine as build
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /bin/app cmd/main.go
RUN apk --no-cache add ca-certificates

FROM scratch
COPY --from=build /bin/app .
COPY --from=build /etc/ssl/certs /etc/ssl/certs
ENTRYPOINT ["/app"]
