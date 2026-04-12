# Golang Self Contained App Boilerplate

## Full stack starter kit for mere mortals

### Just two ingredients

* **HTMX** for interactivity
* **Templ** for server-side HTML templates
* **Golang** server

This project is a self-contained, simple web application you can use for

* hackathons
* experiments 
* prototypes
* making admin tools
* running on your homelab

### Server-side: Go

Simple, **almost-zero-dependency** project to get you started, but not get into your way.
I use dependencies for functions that are not provided in Go standard library. As of now
it is TOML parser.

**App configuration** is read from `data/config.toml`. TOML was chosen because JSON is not well
readable for humans, [YAML is creature from hell](https://ruudvanasseldonk.com/2023/01/11/the-yaml-document-from-hell), and [the most popular YAML parser](https://github.com/goccy/go-yaml) is not 
developed anymore but with 100+ issues and 30+ pull requests.

### Front-end: HTMX + Templ

Pages are rendered server-side using [Templ](https://templ.guide/) templates compiled to Go code.
Dynamic interactions are handled by [HTMX](https://htmx.org/) — no JavaScript framework needed.
Static assets (CSS) are embedded in the binary.

**The entire app is a single binary** (plus configuration file).

## Project Structure

### Application

The `Application` struct (`server/app.go`) is the central component that holds all dependencies:
configuration, logger, services, and (in the future) database connections. It is created early
during startup and passed to the HTTP server.

Route handlers are methods on `Application`, defined in `server/routes.go`.

### Packages and files

- `main.go` — entry point, creates Application and starts the server
- `server/` — root package for all server code:
  - `app.go` — `Application` struct: groups config, logger, and services
  - `routes.go` — route registration and HTTP handlers (methods on Application)
  - `server.go` — HTTP server, middleware setup
  - `config/` — TOML configuration loading
  - `middleware/` — request logging, CORS
  - `rest/` — JSON response helpers
  - `services/` — business logic (e.g. `ExampleService`)
  - `views/` — Templ templates and static assets:
    - `layout/` — base HTML layout
    - `pages/` — full page templates (dashboard, hello, about)
    - `components/` — reusable UI components (navbar, sidebar, card, tile, etc.)
    - `fragments/` — HTMX partial response templates
    - `static/` — CSS and other static files (embedded in binary)

## Building on Top of This Boilerplate

### Adding a new page

1. Create a `.templ` file in `server/views/pages/`
2. Run `templ generate` (or `make templ-generate`)
3. Add a handler method on `Application` in `server/routes.go`
4. Register the route in `Application.RegisterRoutes()`

### Adding a new HTMX endpoint

1. Create a fragment template in `server/views/fragments/`
2. Run `templ generate`
3. Add a handler method on `Application` in `server/routes.go`
4. Register the route in `Application.RegisterRoutes()`
5. Add `hx-get` or `hx-post` attributes to your page template

### Adding a new service

1. Create a file in `server/services/` with your service struct
2. Add the service as a field in `Application` (`server/app.go`)
3. Initialize it in `NewApplication()`
4. Use it from handler methods via `app.YourService`

### Adding an API endpoint (JSON)

1. Add a handler method on `Application` in `server/routes.go`
2. Use `rest.Ok(w, data)` to return JSON responses
3. Register the route in `Application.RegisterRoutes()`

### Removing the example service

Delete `server/services/example.go`, remove the `ExampleService` field from `Application`
in `server/app.go`, and remove or replace the `HtmxGetTime` and `HtmxEcho` handlers in
`server/routes.go`.

## Developing

Start the server with Templ watching for changes:

```shell
make templ-watch
```

This runs `templ generate --watch` with a proxy, auto-reloading on template changes.

Or run manually:

```shell
make run
```

Server starts on port **3000** (configurable in `data/config.toml`).

## Building

```sh
make build
```

This generates Templ code and compiles the Go binary to `bin/server`.

## Docker

### Building the Docker image

```sh
docker build -t golang-fullstack-boilerplate .
```

### Running the container

```sh
docker run -d \
  -p 3000:3000 \
  -v $(pwd)/data:/data \
  --name boilerplate \
  golang-fullstack-boilerplate
```

This will:
- Map port 3000 on your host to port 3000 in the container
- Mount the `data/` directory from your project to `/data` in the container

### Stopping the container

```sh
docker stop boilerplate
docker rm boilerplate
```
