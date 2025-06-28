FROM golang:1.22 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build dengan static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o /hellochi

FROM scratch

COPY --from=build /hellochi /hellochi

EXPOSE 8080

CMD ["/hellochi"]
