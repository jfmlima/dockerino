FROM golang:1.13 as build-proxy
WORKDIR /usr/src/app/proxy
COPY ./proxy ./
RUN go build -o ./proxy ./cmd
CMD ["./proxy"]