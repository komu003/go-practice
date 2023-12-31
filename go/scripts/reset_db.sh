#!/bin/bash

migrate -database 'mysql://root:rootpassword@tcp(mysql)/mydatabase?query' -path migrations drop -f
migrate -database 'mysql://root:rootpassword@tcp(mysql)/mydatabase?query' -path migrations up
migrate -database 'mysql://root:rootpassword@tcp(mysql)/test_mydatabase?query' -path migrations drop -f
migrate -database 'mysql://root:rootpassword@tcp(mysql)/test_mydatabase?query' -path migrations up
go run cmd/seed/main.go
