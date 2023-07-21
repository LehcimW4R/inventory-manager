FROM golang:1.20.6-alpine3.18

RUN mkdir -p app/

COPY . /app

EXPOSE 8080

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]
