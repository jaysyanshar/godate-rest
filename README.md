# GoDate API
GoDate is a dating app which allows you to meet nearby people. This is a REST API project for GoDate app.

## Project Structure
The GoDate REST API project follows a standard project structure, which is organized as follows:

- `cmd/`: The entry point of the application.
  - `migration/main.go`: This is the entry point for database migration.
  - `http/main.go`: This is the entry point for REST API service.
- `config/`: This directory contains configuration file for the application that binds with `.env` file.
- `controllers/`: This directory contains controller functions which handle HTTP requests and responses.
- `internal/`: This directory contains internal implementation details of components in the project, like how database is implemented.
- `middlewares/`: This directory is responsible for organizing and storing all the middleware components used in the application.
- `models/`: This directory contains the data models used in the application.
- `repositories/`: This directory contains the database repositories for interacting with the data.
- `routes/`: This directory contains the route files for the application to define the various routes and their corresponding handlers.
- `services/`: This directory contains the business logic services for the application.
- `utils/`: This directory contains utility functions and helper methods.

Feel free to explore each directory to understand the project structure in more detail.

## Getting Started

To get started with the GoDate REST API project, follow these steps:

1. Clone the repository:

    ```shell
    git clone https://github.com/jaysyanshar/godate-rest.git
    ```

2. Change into the project directory:

    ```shell
    cd godate-rest
    ```

3. Download all dependencies:

    ```shell
    make download
    ```

### Unit Test

Run the following command to perform unit test in the project scope:

```shell
make test
```

### Database Migration

To migrate the database for the GoDate REST API project, follow these steps:

1. Configure the `.env` file with the necessary database credentials. Open the `.env` file located in the project's root directory and update the following variables with your database information:

    ```plaintext
    DB_DRIVER=your_database_driver
    DB_HOST=your_database_host
    DB_PORT=your_database_port
    DB_NAME=your_database_name
    DB_USER=your_database_user
    DB_PASSWORD=your_database_password
    ```

    The `DB_DRIVER` that are currently supported in this project:
    - `mysql`
    - `postgres`
    - `sqlite3`
    - `mssql`

2. Once the `.env` file is configured, open your terminal and navigate to the project's root directory.

3. Run the following command to perform the database migration:

    ```shell
    make migrate
    ```

    This command will execute the migration script located in the `cmd/migration/main.go` file, which will create the necessary database tables and schema.

4. After the migration is complete, you can verify the database changes by checking your database management tool or running queries against the database.

    **Note:** Make sure you have the necessary permissions and access rights to perform the database migration.


### REST API

To run the GoDate REST API project, follow these steps:

1. **(Optional)** To make a production-ready build, run this command:

    ```shell
    make build
    ```

2. Start the server using Makefile:

    ```shell
    make run
    ```

3. The server should now be running on `http://localhost:8080`. You can test the API endpoints using tools like cURL or Postman.

4. Refer to the API documentation for more details on the available endpoints and request/response formats.

## Available Endpoints

### Sign Up

The Sign Up endpoint allows users to create a new account in the GoDate app. The endpoint accepts a `POST` method in the `/api/v1/signup` route.

Example CURL:

```shell
curl --location --request POST 'http://localhost:8080/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@mail.com",
    "password": "admin",
    "first_name": "Test",
    "last_name": "User",
    "birth_date": "1998-01-20",
    "gender": "male"
}'
```

Example Success Response:

```json
{
    "success": true,
    "message": "Account created successfully"
}
```

Example Error Response:

```json
{
    "success": false,
    "message": "failed to insert account: UNIQUE constraint failed: accounts.email"
}
```

### Login

The Login endpoint allows users to authenticate and log into their GoDate app account. The endpoint accepts a `POST` method in the `/api/v1/login` route. If the response is success, it returns a token which can be used to authorize login users. Once you get the token, put it on the next request as an HTTP header named `X-Authorization` with value `Bearer <your-token>`. 

**Note**: Please pay attention to the `Bearer` prefix on the header's value.

Example CURL:

```shell
curl --location --request POST 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@mail.com",
    "password": "admin"
}'
```

Example Success Response:

```json
{
    "success": true,
    "message": "Login successful",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAbWFpbC5jb20iLCJpYXQiOjE2MzI0MzIwMjcsImV4cCI6MTYzMjQzMjE4N30.0YQ0p8Q9e2ZQz2Z7XQJZzXzJ3Q9wzWz8vz7QXJ6X0fA"
}
```

Example Error Response:

```json
{
    "success": false,
    "message": "Invalid email or password"
}
```

### Hello World

Hello World is a testing API to test the Authorization functionality is working correctly or not. It accepts a `GET` method in the root (`/`) route.

Example CURL:

```shell
curl --location 'http://localhost:8080' \
--header 'X-Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjM1NTkxMzUsImlzcyI6IkdvRGF0ZSIsInN1YiI6IjEifQ.yvPhg8skGWTFA6rq7NClHbNOE2uFmNXQZkfbtY0R0DA'
```

Example Success Response:

```plaintext
Hello, World!
```

Example Error Response:

```plaintext
Unauthorized
```

Fell free to explore the project with any possibilities you have in mind. Should you have any questions, please reach me on [Instagram](https://instagram.com/jaysyanshar).
