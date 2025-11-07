# go-graphql
GraphQL server in Go using gqlgen library

## Note
For using backtick(`) in mac use below keys combination
```bash
option + `
```

### Step 1: Initialize a Go module
```bash
go mod init github.com/amitpoitrix/go-graphql
```
where "github.com/amitpoitrix/go-graphql" is naming convention for creating go module which will create go.mod file

### Step 2: Now install gqlgen properly

```bash
go get github.com/99designs/gqlgen@latest
```

Or if you just want to install the CLI globally (to run gqlgen init):

```bash
go install github.com/99designs/gqlgen@latest
```

Make sure your $GOPATH/bin is in your system’s PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Step 3: Initialize gqlgen

Once installed, run:
```bash
gqlgen init
```


This will generate the default structure:
```bash
graph/
  ├── generated.go
  ├── model/
  ├── resolver.go
  └── schema.graphqls
server.go
```