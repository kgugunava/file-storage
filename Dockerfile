FROM golang:1.23-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server cmd/main/main.go


FROM ubuntu
LABEL authors="kgugunava"

COPY --from=build ./app/server ./server

EXPOSE 8010
CMD ./server