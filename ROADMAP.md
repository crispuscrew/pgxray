# ROADMAP

Key features required for release.

## Foundation

- [x] Project skeleton - go mod, directory structure, Makefile
- [x] Config loading - TOML, CLI flags, merge, defaults, pgpassfile chain
- [x] Test infrastructure - containerized postgres, network setup, make test
- [x] DB types - Database, Schema, Table, Column, Index, Constraint, View, Sequence
- [x] DB connection - pgx/v5, timeout, connection string builder
- [x] TUI foundation - bubbletea v2, component system, theme/colors
- [x] Loader - progress bar with smoothing, stage-based progress
- [x] Toast notifications - warnings, errors, info, auto-expire
- [x] Linter - golangci-lint via container

## CI

- [x] Linter container (`make lint`)
- [x] Test infrastructure - containerized postgres, `make test`
- [x] cfg unit tests
- [ ] db integration tests
- [ ] GitHub Actions - run lint + tests on PR

## Core

- [ ] Sidebar tree - databases → schemas → tables/views/sequences, expand/collapse, lazy load
- [ ] Table data viewer - paginated rows, column types, NULL display, sort
- [ ] Structure view - columns, indexes, constraints, DDL tabs
- [ ] SQL prompt - `:` to open, read-only transaction wrap, results in table grid

## UX

- [ ] Help overlay (`?`)
- [ ] Status bar - connection info, location, row count
- [ ] Breadcrumb header
- [ ] Responsive layout on terminal resize
- [ ] Graceful reconnect on lost connection

## Distribution

- [ ] GitHub Actions - binaries for linux/darwin × amd64/arm64
- [ ] GitHub Releases with install script
