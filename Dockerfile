# Use the official golang image as a base image
FROM golang:latest
# Set the working directory inside the container
WORKDIR /app-$(date +%s)

# Clone your git repository
RUN git clone https://github.com/pranjal123662/Subscription-API.git /app-$(date +%s)/Subscription-API            

# Navigate into the cloned directory

# Copy the local package files to the container's workspace
COPY . /Subscription-API

WORKDIR /app-$(date +%s)/Subscription-API
# Expose port 7456 to the outside world
EXPOSE 7689

# Command to run the executable
CMD ["nohup", "./main","&"]