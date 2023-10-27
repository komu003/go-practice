FROM golang:latest

RUN apt-get update && apt-get install -y default-jre

RUN wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.2.0/openapi-generator-cli-5.2.0.jar -O /usr/local/bin/openapi-generator-cli && chmod +x /usr/local/bin/openapi-generator-cli

WORKDIR /app
