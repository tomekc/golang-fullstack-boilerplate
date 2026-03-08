# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go + SvelteKit fullstack boilerplate that embeds the frontend build into the Go binary for self-contained deployment. The Go backend serves both API endpoints and the static frontend files.

## Project Structure

**Workspace Organization:**
- Root uses npm workspaces (pattern: `frontend/*`)
- Frontend webapp: `frontend/webapp/` (package: `@gofullstack/webapp`)
- Root-level npm commands delegate to workspaces using `-w` flag

**Frontend Structure:**
- Standard SvelteKit project at `frontend/webapp/`
- Source code in `frontend/webapp/src/` directory
- Routes in `frontend/webapp/src/routes/`
- Shared utilities and config in `frontend/webapp/src/lib/`
- Build output goes to `frontend/webapp/build/`

**Backend Structure:**
- API handlers organized by feature in `backend/` (e.g., `backend/hello/handler.go`)
- Server framework in `backend/server/`
- Configuration management in `backend/config/`

## Architecture

**Dual-Server Development Setup:**
- Frontend dev server runs on port 3000 (Vite/SvelteKit)
- Backend API server runs on port 3001 (Go)
- In development, frontend proxies API requests to the backend via `baseURL()` helper in `frontend/webapp/src/lib/config.ts`

**Production Build:**
- SvelteKit builds to `frontend/webapp/build/` directory
- Go binary embeds these files using `//go:embed all:frontend/webapp/build` in `main.go`
- Static files served from embedded filesystem via `backend/server/static.go`
- Single deployable binary contains everything

**API Routing:**
- All API endpoints are under `/api/` prefix
- Main router in `backend/server/server_main.go` strips `/api` prefix before passing to API mux
- Static file server handles root path `/` for all non-API routes

## Development Commands

**Start frontend dev server (from root):**
```sh
npm run dev
```
Runs on port 3000 with auto-reload. Delegates to workspace: `@gofullstack/webapp`

**Start backend server:**
```sh
go run main.go
```
Runs on port 3001, serves API endpoints and embedded static files

**Type checking (from root or webapp directory):**
```sh
npm run check
```
For continuous type checking:
```sh
npm run check:watch
```

**Linting (from webapp directory):**
```sh
cd frontend/webapp
npm run lint
```

**Production build (from root):**
```sh
npm run build
```
Outputs to `frontend/webapp/build/` directory

**Preview production build (from webapp directory):**
```sh
cd frontend/webapp
npm run preview
```

**Working with specific workspace:**
```sh
# Run commands in webapp workspace from root
npm run <script> -w @gofullstack/webapp
```

## Adding New API Endpoints

1. Create a new package in `backend/` (e.g., `backend/users/`)
2. Implement handler function with signature `func(w http.ResponseWriter, r *http.Request)`
3. Register in `main.go` using `httpServer.ApiEndpoint("GET /endpoint", handler.Function)`
4. API will be available at `/api/endpoint`

## Key Configuration Files

- `package.json` (root) - Workspace configuration, delegating scripts to `@gofullstack/webapp`
- `frontend/webapp/svelte.config.js` - SvelteKit adapter configured to build to `build/` directory
- `frontend/webapp/src/lib/config.ts` - Contains `baseURL()` helper for API server URL (dev/prod)
- `main.go` - Embeds static files and configures server with API routes
- `backend/server/server_main.go` - Server struct and routing setup
- `backend/server/static.go` - Static file serving from embedded filesystem