FROM node:current-alpine as build-frontend
WORKDIR /usr/src/app/fe
COPY ./fe ./
RUN yarn
RUN yarn build

FROM golang:1.13 as build-proxy
WORKDIR /usr/src/app/proxy
COPY ./proxy ./
RUN go build -o proxy ./cmd
COPY --from=build-frontend /usr/src/app/fe/build ./static
CMD ["./proxy"]