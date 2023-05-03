# Choreo Sample GraphQL Service - Reading List

### Use the following configuration when creating this component in Choreo:

- Build Preset: **Ballerina**
- Project Path: `ballerina/reading-list`

The [endpoints.yaml](.choreo/endpoints.yaml) file contains the endpoint configurations that are used by the Choreo to expose the service.

### Test the service locally

Execute the command below to run this service:

```shell
bal run
```

Execute the following sample cURL command to add a book item to the reading list:

```shell
curl -X POST -H "Content-type: application/json" -d '{ "query": "mutation { addBook(book: {title: \"Sample Book Name\", author: \"Test Author\"}) { id title author status } }" }' 'http://localhost:8090'
```

### Test the service in Choreo

Deploy the created component in Choreo and navigate to the test page. In the GraphQL console, execute the below sample query to add a book item to the reading list:

```
mutation { addBook(book: {title: "Sample Book Name", author: "Test Author"}) { id title author status } }
```