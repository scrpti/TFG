# Medical Ledger Backend

Backend service for managing medical documents with a focus on data integrity, traceability, and a scalable architecture prepared for future blockchain integration.

---

## Overview

This project implements a RESTful API in Go for handling medical-related data such as patients, users, and documents. The system is designed following a layered architecture that separates concerns between HTTP handling, business logic, and data persistence.

The main goal of this project is to provide a robust foundation for:

- Managing medical entities (patients, users, documents)
- Ensuring data integrity using cryptographic hashing
- Storing sensitive data off-chain
- Preparing the system for future integration with a blockchain layer (e.g., Paladin)

---

## Architecture

The project follows a modular structure:

cmd/api/           → Application entry point  
internal/config/   → Configuration management  
internal/db/       → Database connection  
internal/models/   → Domain models  
internal/repository/ → Data access layer (SQL)  
internal/service/  → Business logic  
internal/handlers/ → HTTP layer  
internal/routes/   → Route definitions  
internal/storage/  → File storage logic  
internal/crypto/   → Hashing utilities  
internal/ledger/   → Ledger abstraction (future blockchain integration)

### Layered Flow

HTTP Request → Handler → Service → Repository → Database

---

## Tech Stack

- Language: Go
- Framework: Gin
- Database: PostgreSQL
- Database access: database/sql
- Configuration: Environment variables + godotenv
- Hashing: SHA-256 (planned)
- Storage: Local file system (planned)

---

## Features (Current)

- Project structure and modular architecture
- Configuration management via .env
- PostgreSQL connection
- Health check endpoint

---

## Features (In Progress)

- Patient management (CRUD)
- User management
- Document storage and metadata management
- File storage system
- Hash-based integrity verification

---

## Planned Features

- Document integrity verification using SHA-256
- Off-chain storage for medical files
- Ledger abstraction layer for recording document evidence
- Integration with a blockchain middleware (e.g., Paladin)

---

## Getting Started

### Prerequisites

- Go (1.20+ recommended)
- PostgreSQL
- Git

---

### Environment Variables

Create a .env file in the root directory:

APP_PORT=8080

DB_HOST=localhost  
DB_PORT=5432  
DB_USER=postgres  
DB_PASSWORD=postgres  
DB_NAME=medical_ledger  
DB_SSLMODE=disable  

UPLOAD_DIR=uploads

---

### Database Setup

Create the database:

CREATE DATABASE medical_ledger;

---

### Run the application

go run cmd/api/main.go

The server will start on:

http://localhost:8080

---

### Health Check

GET /health

Response:

{
  "status": "ok"
}

---

## Development Notes

- The project uses database/sql for explicit control over queries and data mapping.
- The architecture is designed to be extensible and maintainable.
- Business logic is isolated in the service layer.
- Future blockchain integration is abstracted via a ledger interface.

---

## Project Status

This project is currently under active development as part of a backend-focused system with future extensions toward secure, traceable medical data management.

---

## License

This project is intended for academic and educational purposes.