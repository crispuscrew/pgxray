# pgxray — Roadmap & TODO

Status legend: `[ ]` todo · `[~]` in progress · `[x]` done

---

## Phase 0 — Project skeleton

- [ ] `go mod init`
- [ ] Directory structure (see below)
- [ ] Linter config (`golangci-lint`)
- [ ] Makefile with `build`, `run`, `lint`, `test` targets
- [ ] Basic Bubbletea app that starts and quits on `q`

**Suggested directory layout:**

```
pgxray/
├── cmd/
│   └── pgxray/
│       └── main.go          # entry point, flag parsing
├── internal/
│   ├── config/              # config file loading & CLI flag merging
│   ├── db/                  # all Postgres interaction (queries, introspection)
│   └── ui/                  # Bubbletea models and views
│       ├── app.go           # root model
│       ├── sidebar/         # tree navigator
│       ├── table/           # data grid view
│       ├── structure/       # DDL / indexes / constraints view
│       ├── query/           # SQL prompt and results
│       ├── help/            # help overlay
│       └── style/           # lipgloss styles, theme
├── README.md
├── ROADMAP.md
└── go.mod
```

---

## Phase 1 — Config & connection

- [ ] Define config struct (`Config`, `ConnectionProfile`, `Keybindings`)
- [ ] Parse `config.toml` with `github.com/BurntSushi/toml`
- [ ] Config file discovery (flag → env → `~/.config/pgxray/config.toml` → `~/.pgxray.toml`)
- [ ] CLI flag parsing with `flag` stdlib или `github.com/spf13/cobra`
- [ ] Merge CLI flags over config file values
- [ ] Keybindings: hardcoded defaults, overridable via `[keybindings]` section in config
- [ ] Open `pgx` connection pool from resolved config
- [ ] Graceful error if connection fails (show message, exit cleanly)
- [ ] Password resolution chain: `PGPASSWORD` → pgpassfile → interactive prompt
- [ ] pgpassfile location: `--pgpassfile` flag → `pgpassfile` in config → `PGPASSFILE` env → `~/.pgpass`
- [ ] Interactive password prompt if nothing else is set (do not pass empty string to pgx)

**Key decisions:**
- Config format: TOML
- Postgres driver: `github.com/jackc/pgx/v5`
- License: Apache 2.0

---

## Phase 2 — Navigation tree (sidebar)

The sidebar is the backbone of the UX. It's a collapsible tree:

```
▼ myapp (database)
  ▼ public (schema)
    ▶ Tables (12)
    ▶ Views (3)
    ▶ Sequences (5)
  ▶ audit (schema)
▶ postgres (database)
```

- [ ] Fetch databases list (`pg_database`)
- [ ] Fetch schemas per database (`information_schema.schemata`)
- [ ] Fetch tables, views, sequences per schema
- [ ] Bubbletea model: tree with expand/collapse
- [ ] Keyboard navigation (`j`/`k`, `Enter`, `Esc`)
- [ ] Lazy loading: only query children when node is expanded
- [ ] Filter/search within tree (`/`)

---

## Phase 3 — Table data viewer

- [ ] Fetch rows with `LIMIT`/`OFFSET` pagination
- [ ] Detect column types, align values accordingly (numbers right, strings left)
- [ ] Truncate long cell values with `…`, expand on selection
- [ ] Display current page / total row count estimate
- [ ] Keyboard: `n`/`p` for pages, `←`/`→` to scroll columns
- [ ] Column header stays fixed while scrolling rows
- [ ] NULL values displayed distinctly (e.g. dimmed `∅`)
- [ ] Sort by column (`s` to toggle asc/desc)
- [ ] Copy cell value to clipboard (`y`) — best-effort, may depend on terminal

---

## Phase 4 — Structure view

Activated with `d` on any table/view.

- [ ] **Columns tab**: name, type, nullable, default, comment
- [ ] **Indexes tab**: name, columns, unique, partial condition
- [ ] **Constraints tab**: PK, FK, CHECK, UNIQUE — with referenced tables for FKs
- [ ] **DDL tab**: rendered `CREATE TABLE` statement (query `pg_get_tabledef` or reconstruct)
- [ ] Tab switching with `Tab` / `Shift+Tab` or number keys `1`–`4`
- [ ] For views: show the view definition SQL
- [ ] For sequences: show current value, min, max, increment

---

## Phase 5 — SQL query mode

- [ ] Prompt opens with `:`, multi-line input
- [ ] Wrap execution: `BEGIN; SET TRANSACTION ISOLATION LEVEL REPEATABLE READ READ ONLY; ... ROLLBACK;`
- [ ] Pre-execution check: reject obvious write statements (`INSERT`, `UPDATE`, `DELETE`, `DROP`, `CREATE`, `ALTER`, `TRUNCATE`, `GRANT`, `REVOKE`) with a clear error message
- [ ] Display results in the same table grid as Phase 3
- [ ] Query history (in-session, navigable with `↑`/`↓` in prompt)
- [ ] Show execution time and row count
- [ ] Error messages from Postgres displayed clearly (not raw stack traces)

---

## Phase 6 — Polish & UX

- [ ] Help overlay (`?`) listing all keybindings
- [ ] Status bar at the bottom: connection info, current location, row count, page
- [ ] Breadcrumb header: `myapp > public > users`
- [ ] Responsive layout: adapt to terminal width/height on resize
- [ ] Meaningful loading indicators (spinner while queries run)
- [ ] Consistent color theme via lipgloss — consider supporting light/dark terminal backgrounds
- [ ] Graceful handling of lost DB connection (reconnect prompt)
- [ ] `--version` flag

---

## Phase 7 — Packaging & distribution

- [ ] GitHub Actions: build binaries for `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`
- [ ] Attach binaries to GitHub Releases
- [ ] Install script (`curl | sh`) that places binary in `/usr/local/bin`
- [ ] Homebrew tap (optional, for macOS users who want it locally)
- [ ] Man page (`man pgxray`)

---

## Ideas & future scope (not committed)

- Named bookmarks for frequently visited tables
- Export query results to CSV (written to server filesystem)
- Explain plan viewer (`EXPLAIN ANALYZE` with visual breakdown)
- Activity monitor tab (wraps `pg_stat_activity`)
- Slow query log viewer (wraps `pg_stat_statements`)
- Side-by-side diff of two table schemas
