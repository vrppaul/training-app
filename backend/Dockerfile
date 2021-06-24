FROM golang:1.16-alpine3.13 AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/training-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/training-app .


FROM alpine:3.9 
COPY --from=build_base /tmp/training-app/out/training-app /app/training-app
EXPOSE 8080
CMD ["/app/training-app"]