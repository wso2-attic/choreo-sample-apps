# Use an official Node.js runtime as a parent image
FROM node:18 AS build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the rest of the application code to the container
COPY . .

# Build the Vue.js app for production
RUN npm run build

# Use a lightweight web server as the final parent image
FROM nginx:alpine

# Copy the built app from the previous stage to the nginx web server directory
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80 to the outside world
EXPOSE 80

# Start the nginx web server
CMD ["nginx", "-g", "daemon off;"]
