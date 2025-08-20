# Use an official Node.js image
FROM node:20

# Set working directory inside the container
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install

# Copy the rest of the source code
COPY . .

# Build the NestJS project
RUN npm run build

# Expose the port (default NestJS port)
EXPOSE 3000

# Start the application
CMD ["npm", "run", "start:prod"]
