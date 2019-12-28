# go-test-example

[more details](https://bastengao.com/blog/2019/12/go-test-practices.html)

## Setup

init db

    sqlite3 test.db < init.sql

generate mock

    go get github.com/golang/mock/mockgen
    mockgen -destination mock/mailer.go -package mock github.com/bastengao/go-testexample Mailer

run test

    go test
