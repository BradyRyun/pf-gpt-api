FROM golang:1.20.4-alpine as build

LABEL maintainer="Brady Ryun <brady@ryunengineering.com>"

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

FROM golang:1.20.4-alpine as runtime

WORKDIR /app

COPY --from=build /app/main ./

EXPOSE 8080

CMD ["/app/main"]
