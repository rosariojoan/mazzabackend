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


## Docker
Build containers: `docker compose up --build`
Run the containers: `docker compose up`

List running services: `docker compose ps`

If there is a problem accessing postgres, reset the password.
Run this command within the container: `sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"`
