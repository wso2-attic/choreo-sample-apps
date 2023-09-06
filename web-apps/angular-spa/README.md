# Choreo sample web app with angular

### Initilized with [Angular CLI](https://angular.io/cli)

```shell
npx @angular/cli new angular-spa --defaults
```

### Use the following configuration when creating this component in Choreo:

- Build Preset: **Angular SPA**
- Build Context Path: `web-apps/angular-spa`
- Build Command: `npm run build`
- Build output directory: `dist/angular-spa`
- Node Version: `18`

### Use thr following commands to build and run the app using Docker:

```shell
docker build -t angular-spa web-apps/angular-spa
docker run -p 8080:80 angular-spa
```
