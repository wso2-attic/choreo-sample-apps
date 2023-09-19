# Choreo sample web app with react

### Initilized with [Create React App](https://create-react-app.dev/docs/getting-started)

```shell
npx create-react-app
```

### Use the following configuration when creating this component in Choreo:

- Build Preset: **React SPA**
- Build Context Path: `web-apps/react-spa`
- Build Command: `npm run build`
- Build output directory: `build`
- Node Version: `18`

### Use thr following commands to build and run the app using Docker:

```shell
docker build -t react-spa bring-your-own-image-components/web-apps/react-spa
docker run -p 8080:80 react-spa
```
