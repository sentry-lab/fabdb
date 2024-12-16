# fabdb

The Sentry Labs Fabrication Database

## Developing

There is a compose file in this repo. It is setup for quick developpment by running the following:
```sh
docker compose up --watch
```
This is configured to automatically rebuild the Go backend and allow Vite's dev server to do it's thing.

We use the [go-migrate v4](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) cli tool to manage DB migrations.
Use the `Makefile` scripts, to easily create and apply database migrations.

Install it with:
```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Environment

You should have a .env file located at the root of the repo. If you are not developing with the compose file,
you need to make sure these environment variables are loaded (with direnv for example).

Here is what the .env file should contain:
```sh
# Go Backend requirements
DB_DATABASE=fabdb
DB_USERNAME=postgres
DB_PASSWORD=postgres

DB_URL=localhost
DB_PORT=5432

BACKEND_PORT=8080

# Vue Frontend requirements
# Must be prepended with VITE_ to be accessible at runtime
FONTEND_PORT=3000
VITE_API_URL="http://127.0.0.1:$BACKEND_PORT"
```

## Licensing

This program is released under the [AGPL-3.0-only](https://www.gnu.org/licenses/agpl-3.0.html).

Copyright (c) 2024 Martin Chaperot
