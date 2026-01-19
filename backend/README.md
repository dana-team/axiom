# Axiom Backend

A Go-based backend service that provides cluster management functionality through a REST API.

## Prerequisites

- Go 1.24.3 or later
- MongoDB instance
- Make (for using Makefile commands)

## Environment Variables

The following environment variables are required to run the application:

| Variable         | Description                                                   | Required |
|------------------|---------------------------------------------------------------|----------|
| `MONGO_URI`      | MongoDB connection string (e.g., `mongodb://localhost:27017`) | Yes |
| `DB_NAME`        | Name of the MongoDB database to use                           | Yes |
| `TEST_MONGO_URI` | url of the MongoDB database to use in tests                   | Yes |

## Setup

1. Clone the repository
2. Set up the required environment variables
3. Install dependencies:
   
```bash
   go mod download
```

## Running the Application

Use the Makefile to run common tasks (check the Makefile for available commands).

To run the application:
```bash 
  make run
```

## API Documentation

Once the application is running, you can access the Swagger UI documentation at:

http://localhost:8080/swagger/index.html

## Development

The project follows standard Go project layout conventions with the following structure:

- `cmd/`: Contains the application entry points
- `internal/`: Houses the private application code
    - `controllers/`: Request handlers and business logic
    - `middleware/`: HTTP middleware components
    - `routes/`: API route definitions
    - `types/`: Data structure definitions
    - `utils/`: Utility functions and database clients
