## Description

Golang server with a clear separation for easy scaling and testing

## Structure

```bash
└── internal/
    └── app/
        └── example/
            ├── api/
            │   └── feature1/
            │       └── main.go
            ├── application/
            │   ├── feature1/
            │   │   └── main.go
            │   └── interfaces.go
            ├── domain/
            │   ├── feature1/
            │   │   └── main.go
            │   └── interfaces.go
            └── infra/
                └── store/
                    ├── sql/
                    │   └── feature1/
                    │       └── main.go
                    └── interfaces.go
````

## Commands

### Run the server Local

```bash
./scripts/run_server --local
```

### Generate sql code

Need to install sqlc first
https://docs.sqlc.dev/en/latest/overview/install.html

```bash
./scripts/generate_sql
```

### Run migration manually

Inside the codebase, we already auto running migration everytime start the server
However, if you want to run it manually, you can do it by running this command

```bash
./scripts/run_migrate
```

## Playground

Perquisite: link .env file to playground folder
```bash
ln -s $PWD/.env $PWD/playground/.env
```

To playground with some application service, we can run the test inside playground folder
For example, to run the test for `example` application service, we can run this command

```bash
go test -v ./playground/example_playground_test.go
```
