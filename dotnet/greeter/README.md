# .NET Greeter Service

Please refer the Choreo documentation under [Develop a REST API](https://wso2.com/choreo/docs/develop-components/develop-services/develop-a-rest-api/#step-1-create-a-service-component-from-a-dockerfile) for deployment steps.

### Use the following config when creating this component in Choreo:

- Dockerfile: `dotnet/greeter/Dockerfile`
- Docker context: `dotnet/greeter/`

### Use the following command to build and run the application locally:

> NOTE: You need to have .Net SDK installed in your machine to build and run this application.

Navigate to the .NET Application Directory:

```bash
cd choreo-sample-apps/dotnet/greeter
```

Build the application

```bash
dotnet build
```

Run the service

```bash
dotnet run
```
