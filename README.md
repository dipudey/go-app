# GO APP

## Overview

**GO APP** is a modular Go application designed with isolated internal modules, clean architecture, and scalable structure.  
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
| | └── router.go # module route definitions
│ │ └── module.go # Register(router, db) resolves dependencies and calls router
│ ├── router/ # Central router
│ │ └── router.go # Router initialization and module registration
│── config/ # App configuration
│   ├── application.go # Application initialization
│   ├── config.go # Configuration management
│   └── db.go # Database connection setup
```

---

## Swagger API Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) and [gin-swagger](https://github.com/swaggo/gin-swagger) for OpenAPI documentation.

### How to Generate Swagger Docs

1. Ensure your handler functions and main router file have proper Swagger annotations. Example:
    ```go
    // @title GO APP API
    // @version 1.0
    // @description This is a standard GO APP API server.
    // @termsOfService http://swagger.io/terms/
    // @contact.name API Support
    // @contact.url http://www.swagger.io/support
    // @contact.email support@swagger.io
    // @license.name Apache 2.0
    // @license.url http://www.apache.org/licenses/LICENSE-2.0.html
    // @host localhost:8080
    // @BasePath /
    ```
    And for endpoints:
    ```go
    // @Summary Ping the server
    // @Description Returns "pong" to test API
    // @Tags Health
    // @Accept json
    // @Produce json
    // @Success 200 {object} map[string]string
    // @Router /api/ping [get]
    ```
2. Run the following command to generate docs:
    ```sh
    $HOME/go/bin/swag init -g ./cmd/api/main.go -o ./docs
    ```
    - `-g ./cmd/api/main.go` points to your main entry file.
    - `-o ./docs` outputs the generated Swagger files to the docs directory.

### How to Access Swagger UI

- Start the application.
- Open your browser and go to: `http://localhost:8000/swagger/index.html`
- All documented API endpoints will be listed and testable.

### References
- [swaggo/swag](https://github.com/swaggo/swag)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [OpenAPI Specification](https://swagger.io/specification/)
