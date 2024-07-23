# Discipline GBatch

## 참고 사항
- [Batch Framework] (https://github.com/chararch/gobatch)
- [Dependancy Injection In GO] (https://syntaxsugar.tistory.com/entry/Golang-Dependency-Injection)
- [Golang-migrate 사용법] (https://velog.io/@suji9709/Golang-migrate)
## Prerequisite

```
go get -u github.com/chararch/gobatch
brew install golang-migrate
```

## CLI Script for DB Migration

```
migrate -source file:<path> -database 'postgresql://<username>:<password>@<host>:<port>/<db_name>?sslmode=disable' 1 up

# ex) migrate -source file:./migrations -database "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose  up 1
```