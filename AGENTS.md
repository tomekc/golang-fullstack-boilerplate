# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go fullstack boilerplate using HTMX + Templ for server-side rendering. The entire application compiles to a single binary with embedded static assets. No JavaScript framework — dynamic interactions are handled by HTMX, and HTML is rendered server-side using Templ templates.

## Project Structure

- `main.go` — entry point; creates Application, starts HTTP server
- `server/app.go` — `Application` struct holding config, logger, services
- `server/routes.go` — route registration and HTTP handlers (methods on Application)
- `server/server.go` — HTTP server struct, middleware wiring
- `server/config/config.go` — TOML config loading from `data/config.toml`
- `server/middleware/` — request logging (`Logging`), CORS middleware, `flightRecorder`
- `server/rest/responses.go` — JSON response helper (`rest.Ok`)
- `server/services/` — business logic services injected into Application
- `server/views/` — Templ templates and embedded static files:
  - `layout/` — base HTML layout wrapping all pages
  - `pages/` — full page templates (dashboard, hello, about)
  - `components/` — reusable UI components (navbar, sidebar, card, tile, etc.)
  - `fragments/` — HTMX partial response templates
  - `static/` — CSS files, embedded via `//go:embed`
  - `static.go` — serves embedded static files at `/static/`

## Architecture

**Application struct** (`server/app.go`) is the central dependency container, created early in `main()`. It holds configuration, logger, and services. Route handlers are methods on Application, defined in `server/routes.go`.

**Request flow:**
1. `main.go` loads config, creates `Application`, creates `Server`
2. `Server.Start()` registers middleware (Logging, CORS), calls `app.RegisterRoutes(mux)`, registers static file routes
3. Requests pass through middleware chain → handler methods on Application

**Templ templates** (`.templ` files) compile to Go code (`_templ.go` files). Run `templ generate` after modifying `.templ` files. Generated `_templ.go` files are committed to the repo.

**Services** are structs in `server/services/` injected into Application. `ExampleService` is a placeholder meant to be replaced.

## Development Commands

**Run the server:**
```sh
make run
```

**Run with Templ file watching (auto-reload on template changes):**
```sh
make templ-watch
```

**Generate Templ code after editing .templ files:**
```sh
make templ-generate
```

**Build production binary:**
```sh
make build
```
Outputs to `bin/server`.

## Adding New Functionality

**New page:** Create `.templ` in `server/views/pages/`, run `templ generate`, add handler method and route in `server/routes.go`.

**New HTMX endpoint:** Create fragment `.templ` in `server/views/fragments/`, run `templ generate`, add handler and route in `server/routes.go`.

**New service:** Create struct in `server/services/`, add as field in `Application` (`server/app.go`), initialize in `NewApplication()`.

**New JSON API endpoint:** Add handler method on Application using `rest.Ok(w, data)`, register route in `server/routes.go`.

## Key Configuration

- `data/config.toml` — server port and app configuration (TOML format)
- Module name: `boilerplate` (in `go.mod`)
- Default port: 3000
