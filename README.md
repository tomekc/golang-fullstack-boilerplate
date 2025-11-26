# Golang Self Contained App Boilerplate

## Project Structure

This project uses npm workspaces:
- Root workspace manages the overall project
- Frontend webapp is located at `frontend/webapp/`
- The webapp follows standard SvelteKit project structure with `src/` directory

## Modes of Operation

### Production

The production-built frontend code is embedded into the Go binary: the program is totally self-contained. The SvelteKit build output from `frontend/webapp/build/` is embedded at compile time.

### Development

Run the frontend in dev mode for auto-reloading and the backend separately:

**Frontend (from project root):**
```sh
npm run dev
```

**Backend:**
```sh
go run main.go
```

The frontend dev server runs on port 3000, and the backend API server runs on port 3001.

## Building

To create a production version:

```sh
npm run build
```

This builds the SvelteKit app to `frontend/webapp/build/`. When you compile the Go binary, these files are embedded automatically.

You can preview the production build with `npm run preview`.
