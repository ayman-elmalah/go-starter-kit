# go-starter-kit

The **go-starter-kit** is a starter kit for Go projects that enables you to easily kickstart your Golang project. This project comes pre-configured with MySQL support and provides the necessary tools to streamline your development process.

Developed by Ayman Elmalah. You can find more of Ayman's work on [GitHub](https://github.com/ayman-elmalah).

## Installation

To install and get started with the **go-starter-kit**, follow these steps:

1. Clone the repository:

```sh
git clone https://github.com/ayman-elmalah/go-starter-kit
```

Change to the project directory:

```sh
cd go-starter-kit
```

## Prerequisites

Before you start using the **go-starter-kit**, make sure you have the following prerequisites installed:

1. **CompileDaemon**: A tool to watch and rebuild your Go application on changes. Install it using the following command:

```sh
go install -mod=mod github.com/githubnemo/CompileDaemon
```

2. **golang-migrate**: A tool to manage database migrations in your Go project. Follow the installation guide [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to install it.

Then run the following command to tidy up your Go modules:

```sh
go mod tidy
```

## Configuration

Copy the example configuration file:

```sh
cp config/config.yaml.example config/config.yaml
```

Open the `config/config.yaml` file and provide your data and database credentials.

## Starting the Project

To start the project, use the following command:

```sh
CompileDaemon -build="go build -o main main.go" -command="./main serve"
```

This command uses CompileDaemon to automatically build and restart your Go application whenever changes are detected.

## Database Migrations

Follow these steps to create and manage database migrations:

### Create Migrations

To create a new migration, run the following command:

```sh
migrate create -ext sql -dir db/migrations -seq create_users_table
```

### Applying Migrations

To apply migrations, use the following commands:

- Migrate Up:

```sh
go run main.go migrate-up
```

- Migrate Down

```sh
go run main.go migrate-down
```

For more advanced migration features and options, refer to the [golang-migrate guide](https://github.com/golang-migrate/migrate).

## Adding a New Module

To add a new module to the project, follow these steps:

1. Create a new directory within `internal/modules` for your new module. You can name it based on the module's functionality.

2. Inside the newly created module directory, set up the following structure:
    - Create a `routes` directory.
    - Create a `controllers` directory.
    - Optionally, create additional directories like `services` and `repositories` as needed for your module's architecture.

3. Within the `routes` directory, create a `routes.go` file. You can model this after the structure in `/internal/modules/home/routes/routes.go`.

4. Similarly, within the `controllers` directory, create a controller file, for example `module_controller.go`. You can follow the structure of `/internal/modules/home/controllers/home_controller.go` as an example.

5. After setting up routes and controllers, add the route(s) defined in your module's `routes/routes.go` to the global router. You can do this by including the route(s) in `/internal/providers/routes/route.go`.

By following these steps, you'll be able to integrate new modules into your project, each with its own set of routes, controllers, and other components. This modular approach helps keep your project organized and maintainable.
