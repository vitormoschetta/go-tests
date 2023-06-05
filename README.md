# Tests


### Run

Current directory:
```bash 
go test
```

All subdirectories:
```bash
go test ./...
```

### Run with coverage

```bash
go test -cover
```

### Coverage

Current directory:
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

All subdirectories:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```



### References

https://go.dev/blog/cover