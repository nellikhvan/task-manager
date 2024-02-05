# task-manager/backend/Dockerfile
FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/jinzhu/gorm/dialects/postgres

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
