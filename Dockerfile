# Use the official golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Clone your git repository
RUN git clone https://github.com/pranjal123662/Subscription-API.git

# Navigate into the cloned directory
WORKDIR /app/Subscription-API
# Copy a file that changes frequently to bust the cache
COPY ./dummy.txt .
# Copy the local package files to the container's workspace
COPY . .

# Build the TwoDB app
RUN go build -o Subscription .
 
# Expose port 7456 to the outside world
EXPOSE 7689

# Command to run the executable
CMD ["nohup", "./Subscription", "&"]