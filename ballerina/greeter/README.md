# Ballerina Greeter Service

## Repository File Structure

The below table gives a brief overview of the important files in the greeter service.\
Note: The following file paths are relative to the path /ballerina/greeter

| Filepath               | Description                                             |
| ---------------------- | ------------------------------------------------------- |
| service.bal            | Greeter service code written in the Ballerina language. |
| tests/service_test.bal | Test files related to the service.bal file.             |
| Ballerina.toml         | Ballerina configuration file.                           |

## Deploy Application

Please refer to the Choreo documentation under the [Develop a
Ballerina REST API](https://wso2.com/choreo/docs/develop-components/develop-services/develop-a-ballerina-rest-api/) section to learn how to deploy the application.

#### Use the following config when creating this component in Choreo:

- Build Preset: **Ballerina**
- Project Path: `ballerina/greeter`

## Execute the Sample Locally

Navigate to the Ballerina application directory

```bash
cd choreo-sample-apps/ballerina/greeter
```

Run the service

```bash
 $ bal run
```
