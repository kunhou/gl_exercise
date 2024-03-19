# Coding Exercise

[![Go Version](https://img.shields.io/badge/Go-v1.17.8-blue.svg)](https://go.dev/doc/go1.17)
[![Test](https://github.com/kunhou/gl_exercise/actions/workflows/test.yml/badge.svg)](https://github.com/kunhou/gl_exercise/blob/master/.github/workflows/test.yml)
[![Docker Build](https://img.shields.io/github/actions/workflow/status/kunhou/gl_exercise/build-image.yml)](https://hub.docker.com/repository/docker/patrice43/task-server)

This repository implements a RESTful API for managing tasks. It allows users to create, list, update, and delete tasks. The application is designed to run within a Docker container, facilitating easy deployment.

## Quick Start

1. **Start the Service:**

   ```bash
   docker-compose -f docker/docker-compose.yml up -d
   ```

   This command starts the Task List API service in a detached mode (-d).

2. **Access Swagger UI:**

   Open http://localhost:8080/swagger/index.html to explore the API endpoints and interact with the service.

3. **Stop the Service:**

   ```bash
   docker-compose -f docker/docker-compose.yml down
   ```

## Development Guide

### Prerequisites

- Go v1.17.8
- Wire
  ```shell
  go install github.com/google/wire/cmd/wire@v0.5.0
  ```
- Swaggo
  ```shell
  go install github.com/swaggo/swag/cmd/swag@v1.7.8
  ```

### Make Commands

The provided Makefile offers convenient commands for managing the project:

- **`make run`:** Launches the Task List API server locally.
- **`make test`:** Executes unit tests for the application.
- **`make gen-doc`:** Generates the Swagger documentation (useful before pushing changes).

### Code Folder Structure Explanation

The codebase follows a clear organization:

- `cmd`: Contains the main entry point and related code.
- `internal` (Core functionality):
  - `common`: Common definitions and configurations.
  - `deliver`: Defines interfaces for external communication (currently HTTP, but could be extended to gRPC etc.).
  - `entity`: Contains object model definitions.
  - `pkg`: Stores reusable packages.
  - `repository`: Handles database interactions (without business logic).
  - `service`: The primary location for business logic.

### CI/CD Flow

The project leverages GitHub Actions for CI/CD:

- Pushing changes to the `main` branch automatically triggers tests.
- Pushing a tag starting with `v1.x` triggers the build and publication of a Docker image to [Docker Hub](https://hub.docker.com/repository/docker/patrice43/task-server).

## API Reference

For detailed API documentation, visit the Swagger UI at http://localhost:8080/swagger/index.html. This interactive interface allows you to explore the API's functionalities and issue test requests.

The API offers the following actions:

- **GET /tasks:** Retrieves a list of all tasks.
- **POST /task:** Creates a new task with a specified name.
- **PUT /task/:id:** Updates an existing task by providing its ID and optionally its name and status.
- **DELETE /task/:id:** Deletes a task by its ID.

## Reference

- [SPEC](document/spec.md)