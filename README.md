# BTG Backend Developer Test

A full-stack CRUD application for managing customer data along with their family members.

- **`go-backend/`** — REST API built with Go, Clean Architecture, `gorilla/mux`, and PostgreSQL
- **`laravel-frontend/`** — Web UI built with Laravel (Blade + Laravel Form), calling the Go API over HTTP

---

## Prerequisites

- [Go](https://go.dev/dl/) 1.21+
- [PostgreSQL](https://www.postgresql.org/download/) 14+
- [PHP](https://www.php.net/downloads) 8.0–8.2 (Laravel 9 does not support PHP 8.3+) and [Composer](https://getcomposer.org/)
- [Laravel installer](https://laravel.com/docs/9.x/installation)

> If your default `php` is a newer version (check with `php --version`), install PHP 8.2 separately, e.g. on Mac: `brew install php@8.2`, then run `composer`/`artisan` commands with the full path: `/opt/homebrew/opt/php@8.2/bin/php`.

---

## 1. Database Setup

Create the database and run the schema migration:

```bash
psql postgres -c "CREATE DATABASE backend_test;"
cd go-backend
psql -d backend_test -f migrations/schema.sql
```

Verify the tables were created:

```bash
psql -d backend_test -c "\dt"
```

You should see `customer`, `family_list`, and `nationality`.

Insert nationality data (required, since `customer` references it, and the frontend dropdown needs options). A seed file with ~48 common nationalities is included:

```bash
psql -d backend_test -f migrations/seed_nationalities.sql
```

---

## 2. Go Backend Setup

```bash
cd go-backend
go mod download
```

Create a `.env` file in `go-backend/` with your local PostgreSQL credentials:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_postgres_username
DB_PASSWORD=
DB_NAME=backend_test
API_KEY=choose-any-secret-string
ALLOWED_ORIGIN=http://127.0.0.1:8000
```

> Mac/Linux (Homebrew Postgres): `DB_USER` = your OS username, `DB_PASSWORD` = empty.
> Check your username with `psql postgres -c "SELECT current_user;"`.

> `API_KEY` — any string, must match `GO_API_KEY` in `laravel-frontend/.env` below.

Run the server:

```bash
go run cmd/server/main.go
```

You should see:

```
✅ Connected to PostgreSQL
🚀 Server running on http://localhost:8080
```

---

## 3. Laravel Frontend Setup

In a **new terminal tab** (keep the Go server running):

```bash
cd laravel-frontend
composer install
```

> If `php --version` isn't 8.0–8.2, use the full path instead, e.g. `/opt/homebrew/opt/php@8.2/bin/php /opt/homebrew/bin/composer install`.

Create/update `.env` in `laravel-frontend/` and add:

```
GO_API_URL=http://localhost:8080/api
GO_API_KEY=choose-any-secret-string
```

> `GO_API_KEY` must match `API_KEY` in `go-backend/.env`.

Generate the app key if not already set:

```bash
php artisan key:generate
```

> Same note as above — use the full PHP 8.2 path if needed: `/opt/homebrew/opt/php@8.2/bin/php artisan key:generate`.

Run the server:

```bash
php artisan serve
```

> Or with explicit PHP 8.2: `/opt/homebrew/opt/php@8.2/bin/php artisan serve`.

Visit **http://127.0.0.1:8000** in your browser. You should see the Customer Management page.

---

## Security

- All `/api/*` routes require header `X-API-Key`, checked against `API_KEY` env var. No key → `401`.
- CORS locked to `ALLOWED_ORIGIN` env var — only the Laravel frontend's origin can call the API from a browser.
- Rate limit: 100 requests/minute per IP → `429` if exceeded.
- No login/session system exists in this test, so the API key is a static shared secret rather than per-user auth.

---

## Usage

- **List customers** — home page (`/`)
- **Add customer** — click "+ Add New Customer", fill in the form, use "+ Tambah Keluarga" to add family member rows (and "Hapus" to remove a row)
- **Edit customer** — click "Edit" on any row
- **Delete customer** — click "Delete" on any row (family members are cascade-deleted automatically)

---

## Architecture Notes

### Go Backend — Clean Architecture

```
go-backend/
├── cmd/server/           # Application entry point
├── internal/
│   ├── entity/           # Plain data structs (Customer, FamilyList, Nationality)
│   ├── repository/       # SQL queries — the only layer that talks to PostgreSQL
│   ├── usecase/          # Business logic (e.g. creating a customer + family in one transaction)
│   ├── delivery/http/    # HTTP handlers and gorilla/mux routing
│   └── config/           # Database connection setup
└── migrations/           # SQL schema
```

Each layer only depends on the layer "inside" it (delivery → usecase → repository → entity), never the reverse. Creating a customer with family members runs inside a single database transaction — if any part fails, everything is rolled back.

### Laravel Frontend

Laravel does not connect to PostgreSQL directly — it is a pure frontend that calls the Go API over HTTP using Laravel's built-in HTTP client (`Illuminate\Support\Facades\Http`). All forms use Laravel's Blade form conventions (`@csrf`, `@method`), and the dynamic "add/remove family member" rows are handled with a small amount of vanilla JavaScript.

---
