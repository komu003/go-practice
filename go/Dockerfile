# ./go/Dockerfile

FROM golang:latest

RUN echo "Asia/Tokyo" > /etc/timezone

RUN go install github.com/cosmtrek/air@v1.29.0
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# OpenAPI Generatorのインストール
RUN curl -L https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/7.0.1/openapi-generator-cli-7.0.1.jar -o /usr/local/bin/openapi-generator-cli.jar
RUN echo '#!/bin/sh\njava -jar /usr/local/bin/openapi-generator-cli.jar "$@"' > /usr/local/bin/openapi-generator-cli
RUN chmod +x /usr/local/bin/openapi-generator-cli

CMD ["air", "-c", ".air.toml"]
