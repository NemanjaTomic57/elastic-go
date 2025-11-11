# elastic-go

A lightweight Go utility for running and managing a local Elasticsearch instance — perfect for development and testing.

## 🚀 Features
- Simple CLI to start Elasticsearch locally  
- Environment configuration with `.env`  
- Minimal dependencies and clean code structure  
- Easily extendable for different setups

## 📁 Project Structure
```
elastic-go/
├── cmd
│   ├── api
│   │   └── main.go
│   └── indexer
│       └── main.go
├── elastic-start-local
│   ├── config
│   │   └── telemetry.yml
│   ├── docker-compose.yml
│   ├── start.sh
│   ├── stop.sh
│   └── uninstall.sh
├── go.mod
├── go.sum
├── internal
│   ├── db
│   │   └── db.go
│   ├── es
│   │   └── client.go
│   └── indexer
│       └── verzeichnisse.go
└── start-local
```

## ⚙️ Requirements
- Go 1.20+  
- Docker (for running Elasticsearch container)
