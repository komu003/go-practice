# ./go/Dockerfile

FROM golang:latest

RUN go install github.com/cosmtrek/air@v1.29.0
CMD ["air", "-c", ".air.toml"]