FROM node:18-alpine

# Create app directory
WORKDIR /usr/src/app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)
COPY package*.json ./

RUN npm install

# Bundle app source
COPY . .

# Create a new user with UID 10014
RUN useradd -u 10014 -m choreouser
# Set a non-root user
USER 10014

CMD [ "node", "index.mjs" ]
