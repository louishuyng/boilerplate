### Description

Golang server with a clear separation for easy scaling and testing

### Structure

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
