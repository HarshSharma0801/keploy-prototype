# Inventory Service (Keploy Prototype)

This is a simulated microservice (like the ecom-service) for demonstrating contract testing with Keploy. It showcases advanced features and can be used for testing schema generation, validation, versioning, and more.

## Features
- Generate and save inventory-related schemas (YAML/JSON)
- Versioned schema storage
- Unified schema management (merge, diff)
- Advanced validation (dependency graph, real-time validation)
- Provider-driven contract publishing and rollback

## Directory Structure
```
inventory-service/
├── v1/
│   └── tests/
│       └── contracts/
│           ├── provider/
│           └── consumer/
```

## Usage
Run the following commands from the root directory:

### Generate Schemas
```
go run main.go generate --service inventory-service
```

### Validate Contracts
```
go run main.go validate --service inventory-service --mode provider
```

### Merge Schemas
```
go run main.go merge --service inventory-service --inputs ...
```

### Diff Schemas
```
go run main.go diff --service inventory-service --old ... --new ...
```

---

## Example Inventory Test Case
- `test-get-inventory.yaml` (provider)
- `mock-get-inventory.yaml` (consumer)

---

## Notes
- This is a simulation for GSOC’25 proposal enhancements.
- See root README for project-wide documentation.
