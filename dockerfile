FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /hellochi

FROM debian:bullseye-slim

WORKDIR /

COPY --from=build /hellochi /hellochi

EXPOSE 8080

CMD ["/hellochi"]