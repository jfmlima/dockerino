
FROM node:current-alpine as build-frontend
WORKDIR /usr/src/app/fe
COPY ./fe ./
RUN yarn
RUN yarn build
RUN yarn global add serve
CMD ["serve", "-s", "build"]