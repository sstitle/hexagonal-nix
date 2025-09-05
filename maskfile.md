# Maskfile

This is a [mask](https://github.com/jacobdeichert/mask) task runner file.

<!-- mdformat-toc start --slug=github --no-anchors --maxlevel=6 --minlevel=1 -->

- [Maskfile](#maskfile)
  - [hello](#hello)
  - [test](#test)
  - [greet [name]](#greet-name)

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

## greet [name]

> Run CLI greeting. Usage: `mask greet Alice` or `mask greet`

```bash
NAME=${name:-World}
go run ./src/adapters/driving "$NAME"
```
