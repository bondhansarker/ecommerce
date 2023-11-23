# ecommerce

A great starting point for building REST APIs in Go using Gin framework, and sqlx for connecting to a PostgreSQL database. The implementation follows Clean Architecture principles as described by Uncle Bob.

### Features

-   Implements the Clean Architecture pattern for a scalable and maintainable codebase
-   Uses the Gin framework for efficient and fast handling of HTTP requests
-   Integrates with PostgreSQL databases using SQLx.DB for powerful and flexible database operations

### Getting Started

To get up and running with the project, follow these simple steps:

```
$ git clone https://github.com/bondhansarker/ecommerce.git
$ cd ecommerce
$ cp internal/config/.env.example internal/config/.env # create a copy of the example environment file, and also follow configuration steps on the difference section below
$ docker compose up
```

#### Configuration

The application can be configured using environment variables to fit your specific needs. A sample environment file is provided as .env.example with the following variables available for customization:

##### App

-   `PORT`: The port on which the server will listen (defaults to 8080)
-   `ENVIRONMENT`: The environment the application is running in (defaults to "development")
-   `DEBUG`: Enable or disable debug mode (defaults to true)

##### Database

PostgreSQL

-   `DB_POSTGRES_DRIVER`: The database driver to use (defaults to "postgres")
-   `DB_POSTGRES_DSN`: The database connection URI in DSN format (defaults to "user=myuser password=mypassword host=myhost port=5432 dbname=mydb sslmode=disable timezone=Asia/Jakarta")
-   `DB_POSTGRES_URL`: The database connection URI in URL format (defaults to "postgres://user:pass@host/db")

### Folder Structure

```
root/
|-- cmd/
|   |-- api
|   |   |-- server/
|   |   |-- main.go
|-- db_scripts/
|   |--init.sql
|-- internal/
|   |-- business/
|   |   |-- domains
|   |   |   |-- v1
|   |   |       |-- domains.product.go
|   |   |-- UseCases
|   |       |-- v1
|   |           |-- UseCase.product.go
|   |-- config/
|   |   |-- .env
|   |   |-- .env.example
|   |   |-- config.go
|   |-- constants/
|   |-- datasources/
|   |   |-- drivers/
|   |   |   |-- driver.sqlx.go
|   |   |-- records/
|   |   |   |-- record.product.go
|   |   |   |-- record.user_mapper_v1.go
|   |   |-- repositories
|   |   |   |-- postgres/
|   |   |   |   |-- v1
|   |   |   |       |-- postgre.product.go
|   |-- http/
|   |   |-- datatransfers/
|   |   |   |-- requests/
|   |   |   |   |-- request.product.go
|   |   |   |-- responses/
|   |   |       |-- response.product.go
|   |   |-- handlers/
|   |   |   |-- v1/
|   |   |       |-- handler.base_response.go
|   |   |       |-- handler.product.go
|   |   |-- routes/
|   |       |-- route.product.go
|   |-- utils/
|-- pkg/
|   |-- helpers/
|   |   |-- helper.common.go
|   |-- logger/
|-- vendor/
|-- go.mod
|-- go.sum
|-- README.md (thisfile)
```

##### `cmd` folder

This folder contains all the entry points of the application. There are four sub-folders in the `cmd` folder:

-   `api`: This folder contains the main entry point of the REST API server. The `main.go` file in the `server` sub-folder is responsible for starting the server and setting up all the necessary routes.


##### `docs` folder

This folder contains the documentation for the REST API, including the `swagger.yaml` file which defines the API specification.

##### `internal` folder

This folder contains all the business logic and other implementation details of the application. It is structured as follows:

-   `business` folder

    -   domains folder: This folder contains domain-specific logic, such as the business rules for creating, updating, and deleting users.

    -   UseCases folder: This folder contains the implementation of the use cases that are defined in the domains folder.

-   `config` folder

    -   `.env`: This file contains the environment variables that are used by the application.
    -   `.env.example`: This file is an example of the .env file, with all the necessary environment variables listed.
    -   `config.go`: This file reads the environment variables and sets up the configuration for the application.

-   `constants` folder

    -   this folder contains constant values used throughout the application.

-   `datasources` folder

    -   `drivers` folder: This folder contains the implementation of database drivers, such as PostgreSQL.
    -   `records` folder: This folder contains the implementation of database records, such as Product.
    -   `repositories` folder: This folder contains the implementation of database repositories, such as PostgreSQL and MongoDB.

-   `http` folder

    -   `datatransfers` folder: This folder contains the implementation of data transfer objects, such as request and response objects.
    -   `handlers` folder: This folder contains the implementation of HTTP handlers, which handle incoming HTTP requests and send responses back to the client.
    -   `routes` folder: This folder contains the implementation of routes, which map URLs to handlers.


-   `utils` folder

    -   this folder contains utility functions and classes used throughout the application.

##### `pkg` folder

This folder contains reusable packages that are shared across different parts of the application.

##### `Dockerfile`: This file is used to build a Docker image of the application.
##### `docker-compose.yml`: This file is used to start the application and its dependencies (such as the database) using Docker Compose.
