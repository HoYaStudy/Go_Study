# HoYa's Coin

## Initialize Project

> $ go mod init github.com/HoYaStudy/Go_Study/hcoin

Then, go.mod file is created.

## Bolt DB

> $ go get github.com/boltdb/bolt

### Browser

> $ go get github.com/evnix/boltdbweb

> $ boltdbweb --db-name=<db_name>.db

## Test Project
### Test All Files
> $ go test -v ./...

### Coverage
> $ go test -v -coverprofile cover.out ./...

> $ go tool cover -html=cover.out