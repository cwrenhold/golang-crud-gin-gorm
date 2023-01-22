# Summary

Simple JSON CRUD API using [Go](https://go.dev/), [Gin](https://gin-gonic.com/), and [GORM](https://gorm.io/), based on the video [here](https://www.youtube.com/watch?v=lf_kiH_NPvM).

# Project setup

There are 3 containers running for this:
* Main dev container
* Postgres container
* PgAdmin container

## Main dev container
This is configured to run on port [3000](http://localhost:3000) by default, and is the main container for the project, running the go web server.

## Postgres container
This is the database for the system, which runs on the default port, however this is not exposed by default. In order to look at the database, you can use the PgAdmin4 container.

## PgAdmin4 container
This is a web based admin tool for the database. It runs on port 5050, and is configured to connect to the Postgres container. This is located at [http://localhost:5050](http://localhost:5050). The login details for this are based on the [devcontainer.env](./.devcontainer/.env) file.

# Configured Tasks

There are two tasks which have been setup in the devcontainer:
* CompileDaemon - This builds and watches the project for changes, and recompiles and restarts the server when changes are detected.
* Run migrations - This runs the [migrations](./migrations/migrate.go) against the database. This will need to be run at least once to initialise the database, and for any subsequent migrations which are required.

# Other notes

Thunder client is installed by default into this dev container, and can be used to test the API. Example requests should be available when the project is loaded, but these are located in the [thunder-tests](./thunder-tests/) folder.

