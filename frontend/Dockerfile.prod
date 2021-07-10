FROM node:slim as builder

WORKDIR /training

COPY package*.json ./
COPY yarn.lock ./
RUN yarn install

COPY tsconfig.json .
COPY src/ ./src
COPY public/ ./public

ARG REACT_APP_SERVER_URI
ENV REACT_APP_SERVER_URI=$REACT_APP_SERVER_URI
RUN yarn build


FROM nginx:alpine-perl
WORKDIR /usr/share/nginx/html
RUN rm -rf ./*
COPY --from=builder /training/build .
COPY ./nginx/nginx.conf /etc/nginx/conf.d/default.conf
ENTRYPOINT ["nginx", "-g", "daemon off;"]