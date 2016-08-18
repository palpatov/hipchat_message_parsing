FROM golang:1.7

# Copy the application files (needed for production)
ADD . /go/src/github.com/palpatov/hipchat_message_parsing

# install dependencies
RUN go get github.com/gorilla/mux

# install this service
RUN go install github.com/palpatov/hipchat_message_parsing

# Expose the service on port 8080.
EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT /go/bin/hipchat_message_parsing