# Use an official Node.js runtime as the base image
FROM node:18 AS build

# Set the working directory in the container
WORKDIR /app

# Copy the package.json and package-lock.json to the container
COPY package*.json ./

# Install the application dependencies
RUN npm install

# Copy the entire Angular app source code to the container
COPY . .

# Build the Angular application for production
RUN npm run build 

# Use a lightweight Nginx image as the final image
FROM nginx:alpine

# Copy the built Angular app from the previous stage to the Nginx web server directory
COPY --from=build /app/dist/angular-spa /usr/share/nginx/html
COPY nginx/default.conf /etc/nginx/conf.d/

# Expose port 80 for serving the web application
EXPOSE 80

# Start the Nginx web server when the container runs
CMD ["nginx", "-g", "daemon off;"]
