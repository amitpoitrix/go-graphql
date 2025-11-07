# go-graphql
GraphQL server in Go using gqlgen library

## Note for Mac users:
For using backtick(`) in mac use below keys combination
```bash
option + `
```

### Step 1: Initialize a Go module
```bash
go mod init github.com/amitpoitrix/go-graphql
```
where "github.com/amitpoitrix/go-graphql" is naming convention for creating go module which will create go.mod file

### Step 2: Install dependencies

```bash
go get github.com/99designs/gqlgen
```

### Step 3: Create folder structure and files as shown below:

```bash
go-graphql/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── graph/
│   │   ├── schema.graphqls
│   │   ├── resolver.go
│   │   ├── schema.resolvers.go
│   │   └── model/
│   │       └── models_gen.go
│   ├── service/
│   │   └── book_service.go
│   ├── repository/
│   │   └── book_repository.go
│   └── domain/
│       └── book.go
├── pkg/
│   └── middleware/
│       └── logger.go
├── config/
│   └── config.go
├── go.mod
├── go.sum
├── gqlgen.yml
└── README.md

```


### Step 4: Generate GraphQL code:

```bash
go run github.com/99designs/gqlgen generate
```

### Step 5: Run server:
```bash
go run cmd/server/main.go
```