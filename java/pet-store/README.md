# Swagger Petstore Sample

This is the java project for a simplified version of the pet store sample hosted at https://petstore3.swagger.io.

## Use the following build config when creating this component in Choreo:

- Build Preset: **Dockerfile**
- Dockerfile Path: `java/pet-store/Dockerfile`
- Docker Context Path: `java/pet-store`

The [endpoints.yaml](.choreo/endpoints.yaml) file contains the endpoint configurations that are used by the Choreo to expose the service.


### Use the following command to build and run the application locally
To run the server, run this task:

```
mvn package jetty:run
```

This will start Jetty embedded on port 8080.
