FROM node:20-alpine as node_buidler

COPY . /app/

WORKDIR /app/frontend
RUN npm install
RUN npm run build

FROM golang:1.22.5-alpine as golang_builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY --from=node_buidler /app/ .

RUN GOOS=linux go build -o main .

FROM alpine:latest

RUN apk add ca-certificates

WORKDIR /app

COPY --from=golang_builder /app/main /app/

CMD ["/app/main"] 