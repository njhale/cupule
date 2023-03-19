# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /src

COPY . .

RUN go build -o /cupule


EXPOSE 9000

ENTRYPOINT ["/cupule"]

