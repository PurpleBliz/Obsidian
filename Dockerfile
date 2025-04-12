FROM golang:1.24-alpine AS build
LABEL authors="purplebliz"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /Obsidian ./cmd

FROM alpine:latest
COPY --from=build /Obsidian /obsidian

EXPOSE 8080

CMD ["/obsidian"]