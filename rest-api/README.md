# HTTP "VANILLA" REST API

My main objective while developing this project was to get familiarize with some of the packages of the [Go standard library](https://pkg.go.dev/std).

More specifically, I focused on the [net/http](https://pkg.go.dev/net/http) and [http/httptest](https://pkg.go.dev/net/http/httptest@go1.18.4) packages.

So I decided to create a "Vanilla" HTTP API to understand how Golang works with routing, querying, testing and database connections (it is important to say though, that it does not follows the all of specifications of a REST API :D).

### Project specifications:

- It was (almost) all developed using TDD and following the **RED, GREEN, REFACTOR** cicle.
- It has only two routes. `/ping` is used for testing and `/user` exposes a set of services for creating, reading, updating and deleting users into a database.
- It uses PostgreSQL DBMS.
