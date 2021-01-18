# Base image
FROM golang:1.15.6

# Configure the repo url so we can configure our work directory.
ENV REPO_URL=github.com/aipetto/go-aipetto-users-api

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# /app/src/github.com/username/go-aipetto-users-api/src
# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
COPY . $WORKPATH
WORKDIR $WORKPATH/src

RUN go build -o go-users-api .

# Expose port 8081 to the service
EXPOSE 8081

CMD ["./go-users-api"]