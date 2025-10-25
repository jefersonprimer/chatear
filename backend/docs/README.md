# Chatear Backend

This is the backend service for the Chatear application, built with Go. It provides a robust and scalable foundation for real-time communication features, user management, and notifications.

## Table of Contents

*   [Prerequisites](#prerequisites)
*   [Setup](#setup)
*   [Running the Application](#running-the-application)
*   [Running Tests](#running-tests)
*   [Documentation](#documentation)

## Prerequisites

Before you begin, ensure you have the following installed:

*   **Go**: Version 1.25.3 or higher.
*   **Docker**: To run the application stack.
*   **Docker Compose**: To orchestrate the Docker containers.
*   **[golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate):** To run database migrations.

## Setup

1.  **Clone the repository**:

    ```bash
    git clone https://github.com/jefersonprimer/chatear-backend.git
    cd chatear-backend
    ```

2.  **Download Go modules**:

    ```bash
    go mod download
    ```

3.  **Configure environment variables**:

    Copy the example environment file and update it with your specific configurations for your Supabase database, Redis, NATS, JWT secrets, and SMTP settings.

    ```bash
    cp env.example .env
    # Open .env in your editor and fill in the details
    ```

4.  **Run database migrations**:

    Before starting the application, you need to apply the database migrations to your Supabase database.

    Run the following command, replacing your Supabase connection string:

    ```bash
    migrate -database "YOUR_SUPABASE_CONNECTION_STRING" -path migrations/postgres up
    ```

## Running the Application

The easiest way to run the entire application stack (API, workers, NATS, Redis) is using `docker-compose`.

1.  **Start the containers**:

    ```bash
    docker-compose -f docker-compose.events.yml up --build
    ```

    This will build the images and start all the services.

2.  **Check the status**:

    ```bash
    docker-compose -f docker-compose.events.yml ps
    ```

The API will be available at `http://localhost:8080`.

## Running Tests

To run all unit and integration tests for the project:

```bash
go test ./...
```

To run tests with verbose output:

```bash
go test -v ./...
```

## Documentation

Detailed documentation for the project's architecture, dependencies, environment variables, and domain-specific workflows can be found in the `docs/` directory:

*   [Quick Start Guide](docs/iniciar.md)
*   [Architecture Overview](docs/architecture.md)
*   [Project Dependencies](docs/dependencies.md)
*   [User Domain](docs/user_domain.md)
*   [Notification Domain](docs/notification_domain.md)