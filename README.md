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
├── cmd/                               # Entry points for different executables (main packages)
│   ├── api/                           # Contains the main file for starting the API service
│   │   └── main.go                    # Bootstraps the API (e.g., starts HTTP server, connects to Elasticsearch)
│   └── indexer/                       # Contains the main file for running indexing jobs or CLI tools
│       └── main.go                    # Runs the indexer logic (e.g., populating Elasticsearch with data)
│
├── elastic-start-local                # Executable or script for starting Elasticsearch locally (helper tool)
│
├── go.mod                             # Go module definition (module path, dependencies, Go version)
├── go.sum                             # Checksums for dependencies (ensures reproducible builds)
│
├── internal/                          # Internal packages (not exposed outside this module)
│   ├── db/                            # Database layer (if any auxiliary data storage or persistence is used)
│   │   └── db.go                      # Contains logic for connecting or mocking DB (optional helper)
│   ├── es/                            # Elasticsearch client utilities
│   │   └── client.go                  # Client initialization and connection handling to Elasticsearch
│   └── indexer/                       # Indexing logic and domain-specific data processing
│       └── verzeichnisse.go           # Example indexer implementation (e.g., indexing "directories" or files)
│
└── start-local                        # Shell or Go script to start local environment (Docker + setup helpers)
```

## ⚙️ Requirements
- Go 1.20+  
- Docker (for running Elasticsearch container)
