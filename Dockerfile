# The base go-image
FROM golang:1.18.7 as build

# Create a directory for the app
RUN mkdir /app

# Copy all files from the current directory to the app directory
COPY . /app

# Set working directory
WORKDIR /app

# Run command as descriptiond:
# go build will build an executable file named server in the current directory
RUN go build -o server ./cmd/todolist

# Run the server executable
CMD [ "/app/server" ]