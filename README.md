# Wallet #
Simple HTTP API allows to make money operations on user wallet.
Database concurrent synchronization is done using optimistic locking.
## How to run the project ##

### Setup postgres db ###
```
docker-compose up -d
```

### Fetch deps ###
```
go mod download
```

### Run the project ###
```
go run ./cmd/wallet/main.go
```

### Run tests ###
```
go test ./...
```
