# GO Spikes
Using Go version `1.17.10`

## SQL
SQLite3

Installed [Link](https://github.com/mattn/go-sqlite3)
`go install  github.com/mattn/go-sqlite3`

- [x] CRUD operations however no wrapper for sql commands written in strings. (Potentially another library to cover this)
Table creation, schema changing available.

- [x] Migration libraries available

## NOSQL
MongoDb
Installed [Link](https://github.com/mongodb/mongo-go-driver)
`go get go.mongodb.org/mongo-driver/mongo`

Pre-req
```
brew tap mongodb/brew
brew update
brew install mongodb-community@5.0
```
To run Mongodb 
`mongod --config /opt/homebrew/etc/mongod.conf`

- [x] CRUD operations.
- [x] Migration libs available.

## Command Executions
- [x] Shell
- [x] CLI
- [x] Read outputs/errors

## AWS SDK
Installed [Link](https://github.com/aws/aws-sdk-go)
`go get github.com/aws/aws-sdk-go`

- [x] Upload
- [x] Percentage Upload 
- [x] Errors
- [x] Download with Percentage [Link](https://github.com/aws/aws-sdk-go/blob/main/example/service/s3/getObjectWithProgress/getObjectWithProgress.go)

Transfer manager available in `aws-sdk-go-v2` seems quite new. [Link](https://github.com/aws/aws-sdk-go-v2)