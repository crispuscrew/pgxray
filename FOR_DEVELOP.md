# Development Guide

## Stack

- **Go 1.24** - all source under `src/`
- **bubbletea v2** (`charm.land/bubbletea/v2`) - TUI framework
- **lipgloss v2** (`charm.land/lipgloss/v2`) - styling
- **bubbles v2** (`charm.land/bubbles/v2`) - UI components (progress bar etc.)
- **pgx/v5** (`github.com/jackc/pgx/v5`) - PostgreSQL driver
- **go-toml/v2** (`github.com/pelletier/go-toml/v2`) - config parsing
- **cobra** (`github.com/spf13/cobra`) - CLI
- **golangci-lint** - internal go linter
- All builds and tests run inside containers (podman preferred, docker fallback)

## Workflow

```bash
make tidy        # update go.mod / go.sum inside container
make build       # compile binary → ./bin/pgxray
make run         # run binary directly (TUI needs live terminal)
make start       # build + run
make lint        # run golangci-lint
make test        # run tests (spins up postgres container)
make db          # start local postgres on :5432 for manual testing
```

Selective testing:
```bash
make test TEST_PKG=./src/internal/db/...
```

Local dev with DB:
```bash
make db                    # terminal 1 - postgres stays up until ctrl+c
PGPASSWORD=test make run   # terminal 2
```

## Project structure

```
src/
├── cmd/pgxray/main.go          # entry point
└── internal/
    ├── opt/                    # generic Opt[T] optional wrapper
    ├── cfg/                    # config loading, CLI flags, merge
    ├── cli/                    # cobra CLI
    ├── db/                     # postgres connection and queries
    └── ui/
        ├── app.go              # program init
        ├── model.go            # root bubbletea model
        ├── view.go             # root view
        ├── common/             # shared interfaces and messages
        ├── colors/             # color palette
        ├── loader/             # loading screen with progress bar
        ├── toast/              # notifications
        ├── db/                 # DB connection cmds and messages
        └── sidebar/            # navigation tree (WIP)
```

## Key conventions

- **No plaintext passwords in config** - use `PGPASSWORD` or pgpassfile
- **All SQL** runs in `REPEATABLE READ + READ ONLY` transaction
- **Component interface** - all UI components implement `common.Component`
- **Messages** shared across packages live in `common/msg.go`
- **opt.Opt[T]** - used for config fields that may or may not be set; implements `encoding.TextUnmarshaler` for TOML
- Containerfile instead of Dockerfile
