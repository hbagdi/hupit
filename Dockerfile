FROM golang:1.12 AS build
WORKDIR /hupit
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hupit

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /hupit/hupit /usr/local/bin
ENTRYPOINT ["hupit"]
