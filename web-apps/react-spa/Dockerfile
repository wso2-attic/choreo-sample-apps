# Use an official Node.js runtime as a parent image
FROM node:16 AS build

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy the rest of the application code to the container
COPY . .

# Build the React app for production
RUN npm run build

# Use a lightweight Node.js runtime as the final parent image
FROM node:16-slim

# Set the working directory in the final image
WORKDIR /app

# Copy the built app from the previous stage
COPY --from=build /app/build ./build

# Expose port 80 to the outside world
EXPOSE 80

# Define environment variables if needed (e.g., REACT_APP_API_URL)

# Start the application
CMD ["npx", "serve", "-s", "build", "-l", "80"]
