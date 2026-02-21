# pgxray

A fast, keyboard-driven TUI for exploring PostgreSQL databases - designed to run **on the server**, accessed over SSH. No local client setup, no port forwarding, no data leaving the host.

```
ssh user@myserver pgxray
```

---

## Philosophy

- **Read-only by design.** pgxray is a viewer, not an editor. It will never modify your data.
- **SSH-native.** The binary runs on the database host. Your terminal is just a window.
- **Safe queries.** Ad-hoc SQL runs inside a `REPEATABLE READ` transaction on a read-only connection - you see a consistent snapshot and cannot accidentally mutate anything.
- **Keyboard first.** Every action has a keybinding. The mouse is never required.

---

## Features

- Browse databases, schemas, tables, views, and sequences in a tree navigator
- Paginated table data viewer with column-aware formatting
- Table structure: columns, types, constraints, indexes, foreign keys
- Full DDL preview (`CREATE TABLE ...`)
- Ad-hoc SQL with snapshot isolation (REPEATABLE READ, read-only)
- Multiple named connection profiles in a single config file
- CLI flags to override any config value on the fly

---

## Installation

```bash
# On the server
go install github.com/crispuscrew/pgxray/cmd/pgxray@latest
```

Or download a prebuilt binary from [Releases](https://github.com/crispuscrew/pgxray/releases) and place it somewhere in `$PATH` on the server.

---

## Configuration

pgxray reads `~/.config/pgxray/config.toml` by default.

```toml
# Default profile - used when no --profile flag is given
[connections.default]
host     = "localhost"
port     = 5432
user     = "postgres"
database = "postgres"
sslmode  = "disable"

[connections.prod]
host       = "127.0.0.1"
port       = 5432
user       = "readonly"
database   = "myapp"
sslmode    = "require"
pgpassfile = "/etc/pgxray/pgpass"   # optional: override pgpass location per profile
```

**Config file locations** (checked in order):

1. Path given by `--config /path/to/config.toml`
2. `$PGXRAY_CONFIG`
3. `~/.config/pgxray/config.toml`
4. `~/.pgxray.toml`

**Passwords** - resolved in this order:

1. `PGPASSWORD` environment variable
2. pgpass file (first match wins):
   - `--pgpassfile /path/to/file` CLI flag
   - `pgpassfile` field in the connection profile
   - `PGPASSFILE` environment variable
   - `~/.pgpass` (default fallback)
3. Interactive prompt - if no password is found anywhere, pgxray asks for it at startup

Plaintext passwords in config are intentionally not supported.

### CLI overrides

Any config value can be overridden at runtime:

```bash
pgxray --profile prod
pgxray --host localhost --port 5433 --user alice --database staging
pgxray --config /etc/pgxray/config.toml --profile prod
```

---

## Keyboard shortcuts

> Full reference available inside the app with `?`

| Key | Action |
|-----|--------|
| `j` / `k` or `↓` / `↑` | Move down / up |
| `h` / `l` or `←` / `→` | Navigate breadcrumbs |
| `Enter` | Open selected item |
| `Esc` | Go back |
| `Tab` | Switch panel focus |
| `/` | Filter / search |
| `g` | Go to top |
| `G` | Go to bottom |
| `d` | View DDL |
| `i` | View indexes & constraints |
| `:` | Open SQL query prompt |
| `n` / `p` | Next / previous page (table view) |
| `q` | Quit |
| `?` | Toggle help |

---

## Usage

```bash
# Connect using default profile
pgxray

# Connect using a named profile
pgxray --profile prod

# One-off connection without a config file
pgxray --host localhost --port 5432 --user postgres --database myapp
```

---

## SQL Query Mode

Press `:` to open the query prompt. Queries run inside a `REPEATABLE READ` transaction on a read-only connection:

```sql
-- This is safe - pgxray wraps it automatically:
-- BEGIN TRANSACTION ISOLATION LEVEL REPEATABLE READ READ ONLY;
-- <your query>
-- ROLLBACK;
```

Write statements (`INSERT`, `UPDATE`, `DELETE`, `DROP`, ...) will be rejected before execution.

---

## Requirements

- Go 1.22+ (to build from source)
- PostgreSQL 12+
- A terminal with 256-color support (most modern terminals qualify)
