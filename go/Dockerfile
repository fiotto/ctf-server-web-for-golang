FROM golang:1.13.5
WORKDIR /app
RUN go get github.com/go-sql-driver/mysql
CMD bash -c "go build app.go && ./app"
EXPOSE 8080
