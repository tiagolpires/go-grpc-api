# Go gRPC API
gRPC API developed with Go

## ⚡️ Start

Download and install [Go](https://go.dev/dl/)

Run docker compose
```
docker-compose up
```
Run seeds and migrations
```
make seed
```
Start the server
```
make start
```
In a new terminal, start the client
```bash
make start_client

# Expected output: 
# UpvoteCrypto...
# Crypto: id:1 name:"Bitcoin" code:"BTC" votes:6 
#
# DownvoteCrypto...
# Crypto: id:1 name:"Bitcoin" code:"BTC" votes:5 
#
# Start crypto streaming...
# Bitcoin Votes: 5...
# Finish crypto streaming
#
# UpdateCrypto...
# Crypto: id:1 name:"Bitcoin" code:"BTC" votes:5 
#
# GetCryptoById...
# Crypto: id:1 name:"Bitcoin" code:"BTC" votes:5 
#
# ListCryptos...
# List: [id:4 name:"Ethereum" votes:10...]
```
## Tests
To run tests
```
make test
```
## Generate Proto Buffer file
```
make gen
```
## Built with
- [Go](https://go.dev/)
- [gRPC](https://grpc.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [GORM](https://gorm.io/index.html)
