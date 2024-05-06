## Rularea testelor de coverage

- Pentru afisarea procentului de coverage: `go test . -cover`
- Pentru analiza detaliata a raportului: `go test -coverprofile=c.out` si apoi `go tool cover -html="c.out"`