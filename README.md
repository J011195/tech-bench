# Tech bench

## Purpose

## Commands

- Build & run

```bash
mkdir -p bin
env GOOS=windows GOARCH=amd64 go build -o ./bin . && ./bin/tech-bench.exe
```

- Install binary

```bash
go install -ldflags="-s -w" -v github.com/J011195/tech-bench@latest
```
