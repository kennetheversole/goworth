FROM golang:1.15 as build

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM ubuntu:20.04
WORKDIR /src
COPY . .
COPY --from=build /src/app .
EXPOSE 8081
CMD ["./app"]
