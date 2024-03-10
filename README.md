# Wallet #

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
