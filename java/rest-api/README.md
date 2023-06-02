# Sample Java REST API

## Use the following build config when creating this component in Choreo:

- Build Preset: **Dockerfile**
- Dockerfile Path: `go/rest-api/Dockerfile`
- Docker Context Path: `go/rest-api`

The [endpoints.yaml](.choreo/endpoints.yaml) file contains the endpoint configurations that are used by the Choreo to expose the service.

## Use the following command to build and run the application locally.

> NOTE: You need to have java 17 installed in your system or a docker and VS Code installed to
> open this in a dev container

Use following command to build the project
```bash
$ ./mvnw clean install
```

Run the service
```bash
$ ./mvnw spring-boot:run
```
