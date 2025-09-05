# Maskfile

This is a [mask](https://github.com/jacobdeichert/mask) task runner file.

<!-- mdformat-toc start --slug=github --no-anchors --maxlevel=6 --minlevel=1 -->

- [Maskfile](#maskfile)
  - [hello](#hello)
  - [test](#test)
  - [greet](#greet)
  - [tea](#tea)
  - [serve [port]](#serve-port)

<!-- mdformat-toc end -->

## hello

> This is an example command you can run with `mask hello`

```bash
echo "Hello World!"
```

## test

> Run Go unit tests

```bash
go test ./...
```

## greet

> Run CLI greeting (prompts for name)

```bash
go run ./src/adapters/driving/main.go
```

## tea

> Run Bubble Tea greeting UI (prompts for name)

```bash
go run ./src/adapters/driving/tea_main.go
```

## serve [port]

> Run HTTP server with greeting form

```bash
PORT=${port:-8080}
PORT="$PORT" go run ./src/adapters/driving/http_main.go
```
