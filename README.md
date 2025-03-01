# go-core-api

go-core-api is a base API framework written in Go that provides a solid, modular structure to jumpstart your API development projects. It is inspired by frameworks like Laravel (specifically its API structure) and leverages the Gin framework for HTTP routing.

## Project Overview

The **go-core-api** project is designed to serve as a foundational template for building robust and scalable APIs in Go. Its key features include:

- **API and CLI Modes**: Run the project as a RESTful API or a command-line tool based on startup arguments.
- **Routing & Middleware**: Centralized routing configuration with support for middlewares (e.g., JWT authentication).
- **Request Validation**: Centralized validation logic in the requests layer.
- **Modular Architecture**: Clear separation of concerns with dedicated packages for handlers, services, repositories, entities, mappers, configuration, logging, and utilities (Utils).
- **Extendable & Reusable**: Designed to be easily extended with your business logic, making it perfect as a template for future projects.

## Features

- **Basic API Structure**: Includes routing, middleware, and request/response handling.
- **JWT Authentication (Mocked)**: Basic JWT middleware that validates the presence of a token.
- **CLI Tool (Under Construction)**: A placeholder CLI for auxiliary development tasks.
- **Modular & Scalable**: Organized into clearly defined layers, facilitating maintenance and scalability.
- **Environment Configuration**: A dedicated configuration module for managing environment variables.

## Project Structure

```
go-core-api/
├── api/
│   ├── api.go      # Defines the API route structure
│   ├── middlewares 
│   ├── requests
│   ├── resources
│   ├── route.go    # Converts route definitions to Gin format
│   └── server.go   # Starts the HTTP server
├── cmd/
│   └── cli
│       └── cli.go  # CLI tool implementation
├── config/
├── docs/
├── entities/
├── logger/
├── mappers/
├── policies/
├── repositories/
├── services/
├── utils/
└── main.go        # Entry point: selects API or CLI mode based on arguments
```

## Installation

Clone the repository:

```bash
git clone https://github.com/mariombn-io/go-core-api.git
cd go-core-api
```

## Running in Development

### API Mode

To run the API in development mode, execute:

```bash
go run main.go --api --host 0.0.0.0 --port 8088
```

This command starts the server on the specified host and port. If no arguments are provided, a help message is displayed.

### CLI Mode

To run the CLI tool (currently under construction), execute:

```bash
go run main.go --cli
```

## Building the Project

To compile the project into an executable binary, run:

```bash
go build -o go-core-api main.go
```

Then you can run the binary:

```bash
./go-core-api --api --host 0.0.0.0 --port 8088
```

## TODO

- [X] Basic project structure
- [x] Configuration of routes, middlewares, handlers, and requests
- [ ] Setup of development environment, Docker, and Docker Compose
- [ ] Utils for handlers with functions for standardizing success and error responses
- [ ] Implementation of a configuration file with environment settings
- [ ] Implementation of the database (Entities, Mappers, Repositories, Migrations, and Seeders)
- [ ] Implementation of an example for the services layer
- [ ] Implementation of JWT integrated with the user entity
- [ ] Implementation of logs Utils
- [ ] Implementation of cache Utils (probably integrated with Redis)
- [ ] Implementation of the CLI tool
- [ ] Tests
- [ ] Documentation and README

Once everything is completed, we will have release **1.0.0**.

## License

This project is licensed under the [MIT License](LICENSE).