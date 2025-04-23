# Keploy Contract Testing Prototype 


![keploy-logo](https://github.com/user-attachments/assets/6df71537-5146-4bc0-aced-4fb12fb8afab)


## Overview
Keploy is a contract testing tool designed to streamline API contract validation for microservices. This enhanced prototype demonstrates advanced features as proposed for including schema versioning, unified schema management, advanced validation, and provider-driven contract testing.

---

## Features

### Current Features
- **Schema Generation**: Converts HTTP interactions to OpenAPI schemas.
- **Schema Saving**: Stores schemas in a structured directory (YAML).
- **Validation**: Compares consumer mocks and provider tests, outputs color-coded results.
- **CLI**: Powered by Cobra, with `generate`, `validate`, and `download` commands.

### Simulated/Planned Enhancements
- **Schema Versioning**: Track version, timestamp, and author for each schema.
- **Multi-Format Support**: Save and load schemas in both YAML and JSON.
- **Mock S3 Integration**: Abstract storage backend (local/S3).
- **Unified Schema Management**: Merge multiple schemas, resolve conflicts, and show diffs.
- **Advanced Validation**: Dependency graph, real-time watch mode, deep body comparison.
- **Provider-Driven Architecture**: Contract publishing, rollback, and history tracking.

---

## Directory Structure
```
keploy-prototype/
├── main.go           # CLI entrypoint, commands
├── contract.go       # Schema structs, conversion, sample data
├── validation.go     # Schema comparison and validation logic
├── README.md         # Project documentation
└── ...
```

---

## Usage

### 1. Generate Schemas
```
go run main.go generate
```
- Saves provider and consumer schemas to `ecom-service/v1/tests/contracts/`

### 2. Validate Contracts
```
go run main.go validate --mode consumer
# or
./keploy validate --mode provider
```
- Compares schemas, outputs results in color-coded table.

---

## Simulated Enhancements (with Examples)

### Schema Versioning
- **How it works**: Each schema file includes metadata (version, timestamp, author).
- **Example**: `test-get-products-v1.1.yaml` with version history retained.
- **Simulated usage**:
  ```bash
  ./keploy generate --author "Harsh Sharma" --version v1.1
  ```

### Multi-Format Support
- **How it works**: Save/load schemas as YAML or JSON.
- **Example**:
  ```bash
  ./keploy generate --format json
  ```

### Unified Schema Management
- **How it works**: Merge multiple schemas, resolve conflicts, and show diffs.
- **Example**:
  ```bash
  ./keploy merge --inputs test-get-products.yaml mock-get-products.yaml --output product-service-v1.yaml
  ./keploy diff --old product-service-v1.yaml --new product-service-v1.1.yaml
  ```

### Advanced Schema Comparison
- **Dependency Resolution**: Build a dependency graph for inter-service validation.
- **Real-Time Validation**: Watch mode for instant feedback on schema changes.
- **Example**:
  ```bash
  ./keploy validate --watch
  ```

### Provider-Driven Architecture
- **Contract Publishing**: Share contracts with dependent teams/services.
- **Automated Rollback**: Revert to previous schema versions on failure.
- **Example**:
  ```bash
  ./keploy publish --contract inventory-v1.1.yaml
  ./keploy rollback --contract inventory-v1.0.yaml
  ```

---

## Project Plan & Milestones

| Week | Dates        | Milestone/Task                                              |
|------|-------------|------------------------------------------------------------|
| 1-3  | Jun 2-22    | Integrate prototype, add versioning, multi-format support  |
| 4-6  | Jun 23-Jul 13| Merging, diffing, conflict resolution                      |
| 7-9  | Jul 14-Aug 3| Dependency graph, real-time validation                     |
| 10-11| Aug 4-17    | Provider-driven features, publishing, rollback             |
| 12-13| Aug 18-31   | End-to-end tests, docs, optimization, final polish         |

---

## Developer Notes
- All enhancements are either implemented, stubbed, or simulated with clear TODOs in the code.
- See code comments for integration points and further development guidance.

---

## Example Enhancement Stubs (see code for details)
- `saveSchemaWithMeta()` — Save schema with version and author metadata.
- `mergeSchemas()` — Merge multiple schemas, resolve conflicts.
- `diffSchemas()` — Show differences between schema versions.
- `watchSchemas()` — Real-time validation loop (stub).
- `publishContract()` — Simulate contract publishing (stub).
- `rollbackContract()` — Simulate rollback (stub).

---

## Contributing
See GSOC proposal for detailed motivation, roadmap, and references to related issues and PRs.

---

## License
[MIT](LICENSE)
