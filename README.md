## Makefile

This project uses a `Makefile` to simplify common tasks like linting, testing, building, running, applying migrations and generate swagger.

### Common commands

| Command        | Description                                      |
|----------------|--------------------------------------------------|
| `make lint`    | Runs Go linter (`golangci-lint`)               |
| `make test`    | Runs all Go tests                               |
| `make migrate` | Applies database migrations via `scripts/migrate-up` |
| `make swagger` | Generate Swagger API documentation from code annotations |

### Example usage

```bash
# Run linter
make linter

# Run tests
make test

# Apply migrations
make migrate

# Generate swagger
make swagger
```

### How to run?

Create your `.env` according to `.env.example` file and create postgresql instance