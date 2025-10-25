# Insomnia Test Guide for Chatear Backend

This guide provides instructions on how to set up and test the Chatear Backend using Insomnia.

## 1. Setup and Start the Services

Before running the application, ensure you have Docker and Docker Compose installed.

### 1.1. Environment Variables

Create a `.env` file in the root of the project based on `env.example`. Make sure to fill in all the necessary details.

```bash
cp env.example .env
# Open .env and fill in your details
```

**Important:**
*   **`SUPABASE_CONNECTION_STRING`**: Make sure this is correctly set to your Supabase database.
*   **`SMTP_USER` and `SMTP_PASS`**: If you are using Gmail, you will need to generate an App Password. Refer to Google's documentation on how to do this.

### 1.2. Start Docker Services

Navigate to the root of your project and start the Docker services (Postgres, Redis, NATS):

```bash
docker-compose -f docker-compose.events.yml up --build -d
```

This will start the `postgres`, `redis`, and `nats` containers in detached mode.

### 1.3. Run the Go Application

You can run the API and worker services separately.

#### 1.3.1. Run the API Service

Open a new terminal in the project root and run the API service:

```bash
go run cmd/api/main.go
```

#### 1.3.2. Run the Worker Services

Open other terminals in the project root and run the worker services you need:

```bash
go run cmd/worker/notification_worker.go
go run cmd/worker/user_registered_worker.go
# ... and so on for other workers
```

## 2. Insomnia Setup

1.  **Create a new Request Collection** in Insomnia.
2.  **Set up Base URL:** In the collection's environment settings, create a variable `base_url` and set it to `http://localhost:8080`.

## 3. REST API Testing

### 3.1. Register User

*   **Method:** `POST`
*   **URL:** `{{ _.base_url }}/api/v1/register`
*   **Body (JSON):**
    ```json
    {
        "name": "Test User",
        "email": "test@example.com",
        "password": "password123"
    }
    ```

### 3.2. Login User

*   **Method:** `POST`
*   **URL:** `{{ _.base_url }}/api/v1/login`
*   **Body (JSON):**
    ```json
    {
        "email": "test@example.com",
        "password": "password123"
    }
    ```
*   **Response:**
    ```json
    {
        "accessToken": "...",
        "refreshToken": "..."
    }
    ```

## 4. GraphQL API Testing

For GraphQL, you will use a single endpoint for all your queries and mutations.

*   **Method:** `POST`
*   **URL:** `{{ _.base_url }}/graphql`
*   **Body:** Use the "GraphQL" body type in Insomnia.

### 4.1. Register User (GraphQL)

*   **Query:**
    ```graphql
    mutation RegisterUser($input: RegisterUserInput!) {
      registerUser(input: $input) {
        id
        name
        email
      }
    }
    ```
*   **Variables (JSON):**
    ```json
    {
      "input": {
        "name": "GraphQL User",
        "email": "graphql@example.com",
        "password": "password123"
      }
    }
    ```

### 4.2. Login (GraphQL)

*   **Query:**
    ```graphql
    mutation Login($input: LoginInput!) {
      login(input: $input) {
        accessToken
        refreshToken
        user {
          id
          name
        }
      }
    }
    ```
*   **Variables (JSON):**
    ```json
    {
      "input": {
        "email": "graphql@example.com",
        "password": "password123"
      }
    }
    ```

### 4.3. Logout (GraphQL)

*   **Query:**
    ```graphql
    mutation Logout {
      logout
    }
    ```
*   **Authentication:** You need to provide the `accessToken` from the login mutation as a Bearer token in the `Authorization` header.

## 5. Clean Up

To stop and remove the Docker containers:

```bash
docker-compose -f docker-compose.events.yml down
```
