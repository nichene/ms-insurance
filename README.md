## MS insurance

API service developed for challenge.

## Pre configurations

- To load env vars copy the `.env.example` to a `.env` file and modify accordingly.

- Point your DB vars to your desired DB. To run a local DB use the `make db` command.

## Migrations

- It is necessary to install Goose to run DB migrations. Please refer to the docs at: [goose](https://github.com/pressly/goose) for instalation guide.

- To run migrations use `make migrate-up`.

## API

To run the server locally you can type `make run-api`.

## Swagger

- To generate endpoint documentation swagger must be installed. Please refer to their docs at: [swagger](https://github.com/pressly/goose) for instalation guide.

- It can usually be installed by running `go install github.com/swaggo/swag/cmd/swag@latest`

- after installed generate docs by running `make swag`

- go to `localhost:8080/swagger` endpoint to see docs

## Testing

Run the project tests by running `make tests`.

## Generate mocks

- To generate test mocks install mockgen `go install github.com/golang/mock/mockgen@v1.6.0`. Docs at: [mockgen](https://github.com/golang/mock)

- run `make mocks`

## Other Commands

Go through makefile to checkout all available commands for this project.
