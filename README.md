# pgxray

> **Work in progress.** pgxray is under active development and does not yet provide its full intended functionality. Not recommended for production use.

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

## Configuration

pgxray reads `~/.config/pgxray/config.toml` by default. Use `--config` to specify a different path.

Config file locations (checked in order):

1. `--config /path/to/config.toml`
2. `$PGXRAY_CONFIG_PATH`
3. `~/.config/pgxray/config.toml`

### Profiles

A profile is a named connection configuration. Profile selection order:

1. `--profile` flag if provided
2. Profile named `default` in the config file

All profile fields can be overridden via the corresponding CLI flag. Run `pgxray --help` for the full flag reference.

```toml
[connections.default]
host      = "localhost"
port      = 5432
user      = "postgres"
database  = "postgres"
sslmode   = "disable"

[connections.prod]
host       = "127.0.0.1"
port       = 5432
user       = "readonly"
database   = "myapp"
sslmode    = "require"
pgpassfile = "/etc/pgxray/pgpass"
```

**Passwords** - resolved in this order:

1. `PGPASSWORD` environment variable
2. pgpass file (first match wins):
   - `--pgpassfile` CLI flag
   - `pgpassfile` field in the active profile
   - `PGPASSFILE` environment variable
   - `~/.pgpass`
3. Interactive prompt at startup

Plaintext passwords in config are intentionally not supported.

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
# Use default profile
pgxray

# Use named profile
pgxray --profile prod

# Custom config, then override a field
pgxray --config /etc/pgxray/config.toml --profile prod --database staging
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

- Go 1.24+ (to build from source)
- PostgreSQL 12+
- A terminal with 256-color support (most modern terminals qualify)
