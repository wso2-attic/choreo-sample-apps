# .NET Greeter Service

## Repository file structure

The below table gives a brief overview of the important files in the greeter service.\
Note: The following file paths are relative to the path /dotnet/greeter

| Filepath               | Description                                                                                                                                                          |
| ---------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Program.cs             | The .NET-based Greeter service code.                                                                                                                                 |
| Dockerfile             | Choreo uses the Dockerfile to build the container image of the application.                                                                                          |
| .choreo/endpoints.yaml | Choreo-specific configuration that provides information about how Choreo exposes the service.                                                                        |
| openapi.yaml           | OpenAPI contract of the greeter service. This is needed to publish our service as a managed API. This openapi.yaml file is referenced by the .choreo/endpoints.yaml. |

### Let's get started!

## Step 1: Create a service component

1.  Go to [https://console.choreo.dev/](https://console.choreo.dev/cloud-native-app-developer) and sign in. This opens the project home page.
2.  If you already have one or more components in your project, click + Create. Otherwise, go to the Service card and click Create.
3.  This opens the Create a Service pane, where you can give your component a name and a description.
4.  Enter a unique name and a description of the service. For this guide, let's enter the following values, and click Next.

    | Field       | Value           |
    | ----------- | --------------- |
    | Name        | Greetings       |
    | Description | Sends greetings |

5.  To allow Choreo to connect to your GitHub account, click **Authorize with GitHub**.
6.  If you have not already connected your GitHub repository to Choreo, enter your GitHub credentials, and select the repository you created by forking [choreo-sample-apps](https://github.com/wso2/choreo-sample-apps) to install the Choreo GitHub App.

7.  In the **Connect Repository** pane, enter the following information:

    | **Field**               | **Description**           |
    | ----------------------- | ------------------------- |
    | **GitHub Account**      | Your account              |
    | **GitHub Repository**   | choreo-sample-apps        |
    | **Branch**              | **`main`**                |
    | **Build Preset**        | Dockerfile                |
    | **Dockerfile Path**     | dotnet/greeter/Dockerfile |
    | **Docker Context path** | dotnet/greeter            |

8.  Click Create. Once the component creation is complete, you will see the component overview page.

You have successfully created a Service component from a Dockerfile. Now let's build and deploy the service.

## Step 2: Configure the service port with endpoints

We expect to run our greeter service on port 9090. To securely expose the service through Choreo, we must provide the port and other required information to Choreo. In Choreo, we expose our services with endpoints. You can read more about endpoints in Choreo [endpoint documentation](https://wso2.com/choreo/docs/develop-components/develop-services/develop-a-service/#what-are-endpoints-in-service-components).

Choreo looks for an endpoints.yaml file inside the `.choreo` directory to configure the endpoint details of a containerized component. Place the `.choreo` directory at the root of the Docker build context path.

In our greeter sample, the endpoints.yaml file is at dotnet/greeter/.choreo/endpoints.yaml.

## Step 3: Build and deploy

Now that we have connected the source repository, and configured the endpoint details, it's time to build and deploy the greeter service.

To build and deploy the service, follow these steps:

1. On the Deploy page, click **Configure & Deploy**.

2. This will open a right drawer that follows the steps to deploy your component.

3. First, you'll be prompted to add environment variables. For this guide, there's no need to add any. Simply click Next.

4. Following that, you'll be asked to add a config file. For this guide, there's no need to add one. Just click Next.

5. Lastly, you'll see the Endpoints details, sourced from the `endpoints.yaml` file in your GitHub repository. Review the provided details carefully and ensure they're correct before clicking the **Deploy** button.

6. Check the deployment progress by observing the console logs on the right of the page.
   You can access the following scans under **Build**.

   - **The Dockerfile scan**: Choreo performs a scan to check if a non-root user ID is assigned to the Docker container to ensure security. If no non-root user is specified, the build will fail.
   - **Container (Trivy) vulnerability scan**: This detects vulnerabilities in the final docker image.
   - **Container (Trivy) vulnerability scan**: The details of the vulnerabilities open in a separate pane. If this scan detects critical vulnerabilities, the build will fail.

\
Once you have successfully deployed your service, you can [test](https://wso2.com/choreo/docs/testing/test-rest-endpoints-via-the-openapi-console/), [manage](https://wso2.com/choreo/docs/api-management/lifecycle-management/), and [observe](https://wso2.com/choreo/docs/monitoring-and-insights/observability-overview/) it like any other component type in Choreo.

\
For additional information, refer the Choreo documentation under [Develop a REST API](https://wso2.com/choreo/docs/develop-components/develop-services/develop-a-rest-api/).
