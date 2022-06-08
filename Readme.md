# GO Spikes
Using Go version `1.17.10`

## SQL
SQLite3

Installed
`go install  github.com/mattn/go-sqlite3`

- [x] CRUD operations however no wrapper for sql commands written in strings. (Potentially another library to cover this)
Table creation, schema changing available.

- [x] Migration libraries available

## NOSQL
MongoDb
Installed
`go get go.mongodb.org/mongo-driver/mongo`

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