# fabdb

The Sentry Labs Fabrication Database

## Developing

We have a `Makefile` for the backend and standard npm run scripts for the front! Use them!

We use the [go-migrate v4](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) cli tool to manage DB migrations.

Install it with:
```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Licensing

This program is released under the [AGPL-3.0-only](https://www.gnu.org/licenses/agpl-3.0.html).

Copyright (c) 2024 Martin Chaperot
