FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o go-demo-backend

EXPOSE 8080

CMD ./go-demo-backend