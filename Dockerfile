# Base image
FROM golang:1.9.0-alpine3.6
MAINTAINER Georgy Ramonov <georgy.ramonov@gmail.com>

# Install dev dependencies -- realize for live reload and dep for
# dependency management
RUN apk update && apk add git
RUN go get github.com/tockins/realize
RUN go get github.com/golang/dep/cmd/dep

# Copy the rest of the files. Unlike copying requirements.txt file in
# a separate layer in Python, we do not copy dependency files
# separately because dep needs existing code to ensure proper sync.
WORKDIR /go/src/app
COPY . /go/src/app

# Start realize watcher
ENV APP_PORT=80
EXPOSE $APP_PORT
CMD ["sh", "-c", "dep ensure && realize start"]
