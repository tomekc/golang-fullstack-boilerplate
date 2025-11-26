# Golang Self Contained App Boilerplate

## Full stack starter kit for mere mortals

### Just two ingredients

* **SvelteKit** front-end
* **Golang** backend

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

**Frontend files are embedded** in binary, so ultimately entire app is just single file (and configuration).

## Project Structure

### Front-end

This project uses npm workspaces:
- Root workspace manages the overall project
- Frontend webapp is located at `frontend/webapp/`
- The webapp follows standard SvelteKit project structure with `src/` directory

### Server

All server endpoints are supposed to be under `/api` path.
All other paths are considered static HTML and served by result of build of the Svelte app.
Example server implement just one endpoint:

```http request
GET /api/hello
```

which returns current time. Just enough to present idea, simple enough to delete out of way.

#### Packages and files

- `main.go` is entry point and is placed in project root, because some IDEs get totally confused if project starts somewhere else
- `backend/` is root package for all functional packages:
  - `server` is boilerplate for HTTP server, middlewares, etc.
  - `rest` helpers to build HTTP responses and serialize to JSON
  - `hello` implements logic of `/api/hello` 

## Developing

Start frontend

```shell
npm run dev
```

Will start Svelte app in auto-reloading mode on port **3000**. 
The app will **automatically detect that** and call backend on port **3001**.

Start backend from your IDE with the `-dev` command line parameter, or:

```shell
go run . -dev
```

Backend **will read port number** from `frontend/webapp/src/lib/devconfig.json` to make sure
frontend and backend are aligned. 

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

## Docker

### Building the Docker image

Build the container image:

```sh
docker build -t golang-fullstack-boilerplate .
```

### Running the container

Run the container with the data directory mounted:

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
- Run the application with the `-docker` flag

### Stopping the container

```sh
docker stop boilerplate
docker rm boilerplate
```
