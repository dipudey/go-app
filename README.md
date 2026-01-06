# POS

## Overview

**POS** is a modular Go application designed with isolated internal modules, clean architecture, and scalable structure.  
Each module contains its own routes, services, repositories, models, DTOs, and validators.

The application uses a **single database connection** injected into modules, supports multiple DB drivers (Postgres, MySQL), and a minimal central router for module registration.

---

## Folder Structure
```aiignore
myapp/
├── go.mod # Go modules definition
|-- go.sum # Go modules checksums
├── cmd/
│ ├── api/
│ │ └── main.go # Application entry point
├── pkg/ # Public packages (if any)
├── internal/ # All internal modules and app logic
│ ├── user/ # User module
│ │ ├── handler.go # HTTP handlers for user endpoints
│ │ ├── service.go # Business logic for users
│ │ ├── repository.go # Database access for users
│ │ ├── model.go # User model definitions
│ │ ├── dto.go # Request and response structures (DTOs)
│ │ ├── validator.go # User-related validation logic
│ │ └── module.go # Register(router, db)
│ ├── router/ # Central router
│ │ └── router.go # Router initialization and module registration
│── config/ # App configuration
│   ├── application.go # Application initialization
│   ├── config.go # Configuration management
│   └── db.go # Database connection setup
```

