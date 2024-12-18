# Tech bench

## Purpose

- Test various developer and operations GO features

## Commands

- Build & run binary

```bash
mkdir -p bin
env GOOS=windows GOARCH=amd64 go build -o ./bin . && ./bin/tech-bench-cli.exe
```

- Build & run a smaller binary

```bash
mkdir -p bin
env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin . && ./bin/tech-bench-cli.exe
```

- Install binary

```bash
# Latest
go install -ldflags="-s -w" -v github.com/J011195/tech-bench-cli@latest
# Specific version
go install -ldflags="-s -w" -v github.com/J011195/tech-bench-cli@v1.0.0
go install -ldflags="-s -w" -v github.com/J011195/tech-bench-cli@v1.0.1
```
