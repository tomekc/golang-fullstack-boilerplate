# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go + SvelteKit fullstack boilerplate that embeds the frontend build into the Go binary for self-contained deployment. The Go backend serves both API endpoints and the static frontend files.

## Architecture

**Dual-Server Development Setup:**
- Frontend dev server runs on port 3000 (Vite/SvelteKit)
- Backend API server runs on port 3001 (Go)
- In development, frontend proxies API requests to the backend via `baseURL()` helper in `frontend/lib/config.ts`

**Production Build:**
- SvelteKit builds to `backend/web/` directory (configured in `svelte.config.js`)
- Go binary serves both the frontend static files and API endpoints
- Single deployable binary contains everything

**API Routing:**
- All API endpoints are under `/api/` prefix
- API handlers are in `backend/` organized by feature (e.g., `backend/hello/handler.go`)
- Main router strips `/api` prefix before passing to API mux (see `main.go:24`)

**Frontend Structure:**
- SvelteKit app with source in `frontend/` directory (configured in `svelte.config.js`)
- Routes in `frontend/routes/`
- Shared utilities and config in `frontend/lib/`

## Development Commands

**Start frontend dev server:**
```sh
npm run dev
```
Runs on port 3000 with auto-reload

**Start backend server:**
```sh
go run main.go
```
Runs on port 3001, serves API endpoints

**Type checking:**
```sh
npm run check
```
For continuous type checking:
```sh
npm run check:watch
```

**Linting:**
```sh
npm run lint
```

**Production build:**
```sh
npm run build
```
Outputs to `backend/web/` directory

**Preview production build:**
```sh
npm run preview
```

## Adding New API Endpoints

1. Create a new package in `backend/` (e.g., `backend/users/`)
2. Implement handler function with signature `func(w http.ResponseWriter, r *http.Request)`
3. Register in `main.go` by adding to `apiMux` (see example at `main.go:20`)
4. API will be available at `/api/<endpoint-name>`

## Key Configuration Files

- `svelte.config.js` - SvelteKit adapter configured to build to `backend/web/`
- `frontend/lib/config.ts` - Contains `baseURL()` helper that returns correct API server URL based on dev/prod environment
- `main.go` - API router configuration and server setup