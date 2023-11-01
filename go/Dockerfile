# ./go/Dockerfile

FROM golang:latest

RUN echo "Asia/Tokyo" > /etc/timezone

RUN go install github.com/cosmtrek/air@v1.29.0
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
CMD ["air", "-c", ".air.toml"]
