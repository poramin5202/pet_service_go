FROM golang:1.21.0-alpine3.18 AS build

WORKDIR /app

COPY . .

RUN go build -o api

#build small image
FROM alpine:3.18
WORKDIR /app

COPY --from=build /app .

EXPOSE 8085

CMD [ "./api" ]