FROM golang:1.13 as build-proxy
WORKDIR /usr/src/app/proxy
COPY ./proxy ./
WORKDIR /usr/src/app/proxy/cmd
RUN go build -o ../proxy
WORKDIR /usr/src/app/proxy/
CMD ["./proxy"]