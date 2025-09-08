## Makefile

This project uses a `Makefile` to simplify common tasks like linting, testing, building, running, and applying migrations.

### Common commands

| Command        | Description                                      |
|----------------|--------------------------------------------------|
| `make lint`    | Runs Go linter (`golangci-lint`)               |
| `make test`    | Runs all Go tests                               |
| `make migrate` | Applies database migrations via `scripts/migrate-up` |

### Example usage

```bash
# Run linter
make linter

# Run tests
make test

# Apply migrations
make migrate
```

### How to run?

Create your `.env` according to `.env.example` file and create postgresql instance