# GO Spikes
Using Go version `1.17.10`
Latest version which is supported by the OS

## SQL
SQLite3

Installed
`go install  github.com/mattn/go-sqlite3`

CRUD operations however no wrapper for sql commands written in strings. (Potentially another library to cover this)
Table creation, schema changing available.

Migration libraries available.

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

CRUD operations.

Migration libs available.

