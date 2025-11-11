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
├── cmd/
│   └── elastic-start-local   # CLI entry point
├── internal/                 # Internal helpers and logic
├── start-local/              # Local setup scripts
├── .env                      # Environment configuration
├── go.mod
└── go.sum
```

## ⚙️ Requirements
- Go 1.20+  
- Docker (for running Elasticsearch container)
