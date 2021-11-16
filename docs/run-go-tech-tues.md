# run go-tech-tues

[go-tech-tues](https://tredfield.github.io/go-tech-tues/) contains a small API built in go

It has a `widgets` endpoint that returns a list of widgets and allows querying a widget by `Id`

It also provides a cli using [cobra](https://github.com/spf13/cobra)

## Layout of project

The basic layout of the project is below

```bash
.
├── README.md
├── cmd
│   ├── root.go
│   └── start.go
├── docs
├── go.mod
├── go.sum
├── main.go
├── pkg
│   └── api
└── static
```

- `main.go` contains single `main` function to run the app

- `cmd` contains `cmd` package and files that implement the CLI

- `pkg/api` contains `api` package files that implement the api

- `static` static html/css/images

- `go.mod, go.sum` provide dependency management

## build and run

```bash
# build
go build
```

```bash
# run cli
./go-tech-tues
```

```bash
# run start help
./go-tech-tues start --help
```

```bash
# finally run sever
./go-tech-tues start
```

Curl endpoints

```bash
# root
curl localhost:8080
```

```bash
# curl widgets
curl localhost:8080/widgets
```

```bash
# curl widgets{3}
curl localhost:8080/widgets/3
```

Serve html

```bash
# open in browser
open localhost:8080
```
