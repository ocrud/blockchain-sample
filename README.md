# Sample Bitcoin Blockchain

Programming a simple blockchain network with node and wallet building.

## Run Blockchain Server

```sh
cd blockchain_server
go run .
OR
go run blokchain_server.go main.go
```
- Default port 5000
- Run multiple blockchain nodes with different ports
```sh
go run . -port 5001
go run . -port 5002
```
Open API url: [localhost:5000](http://localhost:5000)

| API URL | Method | Description |
| ------------- | ------------- | ------------- |
| localhost:5000/chain | GET  | Get all chain of blockchain  |
| localhost:5000/mine  | GET | Manual mining  |
| localhost:5000/mine/start | GET  | Set auto mining  |
| localhost:5000/transactions | GET/POST/PUT/DELETE | Get all transactions pool  |
| localhost:5000/amount?address=  | GET | Get amount of address  |
| localhost:5000/consensus | PUT | Set auto blockchain consensus  |

## Run Wallet Server

```sh
cd wallet_server
go run .
OR
go run wallet_server.go main.go
```
- Default port 8080, gateway localhost:5000
- Run multiple wallets with different ports and gateways
```sh
go run . -port 8081 -gateway http://localhost:5000
go run . -port 8082 -gateway http://localhost:5001
```

Open wallet url: [localhost:8080](http://localhost:8080)