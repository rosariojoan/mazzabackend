Ref: `entgo.io`

Create new model with name `Token`:
    `go run -mod=mod entgo.io/ent/cmd/ent new Token`

generate.go
```
package mazza

//go:generate go run -mod=mod ./ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen@latest --verbose
```

Check and update dependencies: `go mod tidy`
To generate models, run: `go generate .`