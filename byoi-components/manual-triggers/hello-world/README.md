## Choreo Sample Manual Trigger - Hello World

### Use the following commands to create BYOI component

- Select Create Component with `Manual Trigger` type.
- Provide component name and description.
- Select `Deploy an image from a Container Registry` as the source.
- Select `Hello World` tile and create component.

This is a simple hello world program written in goLang.

## Build

```bash
go build -o hello-world
```

## Run

```bash
./hello-world
```

## Docker Build

```bash
docker build -t hello-world .
```

## Docker Run

```bash
docker run hello-world
```
