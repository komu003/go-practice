#!/bin/bash

openapi-generator-cli generate -i /api-spec/openapi.yml -g go-server -o /app/gen --additional-properties=outputAsLibrary=true
