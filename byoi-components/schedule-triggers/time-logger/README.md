# Choreo Sample Schedule Trigger - Time Logger

### Use the following commands to create BYOI component

- Select Create Component with `Schedule Trigger` type.
- Provide component name and description.
- Select `Deploy an image from a Container Registry` as the source.
- Select `Time Logger` tile and create component.

This is a simple program that logs the current time logger

## Build

```bash
go build -o time-logger
```
```

## Run

```bash
./time-logger
```

## Docker Build

```bash
docker build -t time-logger .
```

## Docker Run

```bash
docker run time-logger
```
